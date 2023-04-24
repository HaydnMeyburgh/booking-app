package main

import "fmt"

func main() {
	var conferenceName = "Games Fest"
	const conferenceTickets = 50
	var remainingTickets = 50
	var bookings []string

	fmt.Println("Welcome to Conference bookings!")
	fmt.Printf("In order to attend %v please purchase a ticket.\n", conferenceName)
	fmt.Printf("We have %v tickets in total and there are %v left.\n", conferenceTickets, remainingTickets)

	var firstName string
	var lastName string
	var email string
	var userTickets int

	fmt.Println("Please enter your first name:")
	fmt.Scan(&firstName)

	fmt.Println("Please enter your last name:")
	fmt.Scan(&lastName)

	fmt.Println("Please enter your email address:")
	fmt.Scan(&email)

	fmt.Println("How many tickets would you like to buy?:")
	fmt.Scan(&userTickets)

	remainingTickets -= userTickets
	bookings = append(bookings, firstName+" "+lastName)

	fmt.Printf("Thank you %v %v for booking %v tickets. Expect your confirmation email to be sent to %v shortly.\n", firstName, lastName, userTickets, email)
	fmt.Printf("%v tickets remaining for %v\n", remainingTickets, conferenceName)

	fmt.Printf("These are all our bookings so far: %v", bookings)
}
