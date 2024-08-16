package GoVideoStore

type ChildrensRental struct{}

func (rentalType ChildrensRental) determineAmount(daysRented int) float64 {
	amount := 1.5
	if daysRented > 3 {
		amount += float64(daysRented-3) * 1.5
	}
	return amount
}

func (rentalType ChildrensRental) determinePoints(daysRented int) int {
	return 1
}
