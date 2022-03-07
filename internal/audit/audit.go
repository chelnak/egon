package audit

import (
	"context"
	"strconv"
	"strings"
	"time"

	"github.com/chelnak/gh-iac/internal/cmdutils"
	"github.com/chelnak/gh-iac/internal/modules"
	"github.com/cli/go-gh"
	"github.com/google/go-github/v42/github"
	"github.com/olekukonko/tablewriter"
	"github.com/theckman/yacspin"
)

func ListRepositorySettings() error {
	httpClient, err := gh.HTTPClient(nil)
	if err != nil {
		return err
	}

	m := modules.NewModuleClient(nil)
	g := github.NewClient(httpClient)
	ctx := context.Background()

	modules, err := m.GetSupportedModules(ctx)
	if err != nil {
		return err
	}

	headers := []string{"Repo", "DefaultBranch", "HasIssues", "HasProjects", "HasWiki", "HasPages", "HasDownloads", "IsArchived", "DeleteHead", "IssueCount", "PRCount"}
	colors := []tablewriter.Colors{{tablewriter.Normal, 93}, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil}
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

	for _, module := range *modules {

		s := strings.Split(module.Repo, "/")
		entity, _, err := g.Repositories.Get(ctx, s[0], s[1])
		if err != nil {
			return err
		}

		listRepoOpts := github.IssueListByRepoOptions{
			ListOptions: github.ListOptions{
				PerPage: 100,
			},
		}
		issues, _, err := g.Issues.ListByRepo(ctx, s[0], s[1], &listRepoOpts)
		if err != nil {
			return err
		}

		listPROpts := github.PullRequestListOptions{
			ListOptions: github.ListOptions{
				PerPage: 100,
			},
		}
		prs, _, err := g.PullRequests.List(ctx, s[0], s[1], &listPROpts)
		if err != nil {
			return err
		}

		row := []string{
			module.Repo,
			*entity.DefaultBranch,
			strconv.FormatBool(*entity.HasIssues),
			strconv.FormatBool(*entity.HasProjects),
			strconv.FormatBool(*entity.HasWiki),
			strconv.FormatBool(*entity.HasPages),
			strconv.FormatBool(*entity.HasDownloads),
			strconv.FormatBool(*entity.Archived),
			strconv.FormatBool(entity.GetDeleteBranchOnMerge()),
			strconv.Itoa(len(issues)),
			strconv.Itoa(len(prs)),
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
