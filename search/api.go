package search

type ISearch interface {
	SearchText(string) (SearchResults, error)
}

type SearchResults struct {
	Count   int
	Results []Post
}

type Post struct {
	Type    string
	Title   string
	Summary string
	Details string
}

type search struct {
}

func NewSearch() (ISearch, error) {
	return &search{}, nil
}

func (s *search) SearchText(text string) (SearchResults, error) {
	return SearchResults{
		Count:   0,
		Results: make([]Post, 0),
	}, nil
}
