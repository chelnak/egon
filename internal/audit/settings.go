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

func GetRepositorySettings(repo string) error {
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

	headers := []string{"Repo", "DefaultBranch", "HasIssues", "HasProjects", "HasWiki", "HasPages", "HasDownloads", "IsArchived", "DeleteHead"}
	colors := []tablewriter.Colors{{tablewriter.Normal, 93}, nil, nil, nil, nil, nil, nil, nil, nil}
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
		entity, _, err := g.Repositories.Get(ctx, owner, module.Name)
		if err != nil {
			return err
		}

		row := []string{
			module.Name,
			*entity.DefaultBranch,
			strconv.FormatBool(entity.GetHasIssues()),
			strconv.FormatBool(entity.GetHasProjects()),
			strconv.FormatBool(entity.GetHasWiki()),
			strconv.FormatBool(entity.GetHasPages()),
			strconv.FormatBool(entity.GetHasDownloads()),
			strconv.FormatBool(entity.GetArchived()),
			strconv.FormatBool(entity.GetDeleteBranchOnMerge()),
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
