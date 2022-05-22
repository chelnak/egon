package writers_test

import (
	"bytes"
	"testing"

	"github.com/chelnak/purr/internal/writers"
	"github.com/stretchr/testify/assert"
)

func TestTableWriter(t *testing.T) {
	tests := []struct {
		name    string
		headers []string
		rows    [][]string
		want    string
	}{
		{
			name:    "with empty headers",
			headers: []string{"", "", "", ""},
			rows: [][]string{
				{"foo", "bar", "baz", "qux"},
				{"baz", "qux", "foo", "bar"},
			},
			want: "                      \n──────────────────────\nfoo   bar   baz   qux   \nbaz   qux   foo   bar   \n",
		},
		{
			name:    "with populated headers",
			headers: []string{"foo", "bar", "baz", "qux"},
			rows: [][]string{
				{"foo", "bar", "baz", "qux"},
				{"baz", "qux", "foo", "bar"},
			},
			want: "foo   bar   baz   qux \n──────────────────────\nfoo   bar   baz   qux   \nbaz   qux   foo   bar   \n",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var output bytes.Buffer
			tw := writers.NewTableWriter(tt.headers, tt.rows, nil, &output)
			err := tw.Write()
			assert.NoError(t, err)
			assert.Equal(t, tt.want, output.String())
		})
	}
}
