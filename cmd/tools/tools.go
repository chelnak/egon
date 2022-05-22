//Package tools contains commands for working with a subset
//of supported tools.
package tools

import (
	"github.com/spf13/cobra"
)

// ToolsCmd represents the modules command when called without any subcommands
var ToolsCmd = &cobra.Command{
	Use:   "tools [command]",
	Short: "Tools that are built and supported by the Content and Tooling team.",
	Long:  "Tools that are built and supported by the Content and Tooling team.",
	Run:   nil,
}

func init() {
	ToolsCmd.AddCommand(listCmd)
}
