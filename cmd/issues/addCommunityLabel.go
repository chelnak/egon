package issues

import (
	"github.com/chelnak/gh-iac/internal/issues"
	"github.com/spf13/cobra"
)

// communityCmd will add the label 'community to any issues and prs that have been opened by external contributors.
// Passing dry-run will not actually add the label. This is useful for testing.
var communityCmd = &cobra.Command{
	Use:   "add-community-label [flags]",
	Short: "Identify issues and prs that have been opened by external contributors and assign the community label.",
	Long:  "Identify issues and prs that have been opened by external contributors and assign the community label.",
	RunE: func(command *cobra.Command, args []string) error {
		err := issues.LabelCommunityIssues(repo, limit, dryRun)
		if err != nil {
			return err
		}
		return nil
	},
}

func init() {
	communityCmd.Flags().StringVar(&repo, "repo", "", "Repository to scan for community issues.")
	communityCmd.Flags().IntVar(&limit, "limit", 100, "Limit the number of issues to label.")
	communityCmd.Flags().BoolVar(&dryRun, "dry-run", false, "Performs a no-op run.")

	_ = cobra.MarkFlagRequired(communityCmd.Flags(), "repo")
}
