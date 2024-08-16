package GoVideoStore

type Rental struct {
	movie      *Movie
	daysRented int
}

func (rental Rental) getTitle() string {
	return rental.movie.title
}

func NewRental(movie *Movie, daysRented int) *Rental {
	return &Rental{movie: movie, daysRented: daysRented}
}
