package main

import "fmt"

func main() {
	vehicalName := "Bus"      //value can be change for var
	const vehicalTickets = 50 //value cannot be changed for const
	var remainingTickets uint = 50
	//Array initialization
	//var bookings [50]string

	//slice
	var bookings []string

	//fmt.Printf("Welcome To My %T Booking Application\n",vehicalName)
	fmt.Printf("Welcome To My %v Booking Application\n", vehicalName)
	fmt.Println("We have total of ", vehicalTickets, "tickets and", remainingTickets, "are still available.")
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
		fmt.Println("Enter Your Email:")
		fmt.Scan(&userMail)
		fmt.Println("Enter Your Mobile Number:")
		fmt.Scan(&userMobile)

		remainingTickets = remainingTickets - userBooked
		//bookings[0] = userName
		bookings = append(bookings, userName)

		fmt.Printf("Thank you %v you Booked %v Tickets. You will recieve confirmation @ %v and Mobile %v\n", userName, userBooked, userMail, userMobile)
		fmt.Printf("%v tickets remaining for %v\n", remainingTickets, vehicalName)
		/*
			firstNames:= []string{}
			import "strings"
			for _,booking := range booking{
				names := strings.Fields(booking) //use to split words with space
				firstNames = append(firstNames,names[])
			}
		*/
		fmt.Printf("The People booked Tickets are %v\n", bookings)
		fmt.Printf("The No Of People booked Tickets are %v\n", len(bookings))

		if remainingTickets == 0 {
			fmt.Println("The Bus is Full")
			break
		}
	}

}
// PS C:\Users\N M K\Desktop\SW\Golang> git remote -v
// origin  https://github.com/Noor-30412/Go_Language.git (fetch)
// origin  https://github.com/Noor-30412/Go_Language.git (push)

// git remote set-url origin https://github.com/Noor-30412/Go_Language.git
