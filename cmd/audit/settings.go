package audit

import (
	"github.com/chelnak/gh-iac/internal/audit"
	"github.com/spf13/cobra"
)

// settingsCmd will return settings for a repository.
// if a repository name is not passed then it will return counts for
// all supported modules.
var settingsCmd = &cobra.Command{
	Use:   "settings [flags]",
	Short: "List settings for a repository.",
	Long:  "List settings for a repository. If a repository name is not passed then it will return counts for all supported modules.",
	RunE: func(command *cobra.Command, args []string) error {
		err := audit.GetRepositorySettings(repo)
		if err != nil {
			return err
		}
		return nil
	},
}

func init() {
	settingsCmd.Flags().StringVar(&repo, "repo", "", "The repository name.")
}
