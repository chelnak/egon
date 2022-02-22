package cmd

import (
	"github.com/chelnak/gh-iac/internal/audit"
	"github.com/spf13/cobra"
)

var (
	limit        int
	outputAsJSON bool
	query        string
)

// listCmd represents the list command
var auditCmd = &cobra.Command{
	Use:     "audit",
	Short:   "Audit IAC repositories",
	Long:    "Audit IAC repositories",
	Aliases: []string{"ls"},
	RunE: func(command *cobra.Command, args []string) error {

		audit.GetRepositoriesForTeam(limit)

		return nil
	},
}

func init() {
	rootCmd.AddCommand(auditCmd)
	auditCmd.Flags().IntVarP(&limit, "limit", "l", 30, "the number of environments to show per page")
	auditCmd.Flags().BoolVarP(&outputAsJSON, "json", "j", false, "Output in JSON format")
	auditCmd.Flags().StringVarP(&query, "query", "q", "", "a query string to filter environments")
}
