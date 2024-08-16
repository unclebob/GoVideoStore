package GoVideoStore

type Movie struct {
	title string
}

func NewMovie(title string) *Movie {
	return &Movie{title: title}
}
