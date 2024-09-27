package GoVideoStore

type RentalType interface {
	determineAmount(daysRented int) float64
	determinePoints(daysRented int) int
}

type RentalTypeFactory interface {
	make(typeName string) RentalType
}
