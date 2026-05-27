package main

import "fmt"

func main() {
	//declare variables and constants
	var conferenceName = "Go Conference"
	const conferenceTickets = 50
	var remainingTickets = 50

	//Welcome message, value of conferenceName and conferenceTickets are used in the message using the Println function
	// fmt.Println("Welcome to", conferenceName, "booking application")
	// fmt.Println("We have a total of", conferenceTickets, "tickets and", remainingTickets, "are still available.")
	// fmt.Println("Get your tickets to attend")

	// // Data types using the Println function
	// var userName string
	// var userTickets int

	// //Ask user for their name
	// userName = "Tom"
	// userTickets = 2
	// fmt.Println(userName)
	// fmt.Println(userTickets)

	//Welcome message, value of conferenceName and conferenceTickets are used in the message using the Printf function
	fmt.Printf("Welcome to %v booking application\n", conferenceName)
	fmt.Printf("We have a total of %v tickets and %v are still available.\n", conferenceTickets, remainingTickets)
	fmt.Printf("Get your tickets to attend\n")

	// Data types using the Printf function
	var userName string
	var userTickets int

	//Ask user for their name
	userName = "Tom"
	userTickets = 2
	fmt.Printf("User %v booked %v tickets.\n", userName, userTickets)

}
