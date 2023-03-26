package wiki

import (
	"net/url"
	"strconv"
)

type Extracts struct {
	intro     bool
	plaintext bool
}

func (ex Extracts) FormatURL(article WikiArticle) *url.URL {
	// Initialize URL
	base, _ := url.Parse(baseURL)

	// Add required query params
	params := url.Values{}
	params.Add("action", "query")
	params.Add("titles", article.FormatTitle()) // cast first letter to upper
	params.Add("prop", "extracts")

	// Add optional params
	params.Add("exintro", strconv.FormatBool(ex.intro))
	params.Add("explaintext", strconv.FormatBool(ex.intro))

	// Format and return URL
	base.RawQuery = params.Encode()
	return base
}
