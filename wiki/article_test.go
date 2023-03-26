package wiki

import (
	"testing"
)

func TestFormatTitle(t *testing.T) {

	cases := []struct {
		Description string
		Title       string
		Want        string
	}{
		{
			Title:       "Economics",
			Want:        "Economics",
			Description: "no change",
		},
		{
			Title:       "Computer Science",
			Want:        "Computer science",
			Description: "replace space and fix case",
		},
	}

	for _, test := range cases {
		t.Run(test.Description, func(t *testing.T) {
			article := WikiArticle{title: test.Title}
			got := article.FormatTitle()
			assertEquals(t, got, test.Want)
		})
	}
}

func assertEquals(t *testing.T, got string, want string) {
	t.Helper()

	if got != want {
		t.Errorf("\nGot:    %s \nWanted: %s", got, want)
	}
}
