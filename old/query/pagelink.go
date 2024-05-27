package query

import (
	"fmt"
	"wiki-graph/old/client"
)

const (
	BASE_URL   = "https://en.wikipedia.org/w/api.php"
	ACTION     = "query"
	PROPERTY   = "links"
	FORMAT     = "json"
	PAGE_TYPE  = "0"
	BATCH_SIZE = "max"
)

type PageLinkQuery struct {
	pageTitle   string
	nextResult  string
	batchSize   string
	maxRequests int
}

// OptionFunc is a type for functions that configure a PageLinkQuery.
type OptionFunc func(*PageLinkQuery)

// WithBatchSize configures the query's batch size.
func WithBatchSize(size int) OptionFunc {
	return func(query *PageLinkQuery) {
		query.batchSize = string(rune(size))
	}
}

// WithNextResult configures the query's value for nextResult for pagination.
func WithNextResult(nextResult string) OptionFunc {
	return func(query *PageLinkQuery) {
		query.nextResult = nextResult
	}
}

func NewPageLinkQuery(pageTitle string, options ...OptionFunc) *PageLinkQuery {
	// Initialize new PageLinkQuery with defaults
	query := &PageLinkQuery{
		pageTitle: pageTitle,
		batchSize: BATCH_SIZE,
	}
	// Update additional values with options
	for _, option := range options {
		option(query)
	}
	return query
}

func (query *PageLinkQuery) Fetch(client client.Client) (PageLinksResponse, error) {
	// Set the default query params
	params := map[string]string{
		"action":  ACTION,
		"prop":    PROPERTY,
		"format":  FORMAT,
		"ns":      PAGE_TYPE,
		"pllimit": query.batchSize,
		"titles":  query.pageTitle,
	}
	// If nextResult is set, add it to the query params as 'plcontinue'
	// This param tells the wikipedia API the next result to retrieve
	if query.nextResult != "" {
		params["plcontinue"] = query.nextResult
	}
	// Make the API call
	response, err := client.Get(BASE_URL, params)
	if err != nil {
		err = fmt.Errorf("error making the API request: %s", err)
		return PageLinksResponse{}, err
	}
	// Parse the API response
	parsedResponse, err := PageLinksResponse{}.Parse(response)
	if err != nil {
		err = fmt.Errorf("error parsing API response: %s", err)
		return PageLinksResponse{}, err
	}
	return *parsedResponse, nil
}

func (query *PageLinkQuery) FetchAll(client client.Client) []Page {
	var pageLinks = []Page{}
	for requestCount := 0; requestCount < query.maxRequests; requestCount++ {
		// Make the request and log any errors
		result, err := query.Fetch(client)
		if err != nil {
			fmt.Printf("Error with request number %d: %s\n", requestCount, err)
			fmt.Println("  Page title:", query.pageTitle)
			fmt.Println("  Next result param:", query.pageTitle)
			continue
		}
		// Append the page links returned by the API to the list of outbound links
		if links := result.ExtractLinks(); len(links) > 0 {
			pageLinks = append(pageLinks, links...)
		}
		// If the batch is complete return the list of page links
		if result.BatchComplete == nil {
			return pageLinks
		}
		// Otherwise update the nextResultCursor and continue
		query.nextResult = result.NextResult.Cursor
	}
	return pageLinks
}
