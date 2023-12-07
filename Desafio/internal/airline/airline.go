package airline

import (
	"errors"
	"gobases/Desafio/internal/ticket"
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
	ErrNonZeroValue    = errors.New("parameters must be non-zero values")
	ErrEmptyTicketList = errors.New("there are not any tickets")
	ErrTimeFormat      = errors.New("the string time does not have the correct format")
	ErrInvalidTime     = errors.New("invalid time")
)

type Airline struct {
	Tickets []ticket.Ticket
}

func NewAirline(tickets []ticket.Ticket) *Airline {
	return &Airline{Tickets: tickets}
}

func (a *Airline) GetTotalTicketsByDestination(destination string) (total int, err error) {
	if a.Tickets == nil || destination == "" {
		return 0, ErrNonZeroValue
	}

	if len(a.Tickets) == 0 {
		return 0, ErrEmptyTicketList
	}

	for _, ticket := range a.Tickets {
		if ticket.Destination == destination {
			total++
		}
	}

	return total, nil
}

func (a *Airline) GetTotalTicketsByPeriod(time string) (total int, err error) {
	if a.Tickets == nil || time == "" {
		return 0, ErrNonZeroValue
	}

	if len(a.Tickets) == 0 {
		return 0, ErrEmptyTicketList
	}

	for _, ticket := range a.Tickets {
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

func (a *Airline) GetPercentageTicketsByDestination(destination string) (percentage float64, err error) {
	if a.Tickets == nil || destination == "" {
		return 0, ErrNonZeroValue
	}

	if len(a.Tickets) == 0 {
		return 0, ErrEmptyTicketList
	}

	total := 0
	for _, ticket := range a.Tickets {
		if ticket.Destination == destination {
			total++
		}
	}

	percentage = float64(total) * 100 / float64(len(a.Tickets))

	return
}
