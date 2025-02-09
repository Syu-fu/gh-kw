package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSortSearchResults(t *testing.T) {
	tests := []struct {
		name          string
		searchResults []SearchResult
		want          []SearchResult
	}{
		{
			name:          "zero search result",
			searchResults: []SearchResult{},
			want:          []SearchResult{},
		},
		{
			name: "one search result",
			searchResults: []SearchResult{
				{SearchWord: "foo", SearchCount: 1},
			},
			want: []SearchResult{
				{SearchWord: "foo", SearchCount: 1},
			},
		},
		{
			name: "sorts search results",
			searchResults: []SearchResult{
				{SearchWord: "baz", SearchCount: 3},
				{SearchWord: "foo", SearchCount: 1},
				{SearchWord: "bar", SearchCount: 2},
			},
			want: []SearchResult{
				{SearchWord: "baz", SearchCount: 3},
				{SearchWord: "bar", SearchCount: 2},
				{SearchWord: "foo", SearchCount: 1},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := SortSearchResults(tt.searchResults); !assert.Equal(t, got, tt.want) {
				t.Errorf("SortSearchResults() = %v, want %v", got, tt.want)
			}
		})
	}
}
