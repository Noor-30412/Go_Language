package main

import (
	"fmt"
	"strings"
)

func main() {
	vehicalName := "Bus"      // value can be changed for var
	const vehicalTickets = 50 // value cannot be changed for const
	var remainingTickets uint = 50

	var bookings []string

	fmt.Printf("Welcome To My %v Booking Application\n", vehicalName)
	fmt.Println("We have a total of ", vehicalTickets, "tickets and", remainingTickets, "are still available.")
	fmt.Println("   Get Your Tickets Here    ")

	for remainingTickets > 0 {
		var userName string
		var userBooked uint
		var userMail string
		var userMobile string

		fmt.Println("Enter Your Name:")
		fmt.Scan(&userName)
		fmt.Println("Enter No of Tickets:")
		fmt.Scan(&userBooked)

		if userBooked > 50 {
			fmt.Printf("We only have %v tickets remaining\n", remainingTickets)
			continue
		}
		fmt.Println("Enter Your Email:")
		fmt.Scan(&userMail)
		fmt.Println("Enter Your Mobile Number:")
		fmt.Scan(&userMobile)

		var isValidName = len(userName) >= 2
		isValidEmail := strings.Contains(userMail, "@")

		if isValidEmail && isValidName && userBooked <= remainingTickets {
			remainingTickets = remainingTickets - userBooked
			bookings = append(bookings, userName)

			fmt.Printf("Thank you %v you Booked %v Tickets. You will receive confirmation @ %v and Mobile %v\n", userName, userBooked, userMail, userMobile)
			fmt.Printf("%v tickets remaining for %v\n", remainingTickets, vehicalName)

			fmt.Printf("The People booked Tickets are %v\n", bookings)
			fmt.Printf("The No Of People booked Tickets are %v\n", len(bookings))

			if remainingTickets == 0 {
				fmt.Println("The Bus is Full")
				break
			}
		} else if !isValidEmail || !isValidName {
			fmt.Println("Invalid Input")
		}
	}
}
