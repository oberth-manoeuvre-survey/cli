package cmdtree

import (
	"github.com/ActiveState/cli/internal/captain"
	"github.com/ActiveState/cli/internal/locale"
	"github.com/ActiveState/cli/internal/primer"
	"github.com/ActiveState/cli/internal/runners/clean"
)

func newCleanCommand(prime *primer.Values) *captain.Command {
	return captain.NewCommand(
		"clean",
		locale.T("clean_description"),
		[]*captain.Flag{},
		[]*captain.Argument{},
		func(ccmd *captain.Command, _ []string) error {
			prime.Output().Print(ccmd.Help())
			return nil
		},
	)
}

func newUninstallCommand(prime *primer.Values) *captain.Command {
	params := clean.UninstallParams{}
	return captain.NewCommand(
		"uninstall",
		locale.T("uninstall_description"),
		[]*captain.Flag{
			{
				Name:        "force",
				Shorthand:   "f",
				Description: locale.T("flag_state_clean_uninstall_force_description"),
				Value:       &params.Force,
			},
		},
		[]*captain.Argument{},
		func(ccmd *captain.Command, _ []string) error {
			runner, err := clean.NewUninstall(prime)
			if err != nil {
				return err
			}

			return runner.Run(&params)
		},
	)
}

func newCacheCommand(prime *primer.Values) *captain.Command {
	runner := clean.NewCache(prime)
	params := clean.CacheParams{}
	return captain.NewCommand(
		"cache",
		locale.T("cache_description"),
		[]*captain.Flag{
			{
				Name:        "force",
				Shorthand:   "f",
				Description: locale.T("flag_state_clean_cache_force_description"),
				Value:       &params.Force,
			},
		},
		[]*captain.Argument{
			{
				Name:        "project",
				Description: locale.T("arg_state_clean_cache_project_description"),
				Required:    false,
				Value:       &params.Project,
			},
		},
		func(ccmd *captain.Command, _ []string) error {
			return runner.Run(&params)
		},
	)
}

func newConfigCommand(prime *primer.Values) *captain.Command {
	runner := clean.NewConfig(prime)
	params := clean.ConfigParams{}
	return captain.NewCommand(
		"config",
		locale.T("config_description"),
		[]*captain.Flag{
			{
				Name:        "force",
				Shorthand:   "f",
				Description: locale.T("flag_state_config_cache_force_description"),
				Value:       &params.Force,
			},
		},
		[]*captain.Argument{},
		func(ccmd *captain.Command, _ []string) error {
			return runner.Run(&params)
		},
	)
}
