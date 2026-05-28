package main

import (
	"fmt"
)

func main() {
	//declare variables and constants
	var conferenceName = "Go Conference"
	// or conferenceName := "Go Conference"
	const conferenceTickets = 50
	var remainingTickets uint = 50

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

	// // Data types using the Printf function
	var firstName string
	var lastName string
	var email string
	var userTickets uint

	// //Ask user for their name
	// userName = "Tom"
	// userTickets = 2
	// fmt.Printf("User %v booked %v tickets.\n", userName, userTickets)

	// // getting user's input using scan (A scan function assigns the user's value to the userName variables) and pointers
	// fmt.Println("Enter your first name: ")
	// fmt.Scan(&firstName)

	// fmt.Println("Enter your last name: ")
	// fmt.Scan(&lastName)

	// fmt.Println("Enter your email address: ")
	// fmt.Scan(&email)

	// fmt.Println("Enter number of tickets: ")
	// fmt.Scan(&userTickets)

	// // user's output with a confirmation message
	// fmt.Printf("Thank you %v %v for booking %v tickets. you will receive a confirmation email at %v\n", firstName, lastName, userTickets, email)

	// getting user's input using scan (A scan function assigns the user's value to the userName variables) and pointers. Also showing the tickets left for the other users to be aware of
	fmt.Println("Enter your first name: ")
	fmt.Scan(&firstName)

	fmt.Println("Enter your last name: ")
	fmt.Scan(&lastName)

	fmt.Println("Enter your email address: ")
	fmt.Scan(&email)

	fmt.Println("Enter number of tickets: ")
	fmt.Scan(&userTickets)

	remainingTickets = remainingTickets - userTickets

	// user's output with a confirmation message
	fmt.Printf("Thank you %v %v for booking %v tickets. you will receive a confirmation email at %v\n", firstName, lastName, userTickets, email)
	fmt.Printf("%v tickets remaining for %v\n", remainingTickets, conferenceName)

	// getting user's input using pointers (A pointer is a variable that points to the memory address of another variable (&) or is a special variable)
	// fmt.Println(&remainingTickets)

	//  data type of variables conferenceName, remaininTickets and conferenceTickets are shown using the Printf function
	// fmt.Printf("conferenceName is %T,conferenceTickets is %T, remainingTickets is %T\n", conferenceName, remainingTickets, conferenceTickets)

}
