package pkg

import (
	"errors"
	"runtime"
	"sort"

	"github.com/bndr/gotabulate"
	"github.com/go-openapi/strfmt"
	"github.com/spf13/cobra"

	"github.com/ActiveState/cli/internal/failures"
	"github.com/ActiveState/cli/internal/locale"
	"github.com/ActiveState/cli/internal/logging"
	"github.com/ActiveState/cli/internal/print"
	"github.com/ActiveState/cli/pkg/platform/model"
	"github.com/ActiveState/cli/pkg/project"
)

// ListFlags holds the list-related flag values passed through the command line
var ListFlags struct {
	Commit string
}

// ExecuteList lists the current packages in a project
func ExecuteList(cmd *cobra.Command, allArgs []string) {
	logging.Debug("ExecuteList")

	proj := project.Get()

	commit, fail := targetedCommit(proj, ListFlags.Commit)
	if fail != nil {
		failures.Handle(fail, locale.T("package_err_cannot_obtain_commit"))
		return
	}

	recipe, fail := fetchRecipe(proj, commit)
	if fail != nil {
		failures.Handle(fail, locale.T("package_err_cannot_fetch_recipe"))
		return
	}

	pkgs := makePacks(recipe)

	print.Line(pkgs.table())
}

func targetedCommit(proj *project.Project, commitOpt string) (*strfmt.UUID, *failures.Failure) {
	if commitOpt == "latest" {
		return model.LatestCommitID(proj.Owner(), proj.Name())
	}

	commit := commitOpt
	if commit == "" {
		commit = proj.CommitID()
	}

	if ok := strfmt.Default.Validates("uuid", commit); !ok {
		err := errors.New("invalid uuid value")
		return nil, failures.FailMarshal.Wrap(err)
	}

	var uuid strfmt.UUID
	if err := uuid.UnmarshalText([]byte(commit)); err != nil {
		return nil, failures.FailMarshal.Wrap(err)
	}

	return &uuid, nil
}

func fetchRecipe(proj *project.Project, commit *strfmt.UUID) (*model.Recipe, *failures.Failure) {
	if commit == nil {
		return nil, nil
	}

	mproj, fail := model.FetchProjectByName(proj.Owner(), proj.Name())
	if fail != nil {
		return nil, fail
	}

	return model.FetchRecipeForCommitAndHostPlatform(mproj, *commit, runtime.GOOS)
}

type pack struct {
	Name    string
	Version string
}

type packs []*pack

func makePacks(recipe *model.Recipe) packs {
	if recipe == nil {
		return nil
	}

	filter := func(s *string) string {
		return filterNilString("none", s)
	}

	var pkgs packs
	for _, ing := range recipe.ResolvedIngredients {
		pkg := pack{
			Name:    filter(ing.Ingredient.Name),
			Version: filter(ing.IngredientVersion.Version),
		}

		pkgs = append(pkgs, &pkg)
	}

	return pkgs
}

func (ps packs) table() string {
	if ps == nil {
		return locale.T("package_no_data")
	}

	var rows [][]string
	for _, p := range ps {
		row := []string{
			p.Name,
			p.Version,
		}
		rows = append(rows, row)
	}
	if len(rows) == 0 {
		return locale.T("package_no_packages")
	}
	rowsByFirstCol(rows).Sort()

	headers := []string{
		locale.T("package_name"),
		locale.T("package_version"),
	}

	t := gotabulate.Create(rows)
	t.SetHeaders(headers)
	t.SetAlign("left")

	return t.Render("simple")
}

func filterNilString(fallback string, s *string) string {
	if s == nil {
		return fallback
	}
	return *s
}

type rowsByFirstCol [][]string

func (rs rowsByFirstCol) Len() int {
	return len(rs)
}

func (rs rowsByFirstCol) Less(i, j int) bool {
	if len(rs[i]) < 1 || len(rs[j]) < 1 {
		return false
	}

	return rs[i][0] < rs[j][0]
}

func (rs rowsByFirstCol) Swap(i, j int) {
	rs[i], rs[j] = rs[j], rs[i]
}

func (rs rowsByFirstCol) Sort() {
	sort.Sort(rs)
}
