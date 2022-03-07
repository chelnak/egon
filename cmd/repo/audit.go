package repo

import (
	"github.com/chelnak/gh-iac/internal/audit"
	"github.com/spf13/cobra"
)

// auditCmd represents the list command
var auditCmd = &cobra.Command{
	Use:     "audit audit [flags]",
	Short:   "Audit repo settings.",
	Long:    "Audit repo settings.",
	Aliases: []string{"ls"},
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
