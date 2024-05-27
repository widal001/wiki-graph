package scraper

import (
	"fmt"
	"os"
	"sync"

	"wiki-graph/internals/client"
	"wiki-graph/internals/graph"
)

func BuildLinkGraph(rootPage string, maxDepth int) (*graph.Graph, error) {
	wikiClient := client.NewClient()
	linkGraph := graph.NewGraph()

	var mu sync.Mutex
	var wg sync.WaitGroup

	visited := make(map[string]bool)
	currentQueue := map[string]bool{rootPage: true}
	currentDepth := 0

	for currentDepth <= maxDepth && len(currentQueue) > 0 {
		nextQueue := make(map[string]bool)
		// fetch the links for each page in the current queue
		for page := range currentQueue {
			wg.Add(1)
			go func(page string) {
				defer wg.Done()

				// Check that the page hasn't already been visited
				mu.Lock()
				if visited[page] {
					mu.Unlock()
					return
				}
				visited[page] = true
				mu.Unlock()
				// Fetch the outbound links for this page
				links, err := wikiClient.GetOutboundLinks(page)
				if err != nil {
					fmt.Println(err)
					os.Exit(1)
				}

				// Add those links to the graph
				mu.Lock()
				linkGraph.AddEdges(page, links...)
				mu.Unlock()

				// Add those links to the queue for the next iteration
				mu.Lock()
				for _, link := range links {
					if !visited[link] {
						nextQueue[link] = true
					}
				}
				mu.Unlock()
			}(page)
		}
		wg.Wait()

		// Update the current depth queue
		currentQueue = nextQueue
		currentDepth++
	}

	return linkGraph, nil
}
