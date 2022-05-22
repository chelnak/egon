package writers

import (
	"io"
	"os"

	"github.com/olekukonko/tablewriter"
)

type tableWriter struct {
	headers []string
	data    [][]string
	colors  []tablewriter.Colors
	writer  io.Writer
}

func NewTableWriter(headers []string, data [][]string, colors []tablewriter.Colors, writer io.Writer) Writer {
	return tableWriter{
		headers: headers,
		data:    data,
		colors:  colors,
		writer:  writer,
	}
}

func (t tableWriter) Write() error {
	if t.writer == nil {
		t.writer = io.Writer(os.Stdout)
	}

	table := tablewriter.NewWriter(t.writer)
	table.AppendBulk(t.data)

	table.SetHeader(t.headers)
	table.SetBorder(false)
	//table.SetRowLine(true)
	table.SetHeaderLine(true)
	table.SetRowSeparator("─")
	table.SetCenterSeparator("")
	table.SetHeaderAlignment(tablewriter.ALIGN_LEFT)
	table.SetAutoFormatHeaders(false)
	table.SetTablePadding("   ") // two spaces
	table.SetNoWhiteSpace(true)
	table.SetAutoWrapText(true)
	table.SetReflowDuringAutoWrap(true)

	if t.colors != nil {
		table.SetColumnColor(t.colors...)
	}

	table.Render()

	return nil
}
