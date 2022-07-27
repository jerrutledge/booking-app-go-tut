package main

import (
	"fmt"
)

// Where to start the program - compiler signal
func main() {
	var conferenceName = "Go Conference"
	const conferenceTicketsNum = 50
	var remainingTickets = conferenceTicketsNum

	fmt.Print("Welcome to the " + conferenceName + " booking application.\n")
	fmt.Print("There are ", remainingTickets, " tickets remaining, out of a total of ", remainingTickets, " tickets. \n")
	fmt.Print("Get your tickets here to attend.")
}
