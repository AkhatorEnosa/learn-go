package main

import (
	"fmt"
	"strings"
)

func main()  {
	var conferenceName string = "Go Conference"
	const conferenceTickets uint = 50
	var remainingTickets uint = conferenceTickets

	// greet users
	greetUser(conferenceName, int(conferenceTickets), int(remainingTickets))
	
	// fmt.Println("Hello and welcome to", conferenceName)
	// fmt.Printf("The amount of tickets provided are %v and the amount of tickets available right now are %v. Be fast and grab a ticket for yourself now!!!!\n", conferenceTickets, remainingTickets)

	var bookings []string

	var firstName string
	var lastName string
	var email string
	var userTickets uint

	for {
		if remainingTickets > 0 {
			fmt.Println("Enter Firstname: ")
			// get user input for firstName
			fmt.Scan(&firstName);
			

			fmt.Println("Enter Lastname")
			// get user input for lastname
			fmt.Scan(&lastName);
			

			fmt.Println("Enter Email Address: ")
			// get user input for email
			fmt.Scan(&email);
			

			fmt.Println("Enter number of tickers: ")
			// get user input for tickers
			fmt.Scan(&userTickets);

			// input validation
			isValidName := len(firstName) >= 2 && len(lastName) >= 2
			isValidEmail := strings.Contains(email, "@")
			isValidTicketCount := userTickets > 0 



			if isValidName && isValidEmail && isValidTicketCount {
				if userTickets > remainingTickets {
					fmt.Printf("Only %v tickets remaining\n", remainingTickets)
					continue
				}
				remainingTickets = remainingTickets - userTickets;
				bookings = append(bookings, lastName + " " + firstName)

				fmt.Printf("Thank you %v %v for booking %v tickets for %v.\n Remaining tickets is now %v.\n A confirmation email will be sent to your email %v \n\n", firstName, lastName, userTickets, conferenceName, remainingTickets, email)


				firstNames := []string{}

				for _, booking := range bookings {
					var names = strings.Fields(booking) // split each names from bookings
					firstName := names[0]
					firstNames = append(firstNames, firstName)
				}


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

func greetUser(confName string, confTickets int, remainingTickets int) {
	fmt.Printf("Welcome to %v booking application\n", confName)
	fmt.Printf("%v tickets are still available out of %v.\n", remainingTickets, confTickets)
	fmt.Printf("Get your tickets to attend.\n\n")
}