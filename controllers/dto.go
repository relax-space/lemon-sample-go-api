package controllers

const (
	DefaultMaxResultCount = 30
)

type SearchInput struct {
	Sortby         []string `query:"sortby"`
	Order          []string `query:"order"`
	SkipCount      int      `query:"skipCount"`
	MaxResultCount int      `query:"maxResultCount"`
}

type Book struct {
	Name    string   `json:"name" query:"name"`
	Authors []Author `json:"authors" query:"authors"`
	Sign    string   `json:"sign" query:"sign"`
}

type Author struct {
	Name string `json:"name" query:"name"`
	Age  int64  `json:"age" query:"age"`
}
