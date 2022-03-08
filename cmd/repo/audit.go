package repo

import (
	"github.com/chelnak/gh-iac/internal/audit"
	"github.com/spf13/cobra"
)

// auditCmd represents the audit command
var auditCmd = &cobra.Command{
	Use:   "audit [flags]",
	Short: "Audit repo settings.",
	Long:  "Audit repo settings.",
	RunE: func(command *cobra.Command, args []string) error {
		err := audit.ListRepositorySettings()
		if err != nil {
			return err
		}
		return nil
	},
}

func init() {
	RepoRootCmd.AddCommand(auditCmd)
}
