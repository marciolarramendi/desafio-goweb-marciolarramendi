package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/marciolarramendi/desafio-goweb-marciolarramendi/internal/tickets"
)

type Service struct {
	service tickets.Service
}

func NewService(s tickets.Service) *Service {
	return &Service{
		service: s,
	}
}

func (s *Service) GetAll() gin.HandlerFunc {
	return func(c *gin.Context) {
		tickets, err := s.service.GetAll()
		if err != nil {
			c.JSON(404, err.Error())
			return
		}
		c.JSON(200, tickets)
	}
}

func (s *Service) GetTicketsByCountry() gin.HandlerFunc {
	return func(c *gin.Context) {

		destination := c.Param("dest")

		tickets, err := s.service.GetTicketByDestination(destination)
		if err != nil {
			c.String(http.StatusInternalServerError, err.Error(), nil)
			return
		}

		c.JSON(200, tickets)
	}
}

func (s *Service) AverageDestination() gin.HandlerFunc {
	return func(c *gin.Context) {

		destination := c.Param("dest")

		avg, err := s.service.AverageDestination(destination)
		if err != nil {
			c.String(http.StatusInternalServerError, err.Error(), nil)
			return
		}

		c.JSON(200, avg)
	}
}
