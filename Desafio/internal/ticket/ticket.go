package ticket

import (
	"errors"
	"strconv"
)

var (
	ErrTicketFormat = errors.New("the string ticket does not have the correct format")
)

type Ticket struct {
	ID          int
	Name        string
	Email       string
	Destination string
	FlightTime  string
	Price       int
}

func CreateTicketByLine(line []string) (Ticket, error) {
	if len(line) != 6 {
		return Ticket{}, ErrTicketFormat
	}

	id, err := strconv.Atoi(line[0])
	if err != nil {
		return Ticket{}, err
	}
	price, err := strconv.Atoi(line[5])
	if err != nil {
		return Ticket{}, err
	}

	newTicket := Ticket{
		ID:          id,
		Name:        line[1],
		Email:       line[2],
		Destination: line[3],
		FlightTime:  line[4],
		Price:       price,
	}

	return newTicket, nil
}
