package filemanager

import (
	"encoding/csv"
	"gobases/Desafio/internal/ticket"
	"os"
)

type FileManager struct {
	FilePath string
}

func NewFileManager(filePath string) *FileManager {
	return &FileManager{FilePath: filePath}
}

func (fm *FileManager) CreateTicketsListByFile() ([]ticket.Ticket, error) {
	file, err := os.Open(fm.FilePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	reader := csv.NewReader(file)
	lines, err := reader.ReadAll()
	if err != nil {
		return nil, err
	}

	tickets := []ticket.Ticket{}
	for _, line := range lines {
		ticket, err := ticket.CreateTicketByLine(line)
		if err != nil {
			return nil, err
		}

		tickets = append(tickets, ticket)
	}

	return tickets, nil
}
