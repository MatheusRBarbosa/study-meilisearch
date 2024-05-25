package handlers

import (
	m "squad10x.com.br/boilerplate/internal/pkg/models"
	"squad10x.com.br/boilerplate/internal/pkg/services"
)

type generalHandler struct {
	meili services.MeilisearchService
}

func NewGeneralHandler() *generalHandler {
	return &generalHandler{
		meili: services.NewMeilisearchService(),
	}
}

func (h *generalHandler) Search(term string) ([]m.Pokemon, error) {
	result, err := h.meili.SearchPokemon(term)

	if err != nil {
		return []m.Pokemon{}, err
	}

	return result, nil
}
