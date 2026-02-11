package main

type Movie struct {
	id			  int
    Name            string
    Price           float64
    AvailableTickets int
    ageLimit       int
}

type BoughtTickets struct {
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
            ageLimit:        13,
        },
        {
			id:			   2,
            Name:             "Inception",
            Price:            14.99,
            AvailableTickets: 30,
            ageLimit:        16,
        },
        {
			id:			   3,
            Name:             "Interstellar",
            Price:            15.99,
            AvailableTickets: 25,
            ageLimit:        12,
        },
        {
            id:			   4,
            Name:             "The Dark Knight",
            Price:            13.99,
            AvailableTickets: 40,
            ageLimit:        13,
        },
        {
            id:			   5,
            Name:             "Titanic",
            Price:            12.99,
            AvailableTickets: 35,
            ageLimit:        15,
        },
        {
            id:			   6,
            Name:             "50 Shades of Grey",
            Price:            20.99,
            AvailableTickets: 20,
            ageLimit:        18,
        },
    }

var boughtTickets []BoughtTickets

func main() {
	ticketBookingApp()
}