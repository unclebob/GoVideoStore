package GoVideoStore

type Rental struct {
	movie      *Movie
	daysRented int
}

func NewRental(movie *Movie, daysRented int) *Rental {
	return &Rental{movie: movie, daysRented: daysRented}
}
