package GoVideoStore

type RegularRental struct{}

func (rentalType RegularRental) determineAmount(daysRented int) float64 {
	amount := 2.0
	if daysRented > 2 {
		amount += float64(daysRented-2) * 1.5
	}
	return amount
}

func (rentalType RegularRental) determinePoints(daysRented int) int {
	return 1
}
