package main

import (
	"fmt"
)

const conferenceTickets int = 50

var conferenceName = "Go conference"
var RemainingTickets int = 50
var bookings =  make([]UserData, 0)

type UserData struct {
	firstName string
	lastName string
	email string
	numberOfTickets int

}

func main() {

	greetUsers()

	for {

		firstName, lastName, email, userTickets := getUserInput()
		isValidName, isValidEmail, isValidTicketnumber := ValidateUserInput(firstName, lastName, email, 
			 userTickets, RemainingTickets)

		///////////////////////////////////////////////////////////////////
		// checking for booking
		if isValidName && isValidEmail && isValidTicketnumber {
			bookTicket(userTickets, firstName, lastName, email)
			sendTicket(userTickets, firstName, lastName, email)

			printFirstNames()

			if RemainingTickets == 0 {
				// end program
				fmt.Println("Tickets are sold out. Come back next year")
				break
			}
		} else {
			if !isValidName {
				fmt.Println("The firstname or lastname you entered is too short")
			}
			if !isValidEmail {
				fmt.Println("email address doe not contain @ sign")
			}
			if !isValidTicketnumber {
				fmt.Println("number of ticket is invalid")
			}

		}

	}

}

func greetUsers() {
	fmt.Printf("Welcome to %v booking application\n", conferenceName)
	fmt.Printf("we have a total number of %v tickets and %v tickets are still available for the conference\n",
		conferenceTickets, RemainingTickets)
	fmt.Println("Get your tickets here to attend")
}

func printFirstNames() {
	firstNames := []string{}
	for _, booking := range bookings {
		firstNames = append(firstNames, booking.firstName)
	}
	fmt.Printf("The first names of bookings are: %v\n", firstNames)
}

func getUserInput() (string, string, string, int) {
	var firstName string
	var lastName string
	var email string
	var userTickets int

	// POINTERS
	// A pointer is a variable that points to the memory of another variable
	// *Operator also termed as the dereferencing operator used to declare
	// pointer variable and access the value stored in the address.
	// &operator termed as address operator used to return the
	// address of a variable or to access the address of a variable to a pointer.
	fmt.Println("Enter your name")
	fmt.Scan(&firstName)

	fmt.Println("Enter your last name")
	fmt.Scan(&lastName)

	fmt.Println("Enter your email address")
	fmt.Scan(&email)

	fmt.Println("Enter your of tickets")
	fmt.Scan(&userTickets)
	return firstName, lastName, email, userTickets
}

func bookTicket(userTickets int, firstName string, lastName string, email string) {
	RemainingTickets = RemainingTickets - userTickets

	// create a map for user
	var userData =UserData {
		firstName: firstName,
		lastName: lastName,
		email: email,
		numberOfTickets: userTickets,
	}

	bookings = append(bookings, userData)
	fmt.Printf("List of bookings is %v\n", bookings)

	fmt.Printf("Thank you %v %v for booking %v tickets. You will receive a confirmation email at %v\n",
		firstName, lastName,
		userTickets, email)
	fmt.Printf("%v tickets remaining for %v\n", RemainingTickets, conferenceName)
}


// generating ticket and sending it to email after booking
func sendTicket(userTickets int, firstName string, lastName string, email string) {
	var ticket = fmt.Sprintf("%v tickets for %v %v", userTickets, firstName, lastName)
	fmt.Println("###############")
	fmt.Printf("Sending ticket:\n %v \nto email address %v\n", ticket, email)
	fmt.Println("###############")
	return
}
