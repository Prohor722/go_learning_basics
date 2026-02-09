package main

var movies := []Movie{
        {
            Name:             "The Matrix",
            Price:            12.99,
            AvailableTickets: 50,
        },
        {
            Name:             "Inception",
            Price:            14.99,
            AvailableTickets: 30,
        },
        {
            Name:             "Interstellar",
            Price:            15.99,
            AvailableTickets: 25,
        },
    }

func main() {
	ticketBookingApp()
}