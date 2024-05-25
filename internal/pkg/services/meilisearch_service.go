package services

import (
	"github.com/meilisearch/meilisearch-go"
	"squad10x.com.br/boilerplate/pkg/configs"
)

type MeilisearchService interface {
	SearchPokemon(term string) meilisearch.DocumentsResult
	IndexPokemon(docs map[string]interface{}) error
}

type meilisearchService struct {
	client meilisearch.Client
}

func NewMeilisearchService() MeilisearchService {
	return &meilisearchService{
		client: *meilisearch.NewClient(meilisearch.ClientConfig{
			Host:   configs.GetEnv("MEILISEARCH_HOST"),
			APIKey: configs.GetEnv("MEILISEARCH_API_KEY"),
		}),
	}
}

func (s *meilisearchService) SearchPokemon(term string) meilisearch.DocumentsResult {
	return meilisearch.DocumentsResult{}
}

func (s *meilisearchService) IndexPokemon(docs map[string]interface{}) error {
	index := s.client.Index("pokemons")

	_, err := index.AddDocuments(docs, "id")
	if err != nil {
		return err
	}

	return nil
}
