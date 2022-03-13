package modules

import (
	"github.com/chelnak/gh-iac/internal/modules"
	"github.com/spf13/cobra"
)

// listCmd returns the list of modules supported by the content and tooling team.
var listCmd = &cobra.Command{
	Use:     "list",
	Short:   "List modules that are supported by the modules team.",
	Long:    "List modules that are supported by the modules team.",
	Aliases: []string{"ls"},
	RunE: func(command *cobra.Command, args []string) error {
		err := modules.ListModules()
		if err != nil {
			return err
		}
		return nil
	},
}
