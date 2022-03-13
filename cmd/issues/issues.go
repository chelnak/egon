package issues

import (
	"github.com/spf13/cobra"
)

var (
	repo   string
	limit  int
	dryRun bool
)

// IssuesCmd represents the issues command when called without any subcommands
var IssuesCmd = &cobra.Command{
	Use:   "issues [command]",
	Short: "Work with GitHub Issues",
	Long:  "Work with GitHub Issues",
	Run:   nil,
}

func init() {
	IssuesCmd.AddCommand(closerCmd)
	IssuesCmd.AddCommand(communityCmd)
}
