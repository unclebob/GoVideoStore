package GoVideoStore

import (
	"testing"
)

var statement *RentalStatement
var newRelease1 *Movie
var newRelease2 *Movie
var childrensMovie *Movie
var regular1 *Movie
var regular2 *Movie
var regular3 *Movie

func setup() {
	statement = NewStatement("RentalStatement Name")
	newRelease1 = NewMovie("New Release 1")
	newRelease2 = NewMovie("New Release 2")
	childrensMovie = NewMovie("Childrens")
	regular1 = NewMovie("Regular 1")
	regular2 = NewMovie("Regular 2")
	regular3 = NewMovie("Regular 3")
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
	statement.makeStatement()
	assertFloatsEqual(t, owed, statement.totalAmount)
	assertIntsEqual(t, points, statement.frequentRenterPoints)
}

func TestTotalsForOneNewRelease(t *testing.T) {
	setup()
	statement.addRental(NewRental(newRelease1, NewRelease, 3))
	assertOwedAndPoints(t, 9.0, 2)
}

func TestTotalsForTwoNewReleases(t *testing.T) {
	setup()
	statement.addRental(NewRental(newRelease1, NewRelease, 3))
	statement.addRental(NewRental(newRelease2, NewRelease, 3))
	assertOwedAndPoints(t, 18.0, 4)
}

func TestTotalsForOneChildrensMovie(t *testing.T) {
	setup()
	statement.addRental(NewRental(childrensMovie, Childrens, 3))
	assertOwedAndPoints(t, 1.5, 1)
}

func TestTotalsForManyRegularMovies(t *testing.T) {
	setup()
	statement.addRental(NewRental(regular1, Regular, 1))
	statement.addRental(NewRental(regular2, Regular, 2))
	statement.addRental(NewRental(regular3, Regular, 3))
	assertOwedAndPoints(t, 7.5, 3)
}

func TestStatementFormat(t *testing.T) {
	setup()
	statement.addRental(NewRental(regular1, Regular, 1))
	statement.addRental(NewRental(regular2, Regular, 2))
	statement.addRental(NewRental(regular3, Regular, 3))
	assertStringsEqual(t,
		"Rental Record for RentalStatement Name\n"+
			"\tRegular 1\t2.0\n"+
			"\tRegular 2\t2.0\n"+
			"\tRegular 3\t3.5\n"+
			"Amount owed is 7.5\n"+
			"You earned 3 frequent renter points",
		statement.makeStatement())
}
