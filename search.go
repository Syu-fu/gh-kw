package main

import (
	"fmt"
	"strings"

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
	} `graphql:"search(query: $search, type: REPOSITORY, first: 100)"`
}

// Search searches for repositories with the given search words.
func Search(searchWords []string, language string) ([]SearchResult, error) {
	client, err := api.DefaultGraphQLClient()
	if err != nil {
		return nil, fmt.Errorf("failed to create a GraphQL client: %w", err)
	}

	return fetchSearchResults(client, searchWords, language)
}

// fetchSearchResults fetches search results from the GitHub API.
func fetchSearchResults(client *api.GraphQLClient, searchWords []string, language string) ([]SearchResult, error) {
	var results []SearchResult

	// Construct the search query, incorporating language as a filter.
	joinedSearchWords := strings.Join(searchWords, " ")
	searchQuery := joinedSearchWords

	if language != "" {
		searchQuery += fmt.Sprintf(" language:%s", language)
	}

	variables := map[string]interface{}{
		"search": graphql.String(searchQuery),
	}

	var q query

	err := client.Query("searchCount", &q, variables)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch search counts: %w", err)
	}

	// Process each search word and map the results.
	for _, searchWord := range searchWords {
		results = append(results, SearchResult{
			SearchWord:  searchWord,
			SearchCount: q.Search.RepositoryCount,
		})
	}

	return results, nil
}
