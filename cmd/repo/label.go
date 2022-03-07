package repo

import (
	"errors"

	"github.com/chelnak/gh-iac/internal/label"
	"github.com/spf13/cobra"
)

var (
	repo   string
	issues bool
	prs    bool
)

// labelCommunityIssuesCmd represents the list command
var labelCmd = &cobra.Command{
	Use:     "label [flags]",
	Short:   "Add the community label to issues opened by contributors outside of the puppetlabs organisation.",
	Long:    "Add the community label to issues opened by contributors outside of the puppetlabs organisation.",
	Aliases: []string{"ls"},
	RunE: func(command *cobra.Command, args []string) error {

		if !issues && !prs {
			return errors.New("you must specify either --issues or --prs for this command")
		}

		if issues && prs {
			return errors.New("you cannot specify both --issues and --prs with this command")
		}

		var err error
		if issues {
			err = label.LabelCommunityIssues(repo)
		}

		if prs {
			err = label.LabelCommunityPRs(repo)
		}

		if err != nil {
			return err
		}
		return nil
	},
}

func init() {
	RepoRootCmd.AddCommand(labelCmd)

	labelCmd.Flags().StringVarP(&repo, "name", "n", "", "Repository to scan for community issues.")
	labelCmd.Flags().BoolVarP(&issues, "issues", "i", false, "Scan issues.")
	labelCmd.Flags().BoolVarP(&prs, "prs", "p", false, "Scan pull requests.")
}
