package GoVideoStore

import "fmt"

type Customer struct {
	name                 string
	rentals              []*Rental
	totalAmount          float64
	frequentRenterPoints int
}

func NewCustomer(name string) *Customer {
	return &Customer{name: name}
}

func (customer *Customer) addRental(rental *Rental) {
	customer.rentals = append(customer.rentals, rental)
}

func (customer *Customer) makeStatement() string {
	customer.clearTotals()
	return customer.makeHeader() +
		customer.makeDetails() +
		customer.makeFooter()
}

func (customer *Customer) clearTotals() {
	customer.totalAmount = 0.0
	customer.frequentRenterPoints = 0
}

func (customer *Customer) makeHeader() string {
	return "Rental Record for " + customer.name + "\n"
}

func (customer *Customer) makeDetails() string {
	rentalDetails := ""
	for _, rental := range customer.rentals {
		rentalDetails += customer.makeDetail(rental)
	}
	return rentalDetails
}

func (customer *Customer) makeDetail(rental *Rental) string {
	amount := rental.determineAmount()
	customer.frequentRenterPoints += rental.determinePoints()
	customer.totalAmount += amount
	return customer.formatDetail(rental, amount)
}

func (customer *Customer) formatDetail(rental *Rental, amount float64) string {
	return fmt.Sprintf("\t%s\t%.1f\n", rental.getTitle(), amount)
}

func (customer *Customer) makeFooter() string {
	return fmt.Sprintf("Amount owed is %.1f\n"+
		"You earned %d frequent renter points",
		customer.totalAmount, customer.frequentRenterPoints)
}
