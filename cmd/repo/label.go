package repo

import (
	"github.com/chelnak/gh-iac/internal/label"
	"github.com/spf13/cobra"
)

// labelCmd represents a command that will scan a repo for issues raised by external contributors
// and label them with the community label.
var labelCmd = &cobra.Command{
	Use:   "label [flags]",
	Short: "Add the community label to issues and prs opened by contributors outside of the puppetlabs organisation.",
	Long:  "Add the community label to issues and prs opened by contributors outside of the puppetlabs organisation.",
	RunE: func(command *cobra.Command, args []string) error {
		err := label.LabelCommunityIssues(repo, limit)
		if err != nil {
			return err
		}
		return nil
	},
}

func init() {
	RepoRootCmd.AddCommand(labelCmd)
	labelCmd.Flags().StringVarP(&repo, "name", "n", "", "Repository to scan for community issues.")
	labelCmd.Flags().IntVarP(&limit, "limit", "l", 100, "Limit the number of issues to label.")
}
