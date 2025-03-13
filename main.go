package main

import "fmt"

func main() {
	var conferenceName = "Go Conference"
	const conferenceTicket = 50
	var remainingTicket uint = 40
	//remainingTicket := 50 //other way to declare variable. No need to mention type of variable

	fmt.Printf("Welcome to %v booking application \n", conferenceName)
	//fmt.Println("Welcome to "+conferenceName, "booking application")
	fmt.Println("We have total tickets ", conferenceTicket, "and remaining tickets are ", remainingTicket)
	fmt.Println("Get your booking done in few clicks")

	fmt.Println("Total tickets are : ", conferenceTicket)

	var firstName string
	var lastName string
	var email string
	var userTicket uint

	//Array
	//var booking = [50]string{"John", "Doe"}
	//var booking = [50]string{} //Empty array
	//var bookings [50]string
	//bookings[0] = "John"
	//bookings[1] = "Doe"

	//Slice
	//var bookings []string
	//var bookings = []string{}    //Another way to declare slice
	bookings := []string{} //Another way to declare slice

	//Ask user to enter the name
	//userName = "John"
	//bookedTicket = 5

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

	//To get mamory location of variable
	//fmt.Println("Memory location of userName variable is: ", &userName)

	//To get the type of variable
	//fmt.Printf("Type of bookedTicket is : %T \n", bookedTicket)

	//fmt.Printf("User '%v' has booked %v tickets \n", firstName, userTicket)

	remainingTicket = remainingTicket - userTicket
	//bookings[0] = firstName + " " + lastName
	bookings = append(bookings, firstName+" "+lastName)

	//Printing the whole array
	//fmt.Printf("The whole array: %v\n", bookings)
	fmt.Printf("The whole slice: %v\n", bookings)

	//Printing the array with index
	//fmt.Printf("The first value in array: %v \n", bookings[0])
	fmt.Printf("The first value in slice: %v \n", bookings[0])

	//fmt.Printf("Type of array is : %T\n", bookings)
	fmt.Printf("Type of slice is : %T\n", bookings)

	fmt.Printf("Length of slice is : %v\n", len(bookings))

	fmt.Printf("Thank you '%v %v' for booking %v tickets. You will get confirmation email at %v \n", firstName, lastName, userTicket, email)

	fmt.Println("Remaining tickets are : ", remainingTicket)

	fmt.Printf("There are all our bookings : %v\n", bookings)
}
