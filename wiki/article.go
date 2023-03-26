package wiki

import (
	"strings"
	"unicode"
)

const baseURL = "https://en.wikipedia.org/w/api.php"

type WikiArticle struct {
	title string
}

func (article WikiArticle) FormatTitle() string {
	r := []rune(strings.ToLower(article.title))
	r[0] = unicode.ToUpper(r[0])
	return string(r)
}
