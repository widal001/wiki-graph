package wiki

import (
	"fmt"
	"testing"
)

func TestFormatURL(t *testing.T) {

	cases := []struct {
		Title       string
		Want        string
		Description string
	}{
		{"Economics", "Economics", "no change"},
		{"Computer Science", "Computer+science", "replace space and fix case"},
		{"C++", "C%2B%2B", "replace special characters"},
	}

	for _, test := range cases {
		t.Run(test.Description, func(t *testing.T) {
			article := WikiArticle{title: test.Title}
			want := fmt.Sprintf(
				"https://en.wikipedia.org/w/api.php?action=query&pllimit=500&prop=links&titles=%s",
				test.Want,
			)
			got := article.FormatURL()
			assertEquals(t, got.String(), want)
		})
	}
}

func assertEquals(t *testing.T, got string, want string) {
	t.Helper()

	if got != want {
		t.Errorf("\nGot:    %s \nWanted: %s", got, want)
	}
}
