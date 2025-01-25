package main

import (
	"fmt"

	"github.com/cli/go-gh/v2/pkg/api"
	graphql "github.com/shurcooL/graphql"
)

// SearchResult represents the result of a search.
type SearchResult struct {
	searchWord  string
	searchCount int
}

// query is a GraphQL query for searching repositories.
type query struct {
	Search struct {
		RepositoryCount int
	} `graphql:"search(query: $search, type: REPOSITORY, first:100)"`
}

// Search searches for repositories with the given search words.
func Search(searchWords []string) ([]SearchResult, error) {
	client, err := api.DefaultGraphQLClient()
	if err != nil {
		return nil, fmt.Errorf("failed to create a GraphQL client: %w", err)
	}

	return fetchSearchResults(client, searchWords)
}

// fetchSearchResults fetches search results from the GitHub API.
func fetchSearchResults(client *api.GraphQLClient, searchWords []string) ([]SearchResult, error) {
	var results []SearchResult

	for _, searchWord := range searchWords {
		variables := map[string]interface{}{
			"search": graphql.String(searchWord),
		}

		var q query

		err := client.Query("searchCount", &q, variables)
		if err != nil {
			return nil, fmt.Errorf("failed to search for %s: %w", searchWord, err)
		}

		results = append(results, SearchResult{
			searchWord:  searchWord,
			searchCount: q.Search.RepositoryCount,
		})
	}

	return results, nil
}
