package main

import (
	"fmt"
	"sync"
	"time"

	"example.com/golang/helper"
)

//Package level variables
var conferenceName string = "Go conference"
const conferenceTicket uint8 = 50 
var remainingTickets uint8 = 50
var firstName string
var lastName string
var ticketBought uint8
var email string
var bookings = make([]UserData,0)

// user data struct
type UserData struct {
	firstName 	string
	lastName	string
	email		string
	ticketBought uint8
}


//sync to wait for thread
var wg = sync.WaitGroup{}

func main(){
	// Greeting users
	greetUsers()

	//Get user input
	firstName, lastName, email, ticketBought := getUserInput()
	
	//validate user input
	isEmailValid, isValidName := helper.InputValidator(firstName,lastName,email)
	
	// checking condition for user details 
	if ticketBought <= remainingTickets && isValidName && isEmailValid {
		remainingTickets -= ticketBought
		// creating user object
		var userData = UserData {
			firstName: firstName,
			lastName: lastName,
			email: email,
			ticketBought: ticketBought,
		}

		wg.Add(1)

		// Goroutine
		go sendTicket(ticketBought,firstName,lastName)

		
		// adding objects to list
		bookings = append(bookings, userData)	

		fmt.Printf("List of bookings: %v \n",bookings)
		fmt.Printf("Thank you %v %v, for booking %v tickets. You will receive a confirmation email at %v \n", firstName, lastName, ticketBought,email)
		fmt.Printf("we have total of %v tickets and the remaining tickets are %v \n", conferenceTicket,remainingTickets)

		fmt.Printf("Number of users who booked are %v\n", len(bookings))
		
		fmt.Printf("Booked Users: %v\n", getUserName(bookings))

		noOfTicketRemaining := remainingTickets == 0

		if noOfTicketRemaining{
			fmt.Println("Our conferences is booked out. come back next year")
			// break
		}
	}else{
		if !isValidName{
			fmt.Println("Invalid name length")
		}
		if !isEmailValid{
			fmt.Println("Invalid email")
		}
		fmt.Printf("The remaining ticket is  %v\n", remainingTickets)

	}

	// city := "London"

	// switch city {
	// 	case "New York":
	// 		//code
	// 	case "Hong Kong":
	// 		//
	// 	default:

	
	// }

	wg.Wait()
}

func greetUsers(){
	fmt.Printf("Welcome %v to our conference", firstName)
	fmt.Printf("Welcome to %v booking app\n",conferenceName)
	fmt.Printf("we have total of %v tickets and the remaining tickets are %v \n", conferenceTicket,remainingTickets)
	fmt.Println("You can book all you tickets using this application")
}

func getUserName(bookings []UserData) []string {
	firstNames := []string{}
	for _, booking := range bookings{
		firstNames = append(firstNames, booking.firstName)
	}
	return firstNames
}

func getUserInput() (string,string,string,uint8){
	fmt.Println("Enter first name: ")
	fmt.Scan(&firstName)

	fmt.Println("Enter last name: ")
	fmt.Scan(&lastName)

	fmt.Println("Enter email: ")
	fmt.Scan(&email)

	fmt.Println("How many tickets do you want to buy: ")
	fmt.Scan(&ticketBought)
	return firstName,lastName,email,ticketBought
}

func sendTicket(ticketBought uint8, firstName string, lastName string){
	time.Sleep(10 * time.Second)
	var ticket = fmt.Sprintf("%v tickets for %v %v", ticketBought, firstName, lastName)
	fmt.Println("#######################")
	fmt.Printf("Sending %v  to %v\n", ticket, firstName)
	fmt.Println("#######################")
	wg.Done()
}




