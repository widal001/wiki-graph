package main

import (
	"fmt"
	"os"
	"wiki-graph/internals/scraper"
)

func main() {
	rootPage := "SPSS"
	maxDepth := 2
	wikiGraph, err := scraper.BuildLinkGraph(rootPage, maxDepth)
	if err != nil {
		fmt.Printf("Error building link graph: %v", err)
		os.Exit(1)
	}

	for page, links := range wikiGraph.GetAdjList() {
		fmt.Printf("Page: %s\n", page)
		fmt.Println("- Number of outbound links:", len(links.OutboundLinks))
		fmt.Println("- Number of inbound links:", len(links.InboundLinks))
	}
}
