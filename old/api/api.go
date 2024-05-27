package api

import (
	"wiki-graph/old/client"
	"wiki-graph/old/query"
)

const BASE_URL = "https://en.wikipedia.org/w/api.php"

type Api interface {
	FetchOutboundPageLinks(page string) (PageLinks, error)
}

type WikipediaApi struct {
	baseURL string
	client  *client.Client
}

type PageLinks struct {
	Page  query.Page
	Links []query.Page
}

func New(client *client.Client) *WikipediaApi {
	return &WikipediaApi{
		baseURL: BASE_URL,
		client:  client,
	}
}

func (api *WikipediaApi) FetchOutboundPageLinks(pageTitle string) PageLinks {
	q := query.NewPageLinkQuery(pageTitle)
	links := q.FetchAll(*api.client)
	return PageLinks{
		Page:  query.Page{Title: pageTitle},
		Links: links,
	}

}
