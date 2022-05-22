package writers

import (
	"encoding/csv"
	"io"
)

type csvWriter struct {
	headers []string
	rows    [][]string
	writer  io.Writer
}

func (c *csvWriter) Write() error {
	data := [][]string{c.headers}
	data = append(data, c.rows...)
	writer := csv.NewWriter(c.writer)
	return writer.WriteAll(data)
}

func NewCsvWriter(headers []string, rows [][]string, writer io.Writer) Writer {
	return &csvWriter{headers, rows, writer}
}
