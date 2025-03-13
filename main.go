package main

import (
	"fmt"
	"strings"
)

func main() {
	var conferenceName = "Go Conference"
	const conferenceTicket = 50
	var remainingTicket uint = 40

	fmt.Printf("Welcome to %v booking application \n", conferenceName)
	fmt.Println("We have total tickets ", conferenceTicket, "and remaining tickets are ", remainingTicket)
	fmt.Println("Get your booking done in few clicks")

	for {
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

		if userTicket <= remainingTicket {
			remainingTicket = remainingTicket - userTicket
			bookings = append(bookings, firstName+" "+lastName)

			fmt.Printf("Thank you '%v %v' for booking %v tickets. You will get confirmation email at %v \n", firstName, lastName, userTicket, email)
			fmt.Println("Remaining tickets are : ", remainingTicket)

			firstNames := []string{}

			/*
				for index := range bookings {
					fmt.Println(bookings[index])
					var names = strings.Split(bookings[index], " ")[0]
					firstNames = append(firstNames, names)
				}
			*/
			for _, booking := range bookings {
				var names = strings.Fields(booking)
				firstNames = append(firstNames, names[0])
			}

			fmt.Printf("There are all first name of bookings : %v\n", firstNames)

			if remainingTicket == 0 {
				fmt.Println("All tickets are booked. Thank you for booking")
				break
			}
		} else {
			fmt.Printf("We have only %v tickets remaining so you can not book %v tickets", remainingTicket, userTicket)
			//break
			//continue
		}

	}
}
