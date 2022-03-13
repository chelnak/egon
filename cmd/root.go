package cmd

import (
	"errors"
	"fmt"
	"os"

	"github.com/chelnak/gh-iac/cmd/audit"
	"github.com/chelnak/gh-iac/cmd/issues"
	"github.com/chelnak/gh-iac/cmd/labels"
	"github.com/chelnak/gh-iac/cmd/modules"
	"github.com/spf13/cobra"
)

var version = "dev"
var ErrSilent = errors.New("ErrSilent")

// RootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:           "egon [command]",
	Short:         "Work with Puppet GitHub Repoisitories",
	Long:          "Work with Puppet GitHub Repoisitories",
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
	rootCmd.AddCommand(audit.AuditCmd)
	rootCmd.AddCommand(issues.IssuesCmd)
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
