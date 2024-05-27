package client

import (
	"encoding/json"
	"wiki-graph/pkg/models"
)

const (
	BASE_URL   = "https://en.wikipedia.org/w/api.php"
	ACTION     = "query"
	PROPERTY   = "links"
	FORMAT     = "json"
	PAGE_TYPE  = "0"
	BATCH_SIZE = "5"
)

func (c *Client) GetOutboundLinks(title string) ([]string, error) {

	url := "https://en.wikipedia.org/w/api.php"
	params := map[string]string{
		"action":      ACTION,
		"prop":        PROPERTY,
		"format":      FORMAT,
		"ns":          PAGE_TYPE,
		"pllimit":     BATCH_SIZE,
		"plnamespace": PAGE_TYPE,
		"titles":      title,
	}
	body, err := c.Get(url, params)
	if err != nil {
		return nil, err
	}

	var apiResp models.PageLinksResponse
	if err := json.Unmarshal(body, &apiResp); err != nil {
		return nil, err
	}

	return apiResp.ExtractLinks(), nil
}
