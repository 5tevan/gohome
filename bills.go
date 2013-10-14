package main

import(
	"time"
)

//bill paid by a user to an external company.
//money has left the house and needs to be paid back
type Bill struct{
	payer User
	recipient string
	time time.Time
	amount float32
	balance float32
	// figure out a way to store the division of the bill
}

// single payment for a single bill from a user to another
type BillPayment struct{
	payer User
	recipient User
	time time.Time
	amount float32
	bill Bill
}

// a single payment from one user to another. Covers many bills
type Payment struct {
	amount float32
	billpayments []BillPayment
}