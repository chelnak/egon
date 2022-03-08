package issue

import (
	"context"
	"fmt"

	"github.com/cli/go-gh"
	"github.com/google/go-github/v42/github"
)

func CloseIssuesOlderThan(repoName, date string, close bool) error {
	ctx := context.Background()
	repoOwner := "puppetlabs"

	httpClient, err := gh.HTTPClient(nil)
	if err != nil {
		return err
	}
	client := github.NewClient(httpClient)

	issues, _, err := client.Search.Issues(ctx, fmt.Sprintf("repo:%v/%v updated:<=%v type:issue state:open", repoOwner, repoName, date), &github.SearchOptions{ListOptions: github.ListOptions{PerPage: 100}})
	if err != nil {
		return err
	}

	fmt.Printf("Total: %v\n", *issues.Total)
	for i, iss := range issues.Issues {
		fmt.Printf("Issue %v: %v - %v \n", i+1, *iss.Title, iss.UpdatedAt.Format("2006-01-02"))
	}

	if close {
		for _, issue := range issues.Issues {
			body := "Hello! We are doing some house keeping and noticed that this issue has been open for a long time.\n\nWe're going to close it but please do raise another issue if the issue still persists. 😄"
			_, _, err := client.Issues.CreateComment(ctx, repoOwner, repoName, *issue.Number, &github.IssueComment{Body: &body})
			if err != nil {
				return err
			}

			_, _, err = client.Issues.Edit(ctx, repoOwner, repoName, *issue.Number, &github.IssueRequest{State: github.String("closed")})
			if err != nil {
				return err
			}
		}
		fmt.Printf("%v issues closed in the %v repository", *issues.Total, repoName)
	}

	return nil
}
