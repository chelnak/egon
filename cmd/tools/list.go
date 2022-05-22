package tools

import (
	"github.com/chelnak/purr/internal/tools"
	"github.com/spf13/cobra"
)

// listCmd returns the list of tools supported by the content and tooling team.
var listCmd = &cobra.Command{
	Use:     "list",
	Short:   "List tools that are supported by the Content and Tooling team.",
	Long:    "List tools that are supported by the Content and Tooling team.",
	Aliases: []string{"ls"},
	RunE: func(command *cobra.Command, args []string) error {
		err := tools.ListTools()
		if err != nil {
			return err
		}
		return nil
	},
}
