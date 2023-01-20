package main

import (
	"Booking-App/helper"
	"fmt"
	"time"
)

const conferenceticket = 50

var conferencename = "Go Conference"
var remainingtickets uint = 50
var bookings = make([]userdata, 0)

type userdata struct {
	firstname     string
	lastname      string
	email         string
	total_tickets int
}

func main() {

	greeting()

	for {
		firstname, lastname, email, usertickets := userinput()
		isValidname, isValidemail, isValidticketNumber := helper.Isvalid(firstname, lastname, usertickets, email, remainingtickets)

		if isValidname && isValidemail && isValidticketNumber {

			booking(firstname, lastname, usertickets, email)
			go sendtickets(usertickets, firstname, lastname, email)
			firstname := Printfirstname()
			fmt.Printf("These are the people who booked tickets for %v %v\n", conferencename, firstname)

		} else if remainingtickets == 0 {
			fmt.Println("All the tickets are booked try again next year")

		} else {
			if !isValidemail {
				fmt.Println("your email is invalid")

			} else if !isValidname {
				fmt.Println("your firstname or lastname are invalid")
			} else {
				fmt.Println("your are booking more number of tickets or your are not fill the booking column")
			}

		}
	}

}
func greeting() {

	fmt.Printf("Welcome to %v of booking application\n", conferencename)
	fmt.Printf("we have total of %v tickets and the remaining tickets are %v\n", conferenceticket, remainingtickets)
	fmt.Print("Get your tickets here to attend the ")
	fmt.Println(conferencename)

}
func Printfirstname() []string {

	var firstnames []string
	for _, booking := range bookings {
		firstnames = append(firstnames, booking.firstname)
	}
	return firstnames
}
func userinput() (string, string, string, int) {
	var firstname string
	var lastname string
	var email string
	var usertickets int

	fmt.Println("Enter your first name: ")
	fmt.Scan(&firstname)

	fmt.Println("Enter your last name:")
	fmt.Scan(&lastname)

	fmt.Println("Enter your email:")
	fmt.Scan(&email)

	fmt.Println("How many tickets do you want:")
	fmt.Scan(&usertickets)

	return firstname, lastname, email, usertickets

}
func booking(firstname string, lastname string, usertickets int, email string) {
	var userdata = userdata{
		firstname:     firstname,
		lastname:      lastname,
		email:         email,
		total_tickets: usertickets,
	}

	bookings = append(bookings, userdata)
	remainingtickets = remainingtickets - uint(usertickets)

	fmt.Printf("Thank you %v %v for booking %v tickets for %v and the tickets will send to your email %v\n", firstname, lastname, usertickets, conferencename, email)
	fmt.Println("The remaining tickets are", remainingtickets)
	fmt.Println(bookings)
}
func sendtickets(usertickets int, firstname string, lastname string, email string) {
	time.Sleep(10 * time.Second)
	var tickets = fmt.Sprintf("%v tickets for %v %v", usertickets, firstname, lastname)
	fmt.Printf("#######\n")
	fmt.Printf("sending tickets:\n %v \n to your email address %v \n", tickets, email)
	fmt.Printf("#######\n")
}
