package main

type Movie struct {
	id			  int
    Name            string
    Price           float64
    AvailableTickets int
}

type boughtTickets struct {
	id	int
	Name	string
	email	string
	noOfTickets	int
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

var boughtTickets []boughtTickets

func main() {
	ticketBookingApp()
}