package graph

type LinkType int

const (
	INBOUND LinkType = iota
	OUTBOUND
)

type Graph struct {
	pages map[string]Links
}

type Links struct {
	OutboundLinks []string
	InboundLinks  []string
}

func NewGraph() *Graph {
	return &Graph{
		pages: make(map[string]Links),
	}
}

// Add an edge to the graph
func (g *Graph) AddEdges(from string, to ...string) {
	// Add all outbound links
	srcPage := g.pages[from]
	srcPage.OutboundLinks = append(srcPage.OutboundLinks, to...)
	g.pages[from] = srcPage
	// Add all inbound links
	for _, page := range to {
		dstPage := g.pages[page]
		dstPage.InboundLinks = append(dstPage.InboundLinks, from)
		g.pages[page] = dstPage
	}
}

func (g *Graph) GetLinks(page string, kind LinkType) []string {
	switch kind {
	case OUTBOUND:
		return g.pages[page].OutboundLinks
	case INBOUND:
		return g.pages[page].OutboundLinks
	}
	return []string{}
}

func (g *Graph) GetPages() []string {
	var pages []string
	for page := range g.pages {
		pages = append(pages, page)
	}
	return pages
}

func (g *Graph) GetAdjList() map[string]Links {
	return g.pages
}
