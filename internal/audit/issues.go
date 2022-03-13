package audit

import (
	"context"
	"strconv"
	"time"

	"github.com/chelnak/gh-iac/internal/cmdutils"
	"github.com/chelnak/gh-iac/internal/modules"
	"github.com/cli/go-gh"
	"github.com/google/go-github/v42/github"
	"github.com/olekukonko/tablewriter"
	"github.com/theckman/yacspin"
)

func GetIssueAndPRCount(repo string) error {
	owner := "puppetlabs"

	httpClient, err := gh.HTTPClient(nil)
	if err != nil {
		return err
	}

	g := github.NewClient(httpClient)
	ctx := context.Background()

	var m []modules.Module
	if repo == "" {
		modulesClient := modules.NewModuleClient(nil)
		supportedModules, err := modulesClient.GetSupportedModules(ctx)
		if err != nil {
			return err
		}
		m = *supportedModules
	} else {
		m = make([]modules.Module, 1)
		m[0] = modules.Module{Name: repo}
	}

	headers := []string{"Repo", "IssueCount", "PRCount"}
	colors := []tablewriter.Colors{{tablewriter.Normal, 93}, nil, nil}
	data := [][]string{}

	cfg := yacspin.Config{
		Frequency:       100 * time.Millisecond,
		CharSet:         yacspin.CharSets[11],
		Colors:          []string{"fgGreen"},
		SuffixAutoColon: true,
	}

	spinner, err := yacspin.New(cfg)
	if err != nil {
		return err
	}

	err = spinner.Start()
	if err != nil {
		return err
	}

	for _, module := range m {
		listRepoOpts := github.IssueListByRepoOptions{
			ListOptions: github.ListOptions{
				PerPage: 100,
			},
		}

		issues, _, err := g.Issues.ListByRepo(ctx, owner, module.Name, &listRepoOpts)
		if err != nil {
			return err
		}

		listPROpts := github.PullRequestListOptions{
			ListOptions: github.ListOptions{
				PerPage: 100,
			},
		}

		prs, _, err := g.PullRequests.List(ctx, owner, module.Name, &listPROpts)
		if err != nil {
			return err
		}

		issueCount := len(issues)
		prCount := len(prs)

		row := []string{
			module.Name,
			strconv.Itoa(issueCount - prCount), // issues seems to be inclusive of issues and prs, so we need to subtract prs
			strconv.Itoa(prCount),
		}

		data = append(data, row)
	}

	err = spinner.Stop()
	if err != nil {
		return err
	}

	table := cmdutils.NewTableWriter(headers, data, colors, nil)

	err = table.Write()
	if err != nil {
		return err
	}

	return nil
}
