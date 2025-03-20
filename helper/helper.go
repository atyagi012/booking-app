package helper

import "strings"

//NOTE: To export function to other packages, function name should start with capital letter
// Validate user input
func ValidateUserInput(firstName string, lastName string, email string, userTicket uint, remainingTicket uint) (bool, bool, bool) {
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
