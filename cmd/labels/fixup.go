package labels

import (
	"github.com/spf13/cobra"
)

// fixupCmd will ensure that the correct labels are set on the repo.
// If a repo name is not passed then it will check all supported module
// repositories.
var fixupCmd = &cobra.Command{
	Use:   "fixup [flags]",
	Short: "Make sure repos have the correct label set",
	Long:  "Make sure repos have the correct label set",
	RunE: func(command *cobra.Command, args []string) error {
		return nil
	},
}

func init() {
	fixupCmd.Flags().StringVar(&repo, "repo", "", "The repository name.")
}
