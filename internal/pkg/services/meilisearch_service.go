package services

import (
	"github.com/goccy/go-json"
	meili "github.com/meilisearch/meilisearch-go"
	m "squad10x.com.br/boilerplate/internal/pkg/models"
	"squad10x.com.br/boilerplate/pkg/configs"
)

type MeilisearchService interface {
	SearchPokemon(term string) ([]m.Pokemon, error)
	IndexPokemon(docs map[string]interface{}) error
}

type meilisearchService struct {
	client meili.Client
}

func NewMeilisearchService() MeilisearchService {
	return &meilisearchService{
		client: *meili.NewClient(meili.ClientConfig{
			Host:   configs.GetEnv("MEILISEARCH_HOST"),
			APIKey: configs.GetEnv("MEILISEARCH_API_KEY"),
		}),
	}
}

func (s *meilisearchService) SearchPokemon(term string) ([]m.Pokemon, error) {
	index := s.client.Index("pokemons")

	sort := []string{
		"attack:desc",
	}

	result, err := index.Search(term, &meili.SearchRequest{
		Sort: sort,
	})
	if err != nil {
		return []m.Pokemon{}, err
	}

	pokemons := []m.Pokemon{}
	for i := 0; i < len(result.Hits); i++ {
		var p m.Pokemon
		t, _ := json.Marshal(result.Hits[i])
		json.Unmarshal(t, &p)
		pokemons = append(pokemons, p)
	}

	return pokemons, err
}

func (s *meilisearchService) IndexPokemon(docs map[string]interface{}) error {
	index := s.client.Index("pokemons")

	_, err := index.AddDocuments(docs, "id")
	if err != nil {
		return err
	}

	return nil
}
