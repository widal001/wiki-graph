package wiki

import (
	"net/url"
	"strconv"
)

type PropLinks struct {
	plnamespace []int
	pllimit     int
	plcontinue  string
}

func (p PropLinks) FormatURL(article WikiArticle) *url.URL {

	// Initialize URL
	base, _ := url.Parse(baseURL)

	// Add required query params
	params := url.Values{}
	params.Add("action", "query")
	params.Add("titles", article.FormatTitle()) // cast first letter to upper
	params.Add("prop", "links")

	// Optionally add limit
	switch limit := p.pllimit; true {
	case limit > 500:
		params.Add("pllimit", "500")
	case limit >= 10:
		params.Add("pllimit", strconv.Itoa(limit))
	}

	// Optionally add namespace filter
	namespaces := p.plnamespace
	if len(namespaces) > 0 {
		namespaceString := p.formatNamespaces(namespaces)
		params.Add("plnamespace", namespaceString)
	}

	// Optionally add plcontinue filter
	if p.plcontinue != "" {
		params.Add("plcontinue", p.plcontinue)
	}

	// Format and return URL
	base.RawQuery = params.Encode()
	return base
}

func (p PropLinks) formatNamespaces(namespaces []int) string {
	namespaceString := ""
	for i := 0; i < len(namespaces); i++ {

		if i != 0 {
			namespaceString += "|"
		}
		namespaceString += strconv.Itoa(namespaces[i])
	}
	return namespaceString
}
