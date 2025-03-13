package main

import "fmt"

func main() {
	var conferenceName = "Go Conference"
	const conferenceTicket = 50
	var remainingTicket uint = 40

	fmt.Printf("Welcome to %v booking application \n", conferenceName)
	fmt.Println("We have total tickets ", conferenceTicket, "and remaining tickets are ", remainingTicket)
	fmt.Println("Get your booking done in few clicks")

	var firstName string
	var lastName string
	var email string
	var userTicket uint

	//Slice
	bookings := []string{} //Another way to declare slice

	//Ask user to enter the name
	//Get input from user: First Name
	fmt.Println("Enter first name: ")
	fmt.Scan(&firstName)

	//Last Name
	fmt.Println("Enter last name: ")
	fmt.Scan(&lastName)

	//Email address
	fmt.Println("Enter email address: ")
	fmt.Scan(&email)

	//Number of tickets to book
	fmt.Println("Enter number of tickets you want to book: ")
	fmt.Scan(&userTicket)

	remainingTicket = remainingTicket - userTicket
	bookings = append(bookings, firstName+" "+lastName)

	fmt.Printf("Thank you '%v %v' for booking %v tickets. You will get confirmation email at %v \n", firstName, lastName, userTicket, email)

	fmt.Println("Remaining tickets are : ", remainingTicket)

	fmt.Printf("There are all our bookings : %v\n", bookings)
}
