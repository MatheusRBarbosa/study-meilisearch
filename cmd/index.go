package cmd

import (
	"github.com/goccy/go-json"
	"github.com/spf13/cobra"
	"squad10x.com.br/boilerplate/internal/pkg/services"
	"squad10x.com.br/boilerplate/pkg/logger"
)

var (
	indexCommand = &cobra.Command{
		Use:   "index",
		Short: "Create initial pokemon index",
		Long:  "Make a lot of requests to pokemon API, then index the responses",
		RunE:  indexExecute,
	}
)

func init() {
	rootCmd.AddCommand(indexCommand)
}

func indexExecute(_ *cobra.Command, args []string) error {
	pokeapi := services.NewPokemonService()
	meili := services.NewMeilisearchService()
	logger := logger.GetLogger()

	page := 1
	pageSize := 1000

	offset := (page - 1) * pageSize
	for i := offset; i < pageSize; i++ {
		pokemon, err := pokeapi.GetById(i + 1)
		if err != nil {
			logger.Error(err)
		} else {
			logger.Infof("Indexing >> %s", pokemon.Name)
			var toIndex map[string]interface{}
			t, _ := json.Marshal(pokemon)
			json.Unmarshal(t, &toIndex)

			err = meili.IndexPokemon(toIndex)
			if err != nil {
				logger.Error(err)
			}
		}
	}

	return nil
}
