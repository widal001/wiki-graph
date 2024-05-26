package graph

import (
	"fmt"
	"testing"
	"wiki-graph/api"
	"wiki-graph/query"
)

var MockLinks = map[string][]string{
	"a": {"b", "c", "d"},
	"b": {"a", "d", "e", "f"},
	"c": {"a", "f"},
	"d": {"a", "b"},
	"e": {"a", "g"},
	"f": {"b", "c", "d"},
	"g": {"a", "d", "e"},
}

type mockWikiApi struct {
	linkMap map[string][]string
}

func createMockApi() mockWikiApi {
	api := mockWikiApi{linkMap: make(map[string][]string)}
	for page, links := range MockLinks {
		api.linkMap[page] = links
	}
	return api
}

func (a *mockWikiApi) FetchOutboundPageLinks(page string) (api.PageLinks, error) {
	// Try to fetch the links for a given page
	links, found := a.linkMap[page]
	// Raise an error if the page was already accessed
	if !found {
		err := fmt.Errorf("%s was accessed multiple times", page)
		return api.PageLinks{}, err
	}
	// Property format the links to match the output of the wikipedia API
	var pageLinks []query.Page
	for _, link := range links {
		pageLinks = append(pageLinks, query.Page{
			PageId:    1,
			Title:     link,
			Namespace: 0,
			Summary:   "foo",
		})
	}
	// Remove that page from the map, so that subsequent requests fail
	delete(a.linkMap, page)
	// return the page links
	return api.PageLinks{Page: query.Page{Title: page}, Links: pageLinks}, nil
}

func TestGraph(t *testing.T) {
	t.Run("Mock function raises error if page is requested multiple times.", func(t *testing.T) {
		// Create the mock wikipedia API
		page := "a"
		mockApi := createMockApi()
		// First request for a page should succeed
		_, err := mockApi.FetchOutboundPageLinks(page)
		if err != nil {
			t.Fatal("error on first access")
		}
		// Second request for the same page should fail
		_, err = mockApi.FetchOutboundPageLinks(page)
		if err == nil {
			t.Fatal("mock API didn't fail and it should have")
		}
	})

	t.Run("3 rounds of mapping links should leave only the 'g' page ungraphed.", func(t *testing.T) {
		// Arrange
		pages := map[string]bool{"a": true}
		mockApi := createMockApi()
		// Act
		graph := New()
		graph.MapEdges(0, pages, 3, 10, &mockApi)
		// Second request for the same page should fail
		_, found := mockApi.linkMap["g"]
		if !found {
			t.Fatal("Couldn't find g")
		}
		if len(mockApi.linkMap) > 1 {
			fmt.Println(mockApi.linkMap)
			t.Fatal("Other values in graph")
		}

	})
}
