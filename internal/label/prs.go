package label

import (
	"context"
	"fmt"

	"github.com/cli/go-gh"
	"github.com/google/go-github/v42/github"
)

func LabelCommunityPRs(repo string) error {
	httpClient, err := gh.HTTPClient(nil)
	if err != nil {
		return err
	}

	g := github.NewClient(httpClient)
	ctx := context.Background()

	issueListOptions := &github.PullRequestListOptions{
		ListOptions: github.ListOptions{
			PerPage: 100,
		},
	}
	pullRequests, _, err := g.PullRequests.List(ctx, "puppetlabs", repo, issueListOptions)
	if err != nil {
		return err
	}

	for _, pr := range pullRequests {
		user := pr.User.Login

		isMember, _, err := g.Organizations.IsMember(ctx, "puppetlabs", *user)
		if err != nil {
			return err
		}

		if !isMember {

			labelName := "community"
			_, response, err := g.Issues.GetLabel(ctx, "puppetlabs", repo, labelName)

			if err != nil && response.StatusCode != 404 {
				return err
			}

			if response.StatusCode == 404 {
				_, _, err = g.Issues.CreateLabel(ctx, "puppetlabs", repo, &github.Label{
					Name:  github.String(labelName),
					Color: github.String("5319E7"),
				})

				if err != nil {
					return err
				}
			}

			_, _, err = g.Issues.AddLabelsToIssue(ctx, "puppetlabs", repo, *pr.Number, []string{"community"})
			if err != nil {
				return err
			}

			fmt.Println("PR", *pr.Number, "raised by user", string("\033[33m"), *user, string("\033[0m"), "has been labeled as", string("\033[35m"), "community", string("\033[0m"))
		}
	}

	return nil
}
