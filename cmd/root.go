package cmd

import (
	"errors"
	"fmt"
	"os"

	"github.com/chelnak/gh-iac/cmd/repo"
	"github.com/spf13/cobra"
)

var version = "dev"
var ErrSilent = errors.New("ErrSilent")

// RootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:           "ppt [command]",
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

	rootCmd.AddCommand(repo.RepoRootCmd)
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
