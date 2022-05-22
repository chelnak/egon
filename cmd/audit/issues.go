package audit

import (
	"fmt"
	"os"
	"time"

	"github.com/chelnak/purr/internal/audit"
	"github.com/chelnak/purr/internal/writers"
	"github.com/olekukonko/tablewriter"
	"github.com/spf13/cobra"
	"github.com/theckman/yacspin"
)

var output string
var spinner *yacspin.Spinner

// issiesCmd will return the count of issues and prs for a repository.
// if a repository name is not passed then it will return counts for
// all supported modules.
var issuesCmd = &cobra.Command{
	Use:   "issues [flags]",
	Short: "Get the count of open issues and prs for a repository.",
	Long:  "Get the count of open issues and prs for a repository. If a repository name is not passed then it will return counts for all supported modules.",
	RunE: func(command *cobra.Command, args []string) error {
		var writer writers.Writer
		setupSpinner()

		data, err := audit.GetIssueAndPRCount(repo)
		if err != nil {
			return err
		}

		headers := []string{"repository", "issues", "pull_requests"}

		switch output {
		case "table":
			colors := []tablewriter.Colors{{tablewriter.Normal, 93}, nil, nil}
			writer = writers.NewTableWriter(headers, data, colors, nil)

		case "csv":
			writer = writers.NewCsvWriter(headers, data, os.Stdout)

		default:
			return fmt.Errorf("invalid output format: %s", output)
		}

		_ = spinner.Stop()
		return writer.Write()
	},
}

func init() {
	issuesCmd.Flags().StringVar(&repo, "repo", "", "The repository name.")
	issuesCmd.Flags().StringVar(&output, "output", "table", "Output format. One of: table, json, csv.")
}

func setupSpinner() {
	cfg := yacspin.Config{
		Frequency:       100 * time.Millisecond,
		CharSet:         yacspin.CharSets[11],
		Colors:          []string{"fgGreen"},
		SuffixAutoColon: true,
	}

	spinner, _ = yacspin.New(cfg)
	_ = spinner.Start()
}
