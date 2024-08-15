package GoVideoStore

import (
	"testing"
)

var customer *Customer

func setup() {
	customer = NewCustomer("Fred")
}

func assertEquals(t *testing.T, expected string, statement string) {
	if statement != expected {
		t.Errorf("\nExpected:\n%s\n\nGot:\n%s", expected, statement)
	}
}

func TestSingleNewReleaseStatement(t *testing.T) {
	setup()
	customer.addRental(NewRental(NewMovie("The Cell", NewRelease), 3))
	assertEquals(t,
		"Rental Record for Fred\n"+
			"\tThe Cell\t9.0\n"+
			"Amount owed is 9.0\n"+
			"You earned 2 frequent renter points",
		customer.statement())
}

func TestDualNewReleaseStatement(t *testing.T) {
	setup()
	customer.addRental(NewRental(NewMovie("The Cell", NewRelease), 3))
	customer.addRental(NewRental(NewMovie("The Tigger Movie", NewRelease), 3))
	assertEquals(t,
		"Rental Record for Fred\n"+
			"\tThe Cell\t9.0\n"+
			"\tThe Tigger Movie\t9.0\n"+
			"Amount owed is 18.0\n"+
			"You earned 4 frequent renter points",
		customer.statement())
}

func TestSingleChildrensStatement(t *testing.T) {
	setup()
	customer.addRental(NewRental(NewMovie("The Tigger Movie", Childrens), 3))
	assertEquals(t,
		"Rental Record for Fred\n"+
			"\tThe Tigger Movie\t1.5\n"+
			"Amount owed is 1.5\n"+
			"You earned 1 frequent renter points",
		customer.statement())
}

func TestMultipleRegularStatement(t *testing.T) {
	setup()
	customer.addRental(NewRental(NewMovie("Plan 9 from Outer Space", Regular), 1))
	customer.addRental(NewRental(NewMovie("8 1/2", Regular), 2))
	customer.addRental(NewRental(NewMovie("Eraserhead", Regular), 3))
	assertEquals(t,
		"Rental Record for Fred\n"+
			"\tPlan 9 from Outer Space\t2.0\n"+
			"\t8 1/2\t2.0\n"+
			"\tEraserhead\t3.5\n"+
			"Amount owed is 7.5\n"+
			"You earned 3 frequent renter points",
		customer.statement())
}
