package store

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"os"
	"strconv"

	"github.com/marciolarramendi/desafio-goweb-marciolarramendi/internal/domain"
)

type Store interface {
	Read(data interface{}) ([]domain.Ticket, error)
	Write(data interface{}) error
}

type Type string

const (
	FileType  Type = "file"
	MongoType Type = "Mongo"
)

func New(store Type, fileName string) Store {
	switch store {
	case FileType:
		return &fileStore{fileName}
	}
	return nil
}

type fileStore struct {
	FilePath string
}

func (fs *fileStore) Read(data interface{}) ([]domain.Ticket, error) {
	var ticketList []domain.Ticket

	file, err := os.Open(fs.FilePath)
	if err != nil {
		return nil, fmt.Errorf("could not open file: %w", err)
	}

	csvR := csv.NewReader(file)
	tickets, err := csvR.ReadAll()
	if err != nil {
		return nil, fmt.Errorf("could not open file: %w", err)
	}

	for _, row := range tickets {
		price, err := strconv.ParseFloat(row[5], 64)
		if err != nil {
			return []domain.Ticket{}, err
		}
		ticketList = append(ticketList, domain.Ticket{
			Id:      row[0],
			Name:    row[1],
			Email:   row[2],
			Country: row[3],
			Time:    row[4],
			Price:   price,
		})
	}

	return ticketList, nil
}

func (fs *fileStore) Write(data interface{}) error {
	fileData, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		return err
	}
	return os.WriteFile(fs.FilePath, fileData, 0644)
}
