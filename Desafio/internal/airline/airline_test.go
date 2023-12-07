package airline_test

import (
	"fmt"
	"gobases/Desafio/internal/airline"
	"gobases/Desafio/internal/ticket"
	"testing"

	"github.com/stretchr/testify/assert"
)

var tickets = []ticket.Ticket{
	{ID: 1, Name: "Juan Perez", Email: "jp@gmail.com", Destination: "China", FlightTime: "17:00", Price: 1000},
	{ID: 2, Name: "Maria Lopez", Email: "ml@gmail.com", Destination: "France", FlightTime: "12:00", Price: 1000},
	{ID: 3, Name: "Carlos Sanchez", Email: "cs@gmail.com", Destination: "China", FlightTime: "23:00", Price: 1000},
}

func TestGetTotalTicketsByDestination(t *testing.T) {
	t.Run("Validar cantidad de tickets con destino a China",
		func(t *testing.T) {
			newAirline := airline.NewAirline(tickets)
			expectedResult := 2

			obtainedResult, err := newAirline.GetTotalTicketsByDestination("China")
			if err != nil {
				fmt.Println(err)
			}

			assert.Equal(t, expectedResult, obtainedResult, "Debe devolver 2")
		},
	)

	t.Run("Validar error al no pasar destino",
		func(t *testing.T) {
			newAirline := airline.NewAirline(tickets)
			expectedResult := airline.ErrNonZeroValue

			_, err := newAirline.GetTotalTicketsByDestination("")
			if err != nil {
				fmt.Println(err)
			}

			assert.ErrorIs(t, err, expectedResult, "Debe devolver ErrNonZeroValue")
		},
	)
}

func TestGetTotalTicketsByPeriod(t *testing.T) {
	t.Run("Validar cantidad de tickets que viajan a la tarde",
		func(t *testing.T) {
			newAirline := airline.NewAirline(tickets)
			expectedResult := 1

			obtainedResult, err := newAirline.GetTotalTicketsByPeriod(airline.Afternoon)
			if err != nil {
				fmt.Println(err)
			}

			assert.Equal(t, expectedResult, obtainedResult, "Debe devolver 1")
		},
	)

	t.Run("Validar error al no pasar destino",
		func(t *testing.T) {
			newTicket := ticket.Ticket{
				ID: 4, Name: "Natalie Gonzalez", Email: "nr@gmail.com", Destination: "Italy", FlightTime: "1400", Price: 1000,
			}
			tickets = append(tickets, newTicket)
			newAirline := airline.NewAirline(tickets)
			expectedResult := airline.ErrTimeFormat

			_, err := newAirline.GetTotalTicketsByPeriod(airline.EarlyMorning)
			if err != nil {
				fmt.Println(err)
			}

			assert.ErrorIs(t, err, expectedResult, "Debe devolver ErrTimeFormat")
		},
	)
}

func TestGetPercentageTicketsByDestination(t *testing.T) {
	t.Run("Validar porcentaje de tickets con destino a Francia",
		func(t *testing.T) {
			newAirline := airline.NewAirline(tickets)
			expectedResult := 25

			obtainedResult, err := newAirline.GetPercentageTicketsByDestination("France")
			if err != nil {
				fmt.Println(err)
			}

			assert.Equal(t, expectedResult, obtainedResult, "Debe devolver 25")
		},
	)

	t.Run("Validar error al no tener tickets",
		func(t *testing.T) {
			newAirline := airline.NewAirline(make([]ticket.Ticket, 0))
			expectedResult := airline.ErrEmptyTicketList

			_, err := newAirline.GetTotalTicketsByDestination("")
			if err != nil {
				fmt.Println(err)
			}

			assert.ErrorIs(t, err, expectedResult, "Debe devolver ErrEmptyTicketList")
		},
	)
}
