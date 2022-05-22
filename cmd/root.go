//Package cmd contains top level cli commands.
package cmd

import (
	"errors"
	"fmt"
	"os"

	"github.com/chelnak/purr/cmd/audit"
	"github.com/chelnak/purr/cmd/labels"
	"github.com/chelnak/purr/cmd/modules"
	"github.com/chelnak/purr/cmd/tools"
	"github.com/spf13/cobra"
)

var version = "dev"
var ErrSilent = errors.New("ErrSilent")

// RootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:           "purr [command]",
	Short:         "A utility tool belt for the Content and Tooling team",
	Long:          "A utility tool belt for the Content and Tooling team",
	Version:       version,
	SilenceErrors: true,
	SilenceUsage:  true,
	Run:           nil,
}

func init() {
	rootCmd.SetFlagErrorFunc(func(cmd *cobra.Command, err error) error {
		cmd.Println(err)
		cmd.Println(cmd.UsageString())
		return ErrSilent
	})

	rootCmd.AddCommand(modules.ModulesCmd)
	rootCmd.AddCommand(tools.ToolsCmd)
	rootCmd.AddCommand(audit.AuditCmd)
	rootCmd.AddCommand(labels.LabelsCmd)
}

func Execute() int {
	if err := rootCmd.Execute(); err != nil {
		if err != ErrSilent {
			fmt.Fprintln(os.Stderr, err)
		}
		return 1
	}
	return 0
}
