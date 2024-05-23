package handlers

import (
	"squad10x.com.br/boilerplate/internal/app/api/requests"
	m "squad10x.com.br/boilerplate/internal/pkg/models"
)

type generalHandler struct{}

func NewGeneralHandler() *generalHandler {
	return &generalHandler{}
}

func (h *generalHandler) Search(req requests.SearchRequest) (m.Pokemon, error) {
	return m.Pokemon{}, nil
}
