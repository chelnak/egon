package labels

import (
	"context"
	"fmt"

	"github.com/cli/go-gh"
	"github.com/google/go-github/v42/github"
)

func AddLabelToIssue(issue *github.Issue, labelName string, labelColor string) error {
	if hasLabel(issue.Labels, labelName) {
		fmt.Printf("This issue already has the %v label\n", labelName)
		return nil
	}

	httpClient, err := gh.HTTPClient(nil)
	if err != nil {
		return err
	}

	g := github.NewClient(httpClient)
	ctx := context.Background()

	issue.Labels = append(issue.Labels, &github.Label{Name: &labelName})

	repo := issue.GetRepository()
	_, response, err := g.Issues.GetLabel(ctx, *repo.Owner.Login, *repo.Name, labelName)
	if err != nil && response.StatusCode != 404 {
		return err
	}

	if response.StatusCode == 404 {
		_, _, err = g.Issues.CreateLabel(ctx, *repo.Owner.Login, *repo.Name, &github.Label{
			Name:  github.String(labelName),
			Color: github.String(labelColor),
		})

		if err != nil {
			return err
		}
	}

	_, _, err = g.Issues.AddLabelsToIssue(ctx, *repo.Owner.Login, *repo.Name, *issue.Number, []string{labelName})
	if err != nil {
		return err
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
