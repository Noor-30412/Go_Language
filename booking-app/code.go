package main

import "fmt"

const (
	vehicleName      = "Bus"
	initialTickets   = 50
	remainingTickets = initialTickets
)

type Booking struct {
	UserName   string
	UserBooked uint
	UserMail   string
	UserMobile string
}

func code() {
	// Print welcome message
	fmt.Printf("Welcome To My %v Booking Application\n", vehicleName)
	fmt.Printf("We have a total of %v tickets available.\n", initialTickets)

	// Slice to store bookings
	var bookings []Booking

	// Booking loop
	for {
		var user Booking

		// Get user details
		fmt.Print("Enter Your Name: ")
		fmt.Scan(&user.UserName)

		fmt.Print("Enter No of Tickets: ")
		fmt.Scan(&user.UserBooked)

		// Check if there are enough tickets
		if user.UserBooked <= remainingTickets {
			// Update remaining tickets and add booking
			remainingTickets := user.UserBooked
			bookings = append(bookings, user)

			// Display booking details
			fmt.Printf("Thank you %v, you booked %v tickets. You will receive confirmation @ %v and Mobile %v\n",
				user.UserName, user.UserBooked, user.UserMail, user.UserMobile)
			fmt.Printf("%v tickets remaining for %v\n", remainingTickets, vehicleName)

			// Display booked people and count
			fmt.Printf("People who booked tickets: %v\n", bookings)
			fmt.Printf("Number of people who booked tickets: %v\n", len(bookings))

			// Check if the bus is full
			if remainingTickets == 0 {
				fmt.Println("The Bus is Full")
				break
			}
		} else {
			fmt.Printf("We only have %v tickets remaining\n", remainingTickets)
		}

		// Get additional user details
		fmt.Print("Enter Your Email: ")
		fmt.Scan(&user.UserMail)

		fmt.Print("Enter Your Mobile Number: ")
		fmt.Scan(&user.UserMobile)
	}
}
