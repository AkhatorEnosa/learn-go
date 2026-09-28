package helpers

import "strings"


func ValidateInputs(firstName string, lastName string, email string, userTickets uint) (bool, bool, bool) {
	isValidName := len(firstName) >= 2 && len(lastName) >= 2
	isValidEmail := strings.Contains(email, "@")
	isValidTicketCount := userTickets > 0 

	return isValidName, isValidEmail, isValidTicketCount
}