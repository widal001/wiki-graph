package wiki

import (
	"strings"
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

func TestPropLinkFormatURL(t *testing.T) {

	cases := []struct {
		Description string
		Query       PropQuery
		Include     []string
		Exclude     []string
	}{
		{
			Description: "link limit under 500",
			Query:       PropLinks{pllimit: 30},
			Include:     []string{"&pllimit=30"},
			Exclude:     []string{"plnamespace", "plcontinue"},
		},
		{
			Description: "link limit over 500",
			Query:       PropLinks{pllimit: 750},
			Include:     []string{"&pllimit=500"},
			Exclude:     []string{"plnamespace", "plcontinue"},
		},
		{
			Description: "limit under 10 and one namepspace",
			Query:       PropLinks{pllimit: 3, plnamespace: []int{0}},
			Include:     []string{"&plnamespace=0"},
			Exclude:     []string{"pllimit", "plcontinue"},
		},
		{
			Description: "link limit and multiple namepspaces",
			Query:       PropLinks{pllimit: 10, plnamespace: []int{0, 1}},
			Include:     []string{"&pllimit=10", "&plnamespace=0%7C1"},
			Exclude:     []string{"plcontinue"},
		},
		{
			Description: "link limit and continue",
			Query:       PropLinks{pllimit: 10, plcontinue: "continue"},
			Include:     []string{"&pllimit=10", "&plcontinue=continue"},
			Exclude:     []string{"plnamespace"},
		},
	}

	for _, test := range cases {
		t.Run(test.Description, func(t *testing.T) {
			// Setup
			article := WikiArticle{title: "Economics"}
			// Execution
			url := test.Query.FormatURL(article).String()
			// Validation -- check for required params
			for _, param := range test.Include {
				if !strings.Contains(url, param) {
					t.Errorf("%s is missing param: %s", url, param)
				}
			}
			// Validation -- check for absence of excluded params
			for _, param := range test.Exclude {
				if strings.Contains(url, param) {
					t.Errorf("%s contains param: %s", url, param)
				}
			}
		})
	}

}

func assertEquals(t *testing.T, got string, want string) {
	t.Helper()

	if got != want {
		t.Errorf("\nGot:    %s \nWanted: %s", got, want)
	}
}
