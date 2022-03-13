package audit

import (
	"github.com/chelnak/gh-iac/internal/audit"
	"github.com/spf13/cobra"
)

// issiesCmd will return the count of issues and prs for a repository.
// if a repository name is not passed then it will return counts for
// all supported modules.
var issuesCmd = &cobra.Command{
	Use:   "issues [flags]",
	Short: "Get the count of open issues and prs for a repository.",
	Long:  "Get the count of open issues and prs for a repository. If a repository name is not passed then it will return counts for all supported modules.",
	RunE: func(command *cobra.Command, args []string) error {
		err := audit.GetIssueAndPRCount(repo)
		if err != nil {
			return err
		}
		return nil
	},
}

func init() {
	issuesCmd.Flags().StringVar(&repo, "repo", "", "The repository name.")
}
