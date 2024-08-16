package GoVideoStore

type Rental struct {
	movie      *Movie
	daysRented int
}

func (rental Rental) getTitle() string {
	return rental.movie.title
}

func (rental Rental) determineAmount() float64 {
	amount := 0.0
	switch rental.movie.movieType {
	case NewRelease:
		amount += float64(rental.daysRented * 3)
	case Regular:
		amount += 2
		if rental.daysRented > 2 {
			amount += float64(rental.daysRented-2) * 1.5
		}
	case Childrens:
		amount += 1.5
		if rental.daysRented > 3 {
			amount += float64(rental.daysRented-3) * 1.5
		}
	}
	return amount
}

func (rental Rental) determinePoints() int {
	if rental.movie.movieType == NewRelease && rental.daysRented > 1 {
		return 2
	}
	return 1
}

func NewRental(movie *Movie, daysRented int) *Rental {
	return &Rental{movie: movie, daysRented: daysRented}
}
