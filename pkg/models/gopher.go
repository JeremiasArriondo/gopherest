package gopher

type Gopher struct {
	ID    string `json:"ID"`
	Name  string `json:"name,omitempty"`
	Image string `json:"image,omitempty"`
	Age   int    `json:"age,omitempty"`
}

type GopherRepository interface {
	CreateGopher(g *Gopher) error
	FetchGophers() ([]*Gopher, error)
	FetchGopherByID(ID string) (*Gopher, error)
}
