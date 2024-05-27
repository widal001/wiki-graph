package models

type PageLinksResponse struct {
	BatchComplete *string     `json:"batchcomplete"`
	QueryResult   queryResult `json:"query"`
	NextResult    nextResult  `json:"continue"`
}

type queryResult struct {
	Pages map[string]Page `json:"pages"`
}

type nextResult struct {
	Cursor string `json:"plcontinue"`
}

func (p *PageLinksResponse) ExtractLinks() []string {
	links := []string{}
	for _, page := range p.QueryResult.Pages {
		for _, link := range page.Links {
			links = append(links, link.Title)
		}
	}
	return links
}
