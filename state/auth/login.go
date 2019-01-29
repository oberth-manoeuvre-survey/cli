package auth

import (
	"github.com/ActiveState/cli/internal/api"
	"github.com/ActiveState/cli/internal/api/client/authentication"
	"github.com/ActiveState/cli/internal/api/client/users"
	"github.com/ActiveState/cli/internal/api/models"
	"github.com/ActiveState/cli/internal/constants"
	"github.com/ActiveState/cli/internal/failures"
	"github.com/ActiveState/cli/internal/keypairs"
	"github.com/ActiveState/cli/internal/locale"
	"github.com/ActiveState/cli/internal/print"
	secretsapi "github.com/ActiveState/cli/internal/secrets-api"
	"github.com/ActiveState/cli/internal/surveyor"
	survey "gopkg.in/AlecAivazis/survey.v1"
)

func plainAuth() {
	credentials := &models.Credentials{}
	if err := promptForLogin(credentials); err != nil {
		failures.Handle(err, locale.T("err_prompt_unkown"))
		return
	}

	doPlainAuth(credentials)

	if api.Auth != nil {
		_, failure := keypairs.FetchRaw(secretsapi.DefaultClient)
		if failure != nil {
			if secretsapi.FailKeypairNotFound.Matches(failure.Type) {
				_, failure := keypairs.GenerateAndSaveEncodedKeypair(secretsapi.DefaultClient, credentials.Password, constants.DefaultRSABitLength)
				if failure != nil {
					failures.Handle(failure, locale.T("keypair_err_save"))
				}
			} else {
				failures.Handle(failure, locale.T("keypair_err"))
			}
		}
	}
}

func promptForLogin(credentials *models.Credentials) error {
	var qs = []*survey.Question{
		{
			Name:     "username",
			Prompt:   &survey.Input{Message: locale.T("username_prompt")},
			Validate: surveyor.ValidateRequired,
		},
		{
			Name:     "password",
			Prompt:   &survey.Password{Message: locale.T("password_prompt")},
			Validate: surveyor.ValidateRequired,
		},
	}

	return survey.Ask(qs, credentials)
}

func doPlainAuth(credentials *models.Credentials) {
	loginOK, err := api.Authenticate(credentials)

	// Error checking
	if err != nil {
		switch err.(type) {
		// Authentication failed due to username not existing
		case *authentication.PostLoginUnauthorized:
			params := users.NewUniqueUsernameParams()
			params.SetUsername(credentials.Username)
			_, err := api.Client.Users.UniqueUsername(params)
			if err == nil {
				confirmed := false
				prompt := &survey.Confirm{
					Message: locale.T("prompt_login_to_signup"),
				}
				survey.AskOne(prompt, &confirmed, nil)
				if confirmed {
					signupFromLogin(credentials.Username, credentials.Password)
				}
			} else {
				failures.Handle(err, locale.T("err_auth_failed"))
			}
			return
		case *authentication.PostLoginRetryWith:
			var qs = []*survey.Question{
				{
					Name:     "totp",
					Prompt:   &survey.Input{Message: locale.T("totp_prompt")},
					Validate: surveyor.ValidateRequired,
				},
			}
			survey.Ask(qs, credentials)
			if credentials.Totp == "" {
				print.Line(locale.T("login_cancelled"))
				return
			}
			doPlainAuth(credentials)
			return
		default:
			failures.Handle(err, locale.T("err_auth_failed_unknown_cause"))
			return
		}
	}

	print.Line(locale.T("login_success_welcome_back", map[string]string{
		"Name": loginOK.Payload.User.Username,
	}))
}
