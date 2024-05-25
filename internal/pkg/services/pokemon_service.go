package services

import (
	"fmt"
	"io"
	"net/http"

	"github.com/goccy/go-json"
	"squad10x.com.br/boilerplate/internal/pkg/models"
)

type PokemonService interface {
	GetById(id int) (models.Pokemon, error)
}

type pokemonService struct {
	client *http.Client
}

func NewPokemonService() PokemonService {
	return &pokemonService{
		client: &http.Client{},
	}
}

func (s *pokemonService) GetById(id int) (models.Pokemon, error) {
	url := fmt.Sprintf("https://pokeapi.co/api/v2/pokemon/%d", id)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return models.Pokemon{}, err
	}

	req.Header.Add("Content-Type", "application/json")
	resp, err := s.client.Do(req)
	if err != nil {
		return models.Pokemon{}, err
	}

	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return models.Pokemon{}, err
	}

	result := models.FullPokemonData{}
	err = json.Unmarshal(respBody, &result)
	if err != nil {
		return models.Pokemon{}, err
	}

	return result.ToSlim(), err
}
