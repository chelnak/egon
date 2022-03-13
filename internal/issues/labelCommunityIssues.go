package issues

import (
	"context"
	"fmt"

	"github.com/chelnak/gh-iac/internal/labels"
	"github.com/cli/go-gh"
	"github.com/google/go-github/v42/github"
)

func LabelCommunityIssues(repo string, limit int, dryRun bool) error {
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
		ListOptions: github.ListOptions{PerPage: limit},
	}
	searchResult, _, err := g.Search.Issues(ctx, fmt.Sprintf("repo:%s/%s state:open", orgName, repo), searchOpts)
	if err != nil {
		return err
	}

	for _, issue := range searchResult.Issues {
		user := issue.User.Login
		isMember, _, err := g.Organizations.IsMember(ctx, orgName, *user)
		if err != nil {
			return err
		}

		ignoredUsers := []string{"dependabot[bot]", "github-actions[bot]"}
		if !isMember && !contains(ignoredUsers, *user) {
			var s string
			if dryRun {
				s = fmt.Sprintf("Issue %v raised by \033[33m%v\033[0m is will be labelled as \033[35m%v\033[0m", *issue.Number, *issue.User.Login, labelName)
			} else {
				err := labels.AddLabelToIssue(issue, labelName, labelColor)
				if err != nil {
					return err
				}
				s = fmt.Sprintf("Issue %v raised by \033[33m%v\033[0m is now labelled as \033[35m%v\033[0m", *issue.Number, *issue.User.Login, labelName)
			}
			fmt.Println(s)
		}
	}
	return nil
}

// contains returns true if the given string is in the slice.
func contains(s []string, e string) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}
