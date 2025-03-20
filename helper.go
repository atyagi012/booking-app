package main

import "strings"

// Validate user input
func validateUserInput(firstName string, lastName string, email string, userTicket uint) (bool, bool, bool) {
	isValidName := len(firstName) >= 2 && len(lastName) >= 2
	isValidEmail := strings.Contains(email, "@") && strings.Contains(email, ".")
	isValidTicketNumber := userTicket > 0 && userTicket <= remainingTicket

	//Can return multiple values in GO
	return isValidName, isValidEmail, isValidTicketNumber

	// if isValidName && isValidEmail && isValidTicketNumber {
	// 	return true
	// }
	// return false
}
