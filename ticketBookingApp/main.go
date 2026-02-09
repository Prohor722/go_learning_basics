package main

type Movie struct {
    Name            string
    Price           float64
    AvailableTickets int
}

var movies = []Movie{
        {
			id:			   1,
            Name:             "The Matrix",
            Price:            12.99,
            AvailableTickets: 50,
        },
        {
			id:			   2,
            Name:             "Inception",
            Price:            14.99,
            AvailableTickets: 30,
        },
        {
			id:			   3,
            Name:             "Interstellar",
            Price:            15.99,
            AvailableTickets: 25,
        },
    }

func main() {
	ticketBookingApp()
}