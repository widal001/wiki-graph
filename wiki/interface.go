package wiki

import "net/url"

type PropQuery interface {
	FormatURL(article WikiArticle) *url.URL
}
