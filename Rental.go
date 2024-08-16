package GoVideoStore

type Rental struct {
	movie      *Movie
	daysRented int
	rentalType RentalType
}

func (rental Rental) getTitle() string {
	return rental.movie.title
}

func (rental Rental) determineAmount() float64 {
	return rental.rentalType.determineAmount(rental.daysRented)
}

func (rental Rental) determinePoints() int {
	return rental.rentalType.determinePoints(rental.daysRented)
}

func NewRental(movie *Movie, rentalType RentalType, daysRented int) *Rental {
	return &Rental{
		movie:      movie,
		rentalType: rentalType,
		daysRented: daysRented}
}
