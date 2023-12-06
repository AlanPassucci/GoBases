package main

import (
	"fmt"
	"gobases/Desafio/internal/tickets"
)

func main() {
	fmt.Println("Creando lista de tickets...")
	ticketsList, err := tickets.CreateTicketsListByFile("./tickets.csv")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("¡Lista de tickets creada!")

	fmt.Println("\nCalculando total de tickets con destino a China...")
	totalChinaTickets, err := tickets.GetTotalTicketsByDestination(ticketsList, "China")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Total de tickets con destino a China:", totalChinaTickets)

	fmt.Println("\nCalculando total de viajes durante la noche...")
	totalTimeTickets, err := tickets.GetTotalTicketsByPeriod(ticketsList, tickets.Night)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Total de viajes durante la noche:", totalTimeTickets)

	fmt.Println("\nCalculando promedio de viajes a China...")
	averageTickets, err := tickets.GetAverageTicketsByDestination(ticketsList, "China")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Promedio de viajes a China:", averageTickets, "%")
}
