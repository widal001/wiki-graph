package query

type ResponseSchema interface {
	Parse([]byte) (*ResponseSchema, error)
}

type PageLinksResponse struct {
	BatchComplete *string     `json:"batchcomplete"`
	QueryResult   queryResult `json:"query"`
	NextResult    nextResult  `json:"continue"`
}

type Page struct {
	PageId    uint   `json:"pageid"`
	Namespace uint8  `json:"ns"`
	Title     string `json:"title"`
	Summary   string `json:"extract"`
}

type queryResult struct {
	Pages []Page `json:"pages"`
}

type nextResult struct {
	Cursor string `json:"plcontinue"`
}

func (p PageLinksResponse) Parse(data []byte) (*PageLinksResponse, error) {
	return &PageLinksResponse{}, nil
}

func (p PageLinksResponse) ExtractLinks() []Page {
	return p.QueryResult.Pages
}
