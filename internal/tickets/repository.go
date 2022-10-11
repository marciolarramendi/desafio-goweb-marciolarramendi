package tickets

import (
	"fmt"

	"github.com/marciolarramendi/desafio-goweb-marciolarramendi/internal/domain"
)

var tickets []domain.Ticket

type Repository interface {
	GetAll() ([]domain.Ticket, error)
	GetTicketByDestination(destination string) ([]domain.Ticket, error)
	GetAverageDestination(destination string) (float64, error)
}

type repository struct {
	db []domain.Ticket
}

func NewRepository(db []domain.Ticket) Repository {
	return &repository{
		db: db,
	}
}

func (r *repository) GetAll() ([]domain.Ticket, error) {
	if len(r.db) == 0 {
		return []domain.Ticket{}, fmt.Errorf("empty list of tickets")
	}
	return r.db, nil
}

func (r *repository) GetTicketByDestination(destination string) ([]domain.Ticket, error) {
	var ticketsDest []domain.Ticket

	if len(r.db) == 0 {
		return []domain.Ticket{}, fmt.Errorf("empty list of tickets")
	}

	for _, t := range r.db {
		if t.Country == destination {
			ticketsDest = append(ticketsDest, t)
		}
	}

	return ticketsDest, nil
}

func (r *repository) GetAverageDestination(destination string) (float64, error) {

	//var ticketsDest []domain.Ticket
	var count int

	if len(r.db) == 0 {
		return 0.0, fmt.Errorf("empty list of tickets")
	}

	totalRecords := len(r.db)

	for _, t := range r.db {
		if t.Country == destination {
			count++
		}
	}
	avg := float64(count*100) / float64(totalRecords)

	return avg, nil
}
