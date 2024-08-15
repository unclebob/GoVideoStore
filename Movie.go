package GoVideoStore

const (
	Regular int = iota
	NewRelease
	Childrens
)

type Movie struct {
	title     string
	movieType int
}

func NewMovie(title string, movieType int) *Movie {
	return &Movie{title: title, movieType: movieType}
}
