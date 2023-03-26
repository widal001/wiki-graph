package wiki

import (
	"strings"
	"testing"
)

func TestExtractsFormatURL(t *testing.T) {

	cases := []struct {
		Description string
		Query       PropQuery
		Include     []string
		Exclude     []string
	}{
		{
			Description: "intro in plaintext",
			Query:       Extracts{intro: true, plaintext: true},
			Include:     []string{"&exintro=true", "&explaintext=true"},
			Exclude:     []string{"excontinue"},
		},
		{
			Description: "intro in HTML",
			Query:       Extracts{intro: true, plaintext: false},
			Include:     []string{"&exintro=true", "&explaintext=true"},
			Exclude:     []string{"excontinue"},
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
