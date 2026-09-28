package main

import (
	"fmt"
	"learn-go/helpers"
	"time"
	// "strconv"
)

var conferenceName string = "Go Conference"
const conferenceTickets uint = 50
var remainingTickets uint = conferenceTickets
var bookings = make([]UserData, 0)

// create structure aka struct for user data
type UserData struct{
	firstName string
	lastName string
	email string
	numberOfTickets uint
}

func main()  {

	// greet users
	greetUser(conferenceName, conferenceTickets)
	
	// fmt.Println("Hello and welcome to", conferenceName)
	// fmt.Printf("The amount of tickets provided are %v and the amount of tickets available right now are %v. Be fast and grab a ticket for yourself now!!!!\n", conferenceTickets, remainingTickets)


	for {
		if remainingTickets > 0 {
			firstName, lastName, email, userTickets := getUserInputs()

			// input validation
			isValidName, isValidEmail, isValidTicketCount := helpers.ValidateInputs(firstName, lastName, email, userTickets)


			if isValidName && isValidEmail && isValidTicketCount {
				if userTickets > remainingTickets {
					fmt.Printf("Only %v tickets remaining\n", remainingTickets)
					continue
				}

				// book the ticket
				bookTicket(firstName, lastName, email, userTickets)
				go sendTicket(userTickets, firstName, lastName, email)

				// get first names
 				firstNames := getFirstnames()

				// print first names
				fmt.Printf("These are all the firstnames from our bookings: %v\n", firstNames)

				// check for no tickets remaining
				if remainingTickets == 0 {
					fmt.Println("No tickets remaining")
				}
			} else {
				if !isValidName {
					fmt.Printf("First or last name is too short.\n")
				}

				if !isValidEmail {
					fmt.Printf("Email address is invalid.\n")
				}

				if !isValidTicketCount {
					fmt.Printf("Number of tickets is invalid.\n")
				}
				continue
			}

		} else {
			fmt.Printf("All tickets for %v has been sold. Let's do this again next year. Cheers\n", conferenceName)
			break
		}
	}
	

}

func greetUser(confName string, confTickets uint) {
	fmt.Printf("Welcome to %v booking application\n", confName)
	fmt.Printf("%v tickets are still available out of %v.\n", remainingTickets, confTickets)
	fmt.Printf("Get your tickets to attend.\n\n")
}

func getUserInputs() (string, string, string, uint) {
	var firstName string
	var lastName string
	var email string
	var userTickets uint
	
	fmt.Println("Enter Firstname: ")
	// get user input for firstName
	fmt.Scan(&firstName)
	

	fmt.Println("Enter Lastname")
	// get user input for lastname
	fmt.Scan(&lastName)
	

	fmt.Println("Enter Email Address: ")
	// get user input for email
	fmt.Scan(&email)
	

	fmt.Println("Enter number of tickers: ")
	// get user input for tickers
	fmt.Scan(&userTickets)

	return firstName, lastName, email, userTickets
}

func bookTicket(firstName string, lastName string, email string, userTickets uint) {
	remainingTickets = remainingTickets - userTickets

	// create a map for user data
	// var userData = make(map[string]any)

	// userData["firstName"] = firstName
	// userData["lastName"] = lastName
	// userData["email"] = email
	// userData["numberOfTickets"] = userTickets
	// userData["numberOfTickets"] = strconv.FormatUint(uint64(userTickets), 10)
	// map ends here

	// Make userData struct

	var userData = UserData{
		firstName: firstName,
		lastName: lastName,
		email: email,
		numberOfTickets: userTickets,
	}


	bookings = append(bookings, userData)

	fmt.Printf("Users with bookings is as follows: %v ", bookings)

	fmt.Printf("Thank you %v %v for booking %v tickets for %v.\n Remaining tickets is now %v.\n A confirmation email will be sent to your email %v \n\n", firstName, lastName, userTickets, conferenceName, remainingTickets, email)
}


func getFirstnames() []string {
	var firstNames []string
	
	for _, booking := range bookings {
		// var names = strings.Fields(booking) // split each names from bookings
		firstName := booking.firstName
		firstNames = append(firstNames, firstName)
	}

	return firstNames
}

func sendTicket(userTickets uint, firstName string, lastName string, email string) {
	time.Sleep(25 * time.Second)

	var ticket = fmt.Sprintf("%v tickets for %v %v", userTickets, firstName, lastName)

	fmt.Println("############")
	fmt.Printf("Sending Tickets: \n %v \n to email address %v \n", ticket, email)
	fmt.Println("############")
}