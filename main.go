package main

import (
	"fmt"
)

// Where to start the program - compiler signal
func main() {
	var conferenceName = "Go Conference"
	const conferenceTicketsNum = 50
	var remainingTickets = conferenceTicketsNum

	fmt.Printf("Welcome to the %v booking application.\n", conferenceName)
	fmt.Print("There are ", remainingTickets, " tickets remaining, out of a total of ", remainingTickets, " tickets. \n")
	fmt.Print("Get your tickets here to attend.")
}
