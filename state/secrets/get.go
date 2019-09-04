package secrets

import (
	"encoding/json"

	"github.com/spf13/cobra"

	"github.com/ActiveState/cli/internal/failures"
	"github.com/ActiveState/cli/internal/locale"
	"github.com/ActiveState/cli/internal/print"
	"github.com/ActiveState/cli/pkg/cmdlets/commands"
)

// Flags holds the flag values passed through the command line
var Flags struct {
	JSON bool
}

func buildGetCommand(cmd *Command) *commands.Command {
	return &commands.Command{
		Name:        "get",
		Description: "secrets_get_cmd_description",
		Run:         cmd.ExecuteGet,

		Arguments: []*commands.Argument{
			&commands.Argument{
				Name:        "secrets_get_arg_name",
				Description: "secrets_get_arg_name_description",
				Variable:    &cmd.Args.Name,
				Required:    true,
			},
		},
		Flags: []*commands.Flag{
			{
				Name:        "json",
				Description: "secrets_flag_json",
				Type:        commands.TypeBool,
				BoolVar:     &Flags.JSON,
			},
		},
	}
}

// ExecuteGet processes the `secrets get` command.
func (cmd *Command) ExecuteGet(_ *cobra.Command, args []string) {
	secret, value, fail := getSecretWithValue(cmd.Args.Name)
	if fail != nil {
		failures.Handle(fail, locale.T("secrets_err"))
		return
	}

	if Flags.JSON {
		printJSON(&secretJSONDefinition{secret.Name(), secret.Scope(), secret.Description(), *value})
		return
	}

	if value == nil {
		err := "secrets_err_project_not_defined"
		if secret.IsUser() {
			err = "secrets_err_user_not_defined"
		}
		print.Error(locale.Tr(err, cmd.Args.Name))
		cmd.config.Exiter(1)
		return
	}
	print.Line(*value)

	return
}

func printJSON(secretJSON *secretJSONDefinition) {
	var data []byte

	data, err := json.Marshal(secretJSON)
	if err != nil {
		failures.Handle(failures.FailMarshal.Wrap(err), locale.T("secrets_err"))
	}

	print.Line(string(data))
	return
}
