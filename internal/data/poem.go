package data

// Poem represents a single poem or ci from the huajianji dataset.
type Poem struct {
	Title      string   `json:"title"`
	Author     string   `json:"author"`
	Dynasty    string   `json:"dynasty"`
	Notes      []string `json:"notes"`
	Paragraphs []string `json:"paragraphs"`
}

// Collection holds a named set of poems.
type Collection struct {
	Name  string
	Poems []Poem
}
