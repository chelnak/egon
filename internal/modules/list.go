package modules

import (
	"context"
	"fmt"

	"github.com/chelnak/gh-iac/internal/cmdutils"
	"github.com/olekukonko/tablewriter"
)

func ListModules() error {
	client := NewModuleClient(nil)
	ctx := context.Background()
	modules, err := client.GetSupportedModules(ctx)
	if err != nil {
		return err
	}

	headers := []string{"Name", "Repo"}
	colors := []tablewriter.Colors{{tablewriter.Normal, 93}, nil}

	data := [][]string{}

	for _, module := range *modules {
		row := []string{
			module.Name,
			fmt.Sprintf("https://github.com/puppetlabs/%s", module.Repo),
		}

		data = append(data, row)
	}

	table := cmdutils.NewTableWriter(headers, data, colors, nil)

	err = table.Write()
	if err != nil {
		return err
	}

	return nil
}
