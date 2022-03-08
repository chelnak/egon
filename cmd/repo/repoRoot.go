package repo

import (
	"github.com/spf13/cobra"
)

var (
	repo  string
	limit int
)

// RepoRootCmd represents the base command when called without any subcommands
var RepoRootCmd = &cobra.Command{
	Use:   "repo [command]",
	Short: "Work with Puppet GitHub Repoisitories",
	Long:  "Work with Puppet GitHub Repoisitories",
	Run:   nil,
}
