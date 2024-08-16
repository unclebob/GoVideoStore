package GoVideoStore

import (
	"testing"
)

var customer *Customer
var newRelease1 *Movie
var newRelease2 *Movie
var childrensMovie *Movie
var regular1 *Movie
var regular2 *Movie
var regular3 *Movie

func setup() {
	customer = NewCustomer("Customer Name")
	newRelease1 = NewMovie("New Release 1", NewRelease)
	newRelease2 = NewMovie("New Release 2", NewRelease)
	childrensMovie = NewMovie("Childrens", Childrens)
	regular1 = NewMovie("Regular 1", Regular)
	regular2 = NewMovie("Regular 2", Regular)
	regular3 = NewMovie("Regular 3", Regular)
}

func assertStringsEqual(t *testing.T, expected string, actual string) {
	if actual != expected {
		t.Errorf("\nExpected:\n%s\n\nGot:\n%s", expected, actual)
	}
}

func assertIntsEqual(t *testing.T, expected int, actual int) {
	if actual != expected {
		t.Errorf("\nExpected:%d Got:%d", expected, actual)
	}
}

func assertFloatsEqual(t *testing.T, expected float64, actual float64) {
	if actual != expected {
		t.Errorf("\nExpected:%f Got:%f", expected, actual)
	}
}

func assertOwedAndPoints(t *testing.T, owed float64, points int) {
	customer.makeStatement()
	assertFloatsEqual(t, owed, customer.totalAmount)
	assertIntsEqual(t, points, customer.frequentRenterPoints)
}

func TestTotalsForOneNewRelease(t *testing.T) {
	setup()
	customer.addRental(NewRental(newRelease1, 3))
	assertOwedAndPoints(t, 9.0, 2)
}

func TestTotalsForTwoNewReleases(t *testing.T) {
	setup()
	customer.addRental(NewRental(newRelease1, 3))
	customer.addRental(NewRental(newRelease2, 3))
	assertOwedAndPoints(t, 18.0, 4)
}

func TestTotalsForOneChildrensMovie(t *testing.T) {
	setup()
	customer.addRental(NewRental(childrensMovie, 3))
	assertOwedAndPoints(t, 1.5, 1)
}

func TestTotalsForManyRegularMovies(t *testing.T) {
	setup()
	customer.addRental(NewRental(regular1, 1))
	customer.addRental(NewRental(regular2, 2))
	customer.addRental(NewRental(regular3, 3))
	assertOwedAndPoints(t, 7.5, 3)
}

func TestStatementFormat(t *testing.T) {
	setup()
	customer.addRental(NewRental(regular1, 1))
	customer.addRental(NewRental(regular2, 2))
	customer.addRental(NewRental(regular3, 3))
	assertStringsEqual(t,
		"Rental Record for Customer Name\n"+
			"\tRegular 1\t2.0\n"+
			"\tRegular 2\t2.0\n"+
			"\tRegular 3\t3.5\n"+
			"Amount owed is 7.5\n"+
			"You earned 3 frequent renter points",
		customer.makeStatement())
}
