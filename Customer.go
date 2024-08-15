package GoVideoStore

import "fmt"

type Customer struct {
	name    string
	rentals []*Rental
}

func NewCustomer(name string) *Customer {
	return &Customer{name: name}
}

func (customer *Customer) addRental(rental *Rental) {
	customer.rentals = append(customer.rentals, rental)
}

func (customer *Customer) statement() string {
	totalAmount := 0.0
	frequentRenterPoints := 0
	result := "Rental Record for " + customer.name + "\n"
	for _, rental := range customer.rentals {
		thisAmount := 0.0
		switch rental.movie.movieType {
		case NewRelease:
			thisAmount += float64(rental.daysRented * 3)
		case Regular:
			thisAmount += 2
			if rental.daysRented > 2 {
				thisAmount += float64(rental.daysRented-2) * 1.5
			}
		case Childrens:
			thisAmount += 1.5
			if rental.daysRented > 3 {
				thisAmount += float64(rental.daysRented-3) * 1.5
			}
		}
		frequentRenterPoints++
		if rental.movie.movieType == NewRelease && rental.daysRented > 1 {
			frequentRenterPoints++
		}
		result += fmt.Sprintf("\t%s\t%.1f\n", rental.movie.title, thisAmount)
		totalAmount += thisAmount
	}
	return result + fmt.Sprintf("Amount owed is %.1f\nYou earned %d frequent renter points", totalAmount, frequentRenterPoints)
}
