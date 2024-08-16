package GoVideoStore

type NewReleaseRental struct{}

func (rentalType NewReleaseRental) determineAmount(daysRented int) float64 {
	return float64(daysRented * 3)
}

func (rentalType NewReleaseRental) determinePoints(daysRented int) int {
	if daysRented > 1 {
		return 2
	}
	return 1
}
