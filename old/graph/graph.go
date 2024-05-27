package graph

import (
	"fmt"
	"wiki-graph/old/api"
)

type PageLinkGraph struct {
	Pages map[string]bool
	Links []Link
}

type Link struct {
	From string
	To   string
}

func New() PageLinkGraph {
	return PageLinkGraph{Pages: make(map[string]bool)}
}

func (g *PageLinkGraph) MapEdges(
	round int,
	pagesToGraph map[string]bool,
	maxRound int,
	maxRequests int,
	api api.Api,
) *PageLinkGraph {
	fmt.Println("Starting round", round)
	if round >= maxRound {
		return g
	}
	newPagesToGraph := make(map[string]bool)
	for page := range pagesToGraph {
		// If a page has already been graphed moved to the next page
		if _, alreadyGraphed := g.Pages[page]; alreadyGraphed {
			continue
		}
		// Fetch the links for a page
		result, err := api.FetchOutboundPageLinks(page)
		if err != nil {
			fmt.Printf("error fetching links for %s\n", page)
		}
		// Add the page to the map of graphed pages
		g.Pages[page] = true
		// Add the links to the list of edges
		var links []Link
		for _, targetPage := range result.Links {
			// create a link from source page to target page
			links = append(links, Link{From: page, To: targetPage.Title})
			// Add the target page to the list of pages to graph
			newPagesToGraph[targetPage.Title] = true
		}
		g.Links = append(g.Links, links...)
	}
	if len(newPagesToGraph) == 0 {
		return g
	}
	return g.MapEdges(
		round+1,
		newPagesToGraph,
		maxRound,
		maxRequests,
		api,
	)
}
