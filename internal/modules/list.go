package modules

import (
	"context"
	"fmt"

	"github.com/chelnak/purr/internal/writers"
	"github.com/olekukonko/tablewriter"
)

func ListModules() error {
	client := NewModuleClient(nil)
	ctx := context.Background()
	modules, err := client.GetSupportedModules(ctx)
	if err != nil {
		return err
	}

	headers := []string{"name", "repository"}
	colors := []tablewriter.Colors{{tablewriter.Normal, 93}, nil}

	data := [][]string{}

	for _, module := range *modules {
		row := []string{
			module.Name,
			fmt.Sprintf("https://github.com/%s/%s", module.Owner, module.Name),
		}

		data = append(data, row)
	}

	table := writers.NewTableWriter(headers, data, colors, nil)

	return table.Write()
}
