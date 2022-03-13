package modules

import (
	"github.com/spf13/cobra"
)

// ModulesCmd represents the modules command when called without any subcommands
var ModulesCmd = &cobra.Command{
	Use:   "modules [command]",
	Short: "Work modules that are supported by the content and tooling team.",
	Long:  "Work modules that are supported by the content and tooling team.",
	Run:   nil,
}

func init() {
	ModulesCmd.AddCommand(listCmd)
}
