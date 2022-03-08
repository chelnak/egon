package label

import (
	"context"
	"fmt"

	"github.com/cli/go-gh"
	"github.com/google/go-github/v42/github"
)

func LabelCommunityIssues(repo string) error {
	labelName := "community"
	labelColor := "5319E7"
	orgName := "puppetlabs"

	httpClient, err := gh.HTTPClient(nil)
	if err != nil {
		return err
	}

	g := github.NewClient(httpClient)
	ctx := context.Background()

	searchOpts := &github.SearchOptions{
		ListOptions: github.ListOptions{PerPage: 100},
	}
	searchResult, _, err := g.Search.Issues(ctx, fmt.Sprintf("repo:%s/%s state:open", orgName, repo), searchOpts)
	if err != nil {
		return err
	}

	for _, issue := range searchResult.Issues {
		if hasLabel(issue.Labels, labelName) {
			continue
		}

		user := issue.User.Login
		isMember, _, err := g.Organizations.IsMember(ctx, orgName, *user)
		if err != nil {
			return err
		}

		if !isMember && *user != "dependabot" { // improve this
			labelName := orgName
			_, response, err := g.Issues.GetLabel(ctx, orgName, repo, labelName)
			if err != nil && response.StatusCode != 404 {
				return err
			}

			issue.Labels = append(issue.Labels, &github.Label{Name: &labelName})

			if response.StatusCode == 404 {
				_, _, err = g.Issues.CreateLabel(ctx, orgName, repo, &github.Label{
					Name:  github.String(labelName),
					Color: github.String(labelColor),
				})

				if err != nil {
					return err
				}
			}

			_, _, err = g.Issues.AddLabelsToIssue(ctx, orgName, repo, *issue.Number, []string{"community"})
			if err != nil {
				return err
			}

			s := fmt.Sprintf("Issue %v raised by \033[33m%v\033[0m is now labelled as \033[35m%v\033[0m", *issue.Number, *user, labelName)
			fmt.Println(s)
		}
	}

	return nil
}

func hasLabel(labels []*github.Label, labelName string) bool {
	m := make(map[string]*github.Label, len(labels))
	for _, label := range labels {
		m[*label.Name] = label
	}

	_, ok := m[labelName]
	return ok
}
