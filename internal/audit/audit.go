package audit

import (
	"context"
	"fmt"
	"strconv"

	"github.com/chelnak/gh-iac/internal/cmdutils"
	"github.com/cli/go-gh"
	"github.com/google/go-github/v42/github"
)

func GetRepositoriesForTeam(limit int) error {

	httpClient, err := gh.HTTPClient(nil)
	if err != nil {
		return err
	}

	g := github.NewClient(httpClient)

	ctx := context.Background()

	opts := github.ListOptions{
		PerPage: limit,
	}

	teamRepos, _, _ := g.Teams.ListTeamReposBySlug(ctx, "puppetlabs", "modules", &opts)

	headers := []string{"Repo", "DefaultBranch", "HasIssues", "HasProjects", "HasWiki", "HasPages", "HasDownloads", "IsArchived", "Visibility"}
	data := [][]string{}

	for _, repo := range teamRepos {

		row := []string{
			fmt.Sprintf("%s/%s", *repo.Owner.Login, *repo.Name),
			*repo.DefaultBranch,
			strconv.FormatBool(*repo.HasIssues),
			strconv.FormatBool(*repo.HasProjects),
			strconv.FormatBool(*repo.HasWiki),
			strconv.FormatBool(*repo.HasPages),
			strconv.FormatBool(*repo.HasDownloads),
			strconv.FormatBool(*repo.Archived),
			*repo.Visibility,
		}

		data = append(data, row)

	}

	table := cmdutils.NewTableWriter(headers, data, nil)

	err = table.Write()
	if err != nil {
		return err
	}

	return nil

}
