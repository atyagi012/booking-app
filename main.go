package main

import (
	"booking-app/helper"
	"fmt"
	"strings"
)

var conferenceName = "Go Conference"

const conferenceTicket = 50

var remainingTicket uint = 50

// Slice
var bookings = []string{} //Another way to declare slice

func main() {
	greetUser()

	for remainingTicket > 0 && len(bookings) < 50 {
		//Get user input
		firstName, lastName, email, userTicket := getUserInput()

		//Validate user input
		isValidName, isValidEmail, isValidTicketNumber := helper.ValidateUserInput(firstName, lastName, email, userTicket, remainingTicket)

		if isValidName && isValidEmail && isValidTicketNumber {
			//Book ticket
			bookTicket(userTicket, firstName, lastName, email)

			//Get first name of user
			firstNames := getFirstName()
			fmt.Printf("The first name of bookings : %v\n", firstNames)

			if remainingTicket == 0 {
				fmt.Println("All tickets are booked. Thank you for booking")
				break
			}
		} else {
			if !isValidName {
				fmt.Println("Please enter valid name")
			}
			if !isValidEmail {
				fmt.Println("Please enter valid email")
			}
			if !isValidTicketNumber {
				fmt.Println("Please enter valid ticket number")
			}
		}

		/*
			//switch case
			city := "Delhi"

			switch city {
			case "Delhi":
				//Code for delhi city
			case "Mumbai":
				//Code for Mumbai city
			case "Bangalore", "Chennai":
				//Code for Bangalore and Chennai city
			default:
				//Code for other cities
			}
		*/
	}

	tickets := getRemainingTickets(remainingTicket)
	fmt.Println("Remaining tickets are : ", tickets)
}

// Functions
func greetUser() {
	fmt.Println("Welcome to Go Conference")
	fmt.Printf("Welcome to %v booking application \n", conferenceName)

	fmt.Println("We have total tickets ", conferenceTicket, "and remaining tickets are ", remainingTicket)
	fmt.Println("Get your booking done in few clicks")
}

func getFirstName() []string {
	firstNames := []string{}
	for _, booking := range bookings {
		var names = strings.Fields(booking)
		firstNames = append(firstNames, names[0])
	}
	return firstNames
}

// Return value from function
func getRemainingTickets(remainingTicket uint) uint {
	return remainingTicket
}

// func isValidEmail(email string) []string {
// 	//Get email code Here
// 	allUserEmails := []string{}

// 	//if email.
// 	return allUserEmails
// }

func getUserInput() (string, string, string, uint) {
	var firstName string
	var lastName string
	var email string
	var userTicket uint

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
	fmt.Println("Email is : ", email)
	//isValidEmail(email)

	//Number of tickets to book
	fmt.Println("Enter number of tickets you want to book: ")
	fmt.Scan(&userTicket)

	return firstName, lastName, email, userTicket
}

func bookTicket(userTicket uint, firstName string, lastName string, email string) {
	remainingTicket = remainingTicket - userTicket
	bookings = append(bookings, firstName+" "+lastName)

	fmt.Printf("Thank you '%v %v' for booking %v tickets. You will get confirmation email at %v \n", firstName, lastName, userTicket, email)
	fmt.Printf("%v tickets remaining for %v\n", remainingTicket, conferenceName)
}
