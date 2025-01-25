package main

import (
	"io"
	"strconv"

	"github.com/olekukonko/tablewriter"
)

func Output(searchResults []SearchResult, outstream io.Writer) {
	// markdown table format
	data := make([][]string, len(searchResults))
	for i, result := range searchResults {
		data[i] = []string{strconv.Itoa(i + 1), result.searchWord, strconv.Itoa(result.searchCount)}
	}

	tw := tablewriter.NewWriter(outstream)
	tw.SetHeader([]string{"Rank", "Keyword", "Total"})
	tw.SetBorders(tablewriter.Border{Left: true, Top: false, Right: true, Bottom: false})
	tw.SetCenterSeparator("|")
	tw.AppendBulk(data)
	tw.Render()
}
