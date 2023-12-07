package main

import (
	"fmt"
	"gobases/Desafio/internal/airline"
	"gobases/Desafio/internal/filemanager"
)

func main() {
	fileManager := filemanager.NewFileManager("./tickets.csv")
	fmt.Println("Creando lista de tickets...")
	tickets, err := fileManager.CreateTicketsListByFile()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("¡Lista de tickets creada!")

	newAirline := airline.NewAirline(tickets)
	fmt.Println("\nCalculando total de tickets con destino a Brasil...")
	totalBrazilTickets, err := newAirline.GetTotalTicketsByDestination("Brazil")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Total de tickets con destino a Brasil:", totalBrazilTickets)

	fmt.Println("\nCalculando total de viajes durante la noche...")

	totalTimeTickets, err := newAirline.GetTotalTicketsByPeriod(airline.Night)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Total de viajes durante la noche:", totalTimeTickets)

	fmt.Println("\nCalculando promedio de viajes a Brasil...")
	percentageTickets, err := newAirline.GetPercentageTicketsByDestination("Brazil")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Promedio de viajes a Brasil:", percentageTickets, "%")
}
