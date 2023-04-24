package main

import "fmt"

func main() {
	var conferenceName = "Games Fest"
	const conferenceTickets = 50
	var remainingTickets = 50

	fmt.Println("Welcome to Conference bookings!")
	fmt.Printf("In order to attend %v please purchase a ticket.\n", conferenceName)
	fmt.Printf("We have %v tickets in total and there are %v left.\n", conferenceTickets, remainingTickets)

}
