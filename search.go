package main

import (
	"fmt"

	"github.com/cli/go-gh/v2/pkg/api"
	graphql "github.com/shurcooL/graphql"
)

// SearchResult represents the result of a search.
type SearchResult struct {
	SearchWord  string
	SearchCount int
}

// query is a GraphQL query for searching repositories.
type query struct {
	Search struct {
		RepositoryCount int
	} `graphql:"search(query: $search, type: REPOSITORY, first: 1)"`
}

// Search searches for repositories with the given search words.
func Search(searchWords []string, language string) ([]SearchResult, error) {
	client, err := api.DefaultGraphQLClient()
	if err != nil {
		return nil, fmt.Errorf("failed to create a GraphQL client: %w", err)
	}

	var results []SearchResult

	// Each search word is searched for individually.
	for _, searchWord := range searchWords {
		count, err := fetchSearchCount(client, searchWord, language)
		if err != nil {
			return nil, err
		}

		results = append(results, SearchResult{
			SearchWord:  searchWord,
			SearchCount: count,
		})
	}

	return results, nil
}

// fetchSearchCount fetches repository count for a single search word.
func fetchSearchCount(client *api.GraphQLClient, searchWord string, language string) (int, error) {
	// Add language filter if specified.
	searchQuery := searchWord
	if language != "" {
		searchQuery += fmt.Sprintf(" language:%s", language)
	}

	var q query

	variables := map[string]interface{}{
		"search": graphql.String(searchQuery),
	}

	err := client.Query("searchCount", &q, variables)
	if err != nil {
		return 0, fmt.Errorf("failed to fetch search count for '%s': %w", searchWord, err)
	}

	return q.Search.RepositoryCount, nil
}
