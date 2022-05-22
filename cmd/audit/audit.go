//Package audit contains commands that collect and return data.
package audit

import (
	"github.com/spf13/cobra"
)

var (
	repo string
)

// AuditCmd represents the audit command when called without any subcommands
var AuditCmd = &cobra.Command{
	Use:   "audit [flags]",
	Short: "Commands for auditing modules and tools.",
	Long:  "Commands for auditing modules and tools.",
	Run:   nil,
}

func init() {
	AuditCmd.AddCommand(issuesCmd)
}
