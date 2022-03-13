package labels

import (
	"github.com/spf13/cobra"
)

var (
	repo string
)

// LabelsCmd represents the labels command when called without any subcommands
var LabelsCmd = &cobra.Command{
	Use:   "labels [command]",
	Short: "Work with Puppet GitHub Repoisitories",
	Long:  "Work with Puppet GitHub Repoisitories",
	Run:   nil,
}

func init() {
	LabelsCmd.AddCommand(fixupCmd)
}
