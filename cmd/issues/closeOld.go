package issues

import (
	"github.com/chelnak/gh-iac/internal/issues"
	"github.com/spf13/cobra"
)

var (
	date string
)

// closerCmd will close issues that were opened before a given date.
// Passing dry-run will not actually close the issues. This is useful
// for testing.
var closerCmd = &cobra.Command{
	Use:   "close-old [flags]",
	Short: "Query for and close issues for a given date.",
	Long:  "Query for and close issues for a given date.",
	RunE: func(command *cobra.Command, args []string) error {
		err := issues.CloseIssuesOlderThan(repo, date, dryRun, limit)
		if err != nil {
			return err
		}
		return nil
	},
}

func init() {
	closerCmd.Flags().StringVar(&repo, "repo", "", "Name of the repository.")
	closerCmd.Flags().StringVar(&date, "date", "2017-01-01", "Returns issues before the specified date. Date format: 'yyyy-MM-dd' (Default: '2017-01-01')")
	closerCmd.Flags().BoolVar(&dryRun, "dry-run", false, "Performs a no-op run.")
	closerCmd.Flags().IntVar(&limit, "limit", 100, "Limit the number of issues to label.")

	_ = cobra.MarkFlagRequired(closerCmd.Flags(), "name")
}
