package main

import "fmt"

func main() {
	vehicalName := "Bus"      // value can be changed for var
	const vehicalTickets = 50 // value cannot be changed for const
	var remainingTickets uint = 50
	// Array initialization
	// var bookings [50]string

	// Slice
	var bookings []string

	// fmt.Printf("Welcome To My %T Booking Application\n",vehicalName)
	fmt.Printf("Welcome To My %v Booking Application\n", vehicalName)
	fmt.Println("We have a total of ", vehicalTickets, "tickets and", remainingTickets, "are still available.")
	fmt.Println("   Get Your Tickets Here    ")

	for {
		var userName string
		var userBooked uint
		var userMail string
		var userMobile string

		fmt.Println("Enter Your Name:")
		fmt.Scan(&userName)
		fmt.Println("Enter No of Tickets:")
		fmt.Scan(&userBooked)

		if userBooked <= remainingTickets {
			remainingTickets = remainingTickets - userBooked
			// bookings[0] = userName
			bookings = append(bookings, userName)

			fmt.Printf("Thank you %v you Booked %v Tickets. You will receive confirmation @ %v and Mobile %v\n", userName, userBooked, userMail, userMobile)
			fmt.Printf("%v tickets remaining for %v\n", remainingTickets, vehicalName)

			fmt.Printf("The People booked Tickets are %v\n", bookings)
			fmt.Printf("The No Of People booked Tickets are %v\n", len(bookings))

			// var noTickets bool = remainingTickets == 0
			// noTickets := remainingTickets == 0
			if remainingTickets == 0 {
				fmt.Println("The Bus is Full")
				break
			}
		} else {
			fmt.Printf("We only have %v tickets remaining\n", remainingTickets)
			continue
		}
		fmt.Println("Enter Your Email:")
		fmt.Scan(&userMail)
		fmt.Println("Enter Your Mobile Number:")
		fmt.Scan(&userMobile)
	}
}
