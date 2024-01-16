package tui

import (
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/jedib0t/go-pretty/v6/list"
	"github.com/jedib0t/go-pretty/v6/table"
)

type PrettyTableFilter struct {
	Key        string `help:"filter target key (column)"`
	Value      string `help:"filter target value"`
	FizzyMatch bool   `default:"true" help:"enable fizzy match on value"`

	SortedBy []string `help:"sort keys"`
	Show     []string `help:"show specific columns"`
	Hide     []string `help:"hide specific columns"`
}

type PrettyTable struct {
	headers   []string
	rows      [][]string
	title     string
	autoIndex bool

	filter PrettyTableFilter
}

func (t *PrettyTable) filterRows() [][]string {
	targetIndex := indexOfHeader(t.filter.Key, t.headers)
	if targetIndex < 0 {
		return t.rows
	}

	filtered := [][]string{}
	for _, row := range t.rows {
		expect := t.filter.Value
		actual := row[targetIndex]
		if expect == actual || (t.filter.FizzyMatch && strings.Contains(actual, expect)) {
			filtered = append(filtered, row)
		}
	}
	return filtered
}

func indexOfHeader(target string, cols []string) int {
	if target == "" {
		return -1
	}

	for i, key := range cols {
		expect := strings.ToLower(target)
		actual := strings.ToLower(key)
		// equal or match the prefix
		if expect == actual || strings.HasPrefix(actual, expect) {
			return i
		}
	}

	return -1
}

func (t *PrettyTable) Print() {
	t.Write(os.Stdout)
}

func (t *PrettyTable) Write(wr io.StringWriter) {
	writer := table.NewWriter()

	// header
	h := table.Row{}
	for _, key := range t.headers {
		h = append(h, key)
	}
	writer.AppendHeader(h)

	// sorting
	sorts := []table.SortBy{}
	for _, key := range t.filter.SortedBy {
		sortIndex := indexOfHeader(key, t.headers)
		if sortIndex >= 0 {
			sorts = append(sorts, table.SortBy{
				Number: sortIndex + 1,
			})
		}
	}
	writer.SortBy(sorts)

	// visibility
	colsettings := []table.ColumnConfig{}
	if len(t.filter.Show) > 0 {
		visible := make([]bool, len(t.headers)+1)
		for _, key := range t.filter.Show {
			visible[indexOfHeader(key, t.headers)+1] = true
		}
		for i := range t.headers {
			colsettings = append(colsettings, table.ColumnConfig{
				Number: i + 1,
				Hidden: !visible[i+1],
			})
			// fmt.Printf("%s is visible %s\n", t.headers[i], !visible[i+1])
		}
	} else if len(t.filter.Hide) > 0 {
		for _, key := range t.filter.Show {
			hiddenIndex := indexOfHeader(key, t.headers)
			colsettings = append(colsettings, table.ColumnConfig{
				Number: hiddenIndex + 1,
				Hidden: true,
			})
		}
	}
	writer.SetColumnConfigs(colsettings)

	// title
	writer.SetTitle(t.title)
	writer.SetAutoIndex(t.autoIndex)

	// rows
	for _, row := range t.filterRows() {
		r := table.Row{}
		for _, value := range row {
			r = append(r, value)
		}
		writer.AppendRow(r)
	}

	wr.WriteString(writer.Render())
}

func (t *PrettyTable) SetAutoIndex(autoIndex bool) {
	t.autoIndex = autoIndex
}

func (t *PrettyTable) SetFilter(filter PrettyTableFilter) {
	t.filter = filter
}

func NewPrettyTable(title string, headers []string, rows [][]string) PrettyTable {
	return PrettyTable{
		title:   title,
		headers: headers,
		rows:    rows,
	}
}

type PrettyList struct {
	items []string
}

func (l PrettyList) Print() {
	writer := list.NewWriter()
	writer.SetStyle(list.StyleConnectedRounded)
	for _, item := range l.items {
		writer.AppendItem(item)
	}
	fmt.Println(writer.Render())
}

func NewPrettyList(items []string) PrettyList {
	return PrettyList{items}
}
