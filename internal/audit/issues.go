// Package audit is responsible for gathering data from the requested source
// and returning to the caller in a standardised way.
package audit

import (
	"context"
	"strconv"

	"github.com/chelnak/purr/internal/modules"
	"github.com/cli/go-gh"
	"github.com/google/go-github/v42/github"
)

func GetIssueAndPRCount(repo string) ([][]string, error) {
	owner := "puppetlabs"

	httpClient, err := gh.HTTPClient(nil)
	if err != nil {
		return nil, err
	}

	g := github.NewClient(httpClient)
	ctx := context.Background()

	var m []modules.Module
	if repo == "" {
		modulesClient := modules.NewModuleClient(nil)
		supportedModules, err := modulesClient.GetSupportedModules(ctx)
		if err != nil {
			return nil, err
		}

		m = append(m, *supportedModules...)
	} else {
		m = append(m, modules.Module{Name: repo})
	}

	data := [][]string{}

	for _, module := range m {
		listRepoOpts := github.IssueListByRepoOptions{
			ListOptions: github.ListOptions{
				PerPage: 100,
			},
		}

		issues, _, err := g.Issues.ListByRepo(ctx, owner, module.Name, &listRepoOpts)
		if err != nil {
			return nil, err
		}

		listPROpts := github.PullRequestListOptions{
			ListOptions: github.ListOptions{
				PerPage: 100,
			},
		}

		prs, _, err := g.PullRequests.List(ctx, owner, module.Name, &listPROpts)
		if err != nil {
			return nil, err
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

	return data, nil
}
