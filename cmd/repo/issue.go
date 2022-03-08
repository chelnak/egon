package repo

import (
	"github.com/chelnak/gh-iac/internal/issue"
	"github.com/spf13/cobra"
)

var (
	date  string
	close bool
	name  string
)

var issueCmd = &cobra.Command{
	Use:   "issue [flags]",
	Short: "Query for and optionally close issues for a given date.",
	Long:  "Query for and optionally close issues for a given date.",
	RunE: func(command *cobra.Command, args []string) error {
		err := issue.CloseIssuesOlderThan(name, date, close)
		if err != nil {
			return err
		}
		return nil
	},
}

func init() {
	RepoRootCmd.AddCommand(issueCmd)

	issueCmd.Flags().StringVarP(&name, "name", "n", "", "Name of the repository.")
	issueCmd.Flags().StringVarP(&date, "date", "d", "2017-01-01", "Returns issues before the specified date. Date format: 'yyyy-MM-dd' (Default: '2017-01-01')")
	issueCmd.Flags().BoolVarP(&close, "close", "c", false, "Closes issues returned by the query.")
}
