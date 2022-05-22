package tools

import (
	"context"
	"fmt"

	"github.com/chelnak/purr/internal/writers"
	"github.com/olekukonko/tablewriter"
)

func ListTools() error {
	client := NewToolClient(nil)
	ctx := context.Background()
	tools, err := client.GetTools(ctx)
	if err != nil {
		return err
	}

	headers := []string{"Name", "Repo"}
	colors := []tablewriter.Colors{{tablewriter.Normal, 93}, nil}

	data := [][]string{}

	for _, tool := range *tools {
		row := []string{
			tool.Name,
			fmt.Sprintf("https://github.com/%s/%s", tool.Owner, tool.Name),
		}

		data = append(data, row)
	}

	table := writers.NewTableWriter(headers, data, colors, nil)

	return table.Write()
}
