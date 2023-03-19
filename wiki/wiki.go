package wiki

import (
	"net/url"
	"strings"
	"unicode"
)

const baseURL = "https://en.wikipedia.org/w/api.php"

type WikiArticle struct {
	title string
}

func (article WikiArticle) FormatURL() *url.URL {

	// Initialize URL
	base, _ := url.Parse(baseURL)

	// Add query params
	params := url.Values{}
	params.Add("action", "query")
	params.Add("prop", "links")
	params.Add("pllimit", "500")
	params.Add("titles", article.FormatTitle()) // cast first letter to upper

	// Format and return URL
	base.RawQuery = params.Encode()
	return base
}

func (article WikiArticle) FormatTitle() string {
	r := []rune(strings.ToLower(article.title))
	r[0] = unicode.ToUpper(r[0])
	return string(r)
}
