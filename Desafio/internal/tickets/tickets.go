package tickets

import (
	"bufio"
	"errors"
	"os"
	"strconv"
	"strings"
)

const (
	EarlyMorning = "Madrugada"
	Morning      = "Mañana"
	Afternoon    = "Tarde"
	Night        = "Noche"
)

var (
	ErrTicketFormat    = errors.New("the string ticket does not have the correct format")
	ErrNonZeroValue    = errors.New("parameters must be non-zero values")
	ErrEmptyTicketList = errors.New("there are not any tickets")
	ErrTimeFormat      = errors.New("the string time does not have the correct format")
	ErrInvalidTime     = errors.New("invalid time")
)

type Ticket struct {
	ID          int
	Name        string
	Email       string
	Destination string
	FlightTime  string
	Price       int
}

func CreateTicketsListByFile(path string) ([]Ticket, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	tickets := []Ticket{}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		ticket, err := CreateTicketByString(scanner.Text())
		if err != nil {
			return nil, err
		}

		tickets = append(tickets, ticket)
	}

	return tickets, nil
}

func CreateTicketByString(stringTicket string) (Ticket, error) {
	slicedTicket := strings.Split(stringTicket, ",")
	if len(slicedTicket) != 6 {
		return Ticket{}, ErrTicketFormat
	}

	id, err := strconv.Atoi(slicedTicket[0])
	if err != nil {
		return Ticket{}, err
	}
	price, err := strconv.Atoi(slicedTicket[5])
	if err != nil {
		return Ticket{}, err
	}

	newTicket := Ticket{
		ID:          id,
		Name:        slicedTicket[1],
		Email:       slicedTicket[2],
		Destination: slicedTicket[3],
		FlightTime:  slicedTicket[4],
		Price:       price,
	}

	return newTicket, nil
}

func GetTotalTicketsByDestination(tickets []Ticket, destination string) (total int, err error) {
	if tickets == nil || destination == "" {
		return 0, ErrNonZeroValue
	}

	if len(tickets) == 0 {
		return 0, ErrEmptyTicketList
	}

	for _, ticket := range tickets {
		if ticket.Destination == destination {
			total++
		}
	}

	return total, nil
}

func GetTotalTicketsByPeriod(tickets []Ticket, time string) (total int, err error) {
	if tickets == nil || time == "" {
		return 0, ErrNonZeroValue
	}

	if len(tickets) == 0 {
		return 0, ErrEmptyTicketList
	}

	for _, ticket := range tickets {
		slicedTime := strings.Split(ticket.FlightTime, ":")
		if len(slicedTime) != 2 {
			return 0, ErrTimeFormat
		}

		hour, err := strconv.Atoi(slicedTime[0])
		if err != nil {
			return 0, err
		}

		switch time {
		case EarlyMorning:
			if hour >= 0 && hour <= 6 {
				total++
			}
		case Morning:
			if hour >= 7 && hour <= 12 {
				total++
			}
		case Afternoon:
			if hour >= 13 && hour <= 19 {
				total++
			}
		case Night:
			if hour >= 20 && hour <= 23 {
				total++
			}
		default:
			return 0, ErrInvalidTime
		}
	}

	return
}

func GetAverageTicketsByDestination(tickets []Ticket, destination string) (average float64, err error) {
	if tickets == nil || destination == "" {
		return 0, ErrNonZeroValue
	}

	if len(tickets) == 0 {
		return 0, ErrEmptyTicketList
	}

	total := 0
	for _, ticket := range tickets {
		if ticket.Destination == destination {
			total++
		}
	}

	average = float64(total) / float64(len(tickets)) * 100

	return
}
