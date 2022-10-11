package main

import (
	"encoding/csv"
	"fmt"
	"os"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/marciolarramendi/desafio-goweb-marciolarramendi/cmd/server/handler"
	"github.com/marciolarramendi/desafio-goweb-marciolarramendi/internal/domain"
	"github.com/marciolarramendi/desafio-goweb-marciolarramendi/internal/tickets"
)

func main() {
	// Cargo csv.
	list, err := LoadTicketsFromFile("../../tickets.csv")
	if err != nil {
		panic("Couldn't load tickets")
	}
	//db := store.New(store.FileType, "./tickets.csv")
	repo := tickets.NewRepository(list)
	service := tickets.NewService(repo)

	u := handler.NewService(service)

	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) { c.String(200, "pong") })
	// Rutas a desarollar:

	ur := r.Group("/ticket")
	ur.GET("/", u.GetAll())
	ur.GET("/getByCountry/:dest", u.GetTicketsByCountry())
	ur.GET("/getAverage/:dest", u.AverageDestination())

	// GET - “/ticket/getByCountry/:dest”
	// GET - “/ticket/getAverage/:dest”
	if err := r.Run(); err != nil {
		panic(err)
	}

}

func LoadTicketsFromFile(path string) ([]domain.Ticket, error) {

	var ticketList []domain.Ticket

	file, err := os.Open(path)
	if err != nil {
		return nil, fmt.Errorf("could not open file: %w", err)
	}

	csvR := csv.NewReader(file)
	data, err := csvR.ReadAll()
	if err != nil {
		return nil, fmt.Errorf("could not open file: %w", err)
	}

	for _, row := range data {
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