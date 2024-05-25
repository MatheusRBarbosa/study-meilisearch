package models

type FullPokemonData struct {
	Abilities []struct {
		Ability struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"ability"`
		IsHidden bool `json:"is_hidden"`
		Slot     int  `json:"slot"`
	} `json:"abilities"`
	BaseExperience int `json:"base_experience"`
	Cries          struct {
		Latest string      `json:"latest"`
		Legacy interface{} `json:"legacy"`
	} `json:"cries"`
	Forms []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"forms"`
	GameIndices            []interface{} `json:"game_indices"`
	Height                 int           `json:"height"`
	HeldItems              []interface{} `json:"held_items"`
	ID                     int           `json:"id"`
	IsDefault              bool          `json:"is_default"`
	LocationAreaEncounters string        `json:"location_area_encounters"`
	Moves                  []struct {
		Move struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"move"`
		VersionGroupDetails []struct {
			LevelLearnedAt  int `json:"level_learned_at"`
			MoveLearnMethod struct {
				Name string `json:"name"`
				URL  string `json:"url"`
			} `json:"move_learn_method"`
			VersionGroup struct {
				Name string `json:"name"`
				URL  string `json:"url"`
			} `json:"version_group"`
		} `json:"version_group_details"`
	} `json:"moves"`
	Name          string        `json:"name"`
	Order         int           `json:"order"`
	PastAbilities []interface{} `json:"past_abilities"`
	PastTypes     []interface{} `json:"past_types"`
	Species       struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"species"`
	Sprites struct {
		BackDefault      interface{} `json:"back_default"`
		BackFemale       interface{} `json:"back_female"`
		BackShiny        interface{} `json:"back_shiny"`
		BackShinyFemale  interface{} `json:"back_shiny_female"`
		FrontDefault     string      `json:"front_default"`
		FrontFemale      interface{} `json:"front_female"`
		FrontShiny       string      `json:"front_shiny"`
		FrontShinyFemale interface{} `json:"front_shiny_female"`
		Other            struct {
			DreamWorld struct {
				FrontDefault interface{} `json:"front_default"`
				FrontFemale  interface{} `json:"front_female"`
			} `json:"dream_world"`
			Home struct {
				FrontDefault     string      `json:"front_default"`
				FrontFemale      interface{} `json:"front_female"`
				FrontShiny       string      `json:"front_shiny"`
				FrontShinyFemale interface{} `json:"front_shiny_female"`
			} `json:"home"`
			OfficialArtwork struct {
				FrontDefault string `json:"front_default"`
				FrontShiny   string `json:"front_shiny"`
			} `json:"official-artwork"`
			Showdown struct {
				BackDefault      interface{} `json:"back_default"`
				BackFemale       interface{} `json:"back_female"`
				BackShiny        interface{} `json:"back_shiny"`
				BackShinyFemale  interface{} `json:"back_shiny_female"`
				FrontDefault     interface{} `json:"front_default"`
				FrontFemale      interface{} `json:"front_female"`
				FrontShiny       interface{} `json:"front_shiny"`
				FrontShinyFemale interface{} `json:"front_shiny_female"`
			} `json:"showdown"`
		} `json:"other"`
		Versions struct {
			GenerationI struct {
				RedBlue struct {
					BackDefault      interface{} `json:"back_default"`
					BackGray         interface{} `json:"back_gray"`
					BackTransparent  interface{} `json:"back_transparent"`
					FrontDefault     interface{} `json:"front_default"`
					FrontGray        interface{} `json:"front_gray"`
					FrontTransparent interface{} `json:"front_transparent"`
				} `json:"red-blue"`
				Yellow struct {
					BackDefault      interface{} `json:"back_default"`
					BackGray         interface{} `json:"back_gray"`
					BackTransparent  interface{} `json:"back_transparent"`
					FrontDefault     interface{} `json:"front_default"`
					FrontGray        interface{} `json:"front_gray"`
					FrontTransparent interface{} `json:"front_transparent"`
				} `json:"yellow"`
			} `json:"generation-i"`
			GenerationIi struct {
				Crystal struct {
					BackDefault           interface{} `json:"back_default"`
					BackShiny             interface{} `json:"back_shiny"`
					BackShinyTransparent  interface{} `json:"back_shiny_transparent"`
					BackTransparent       interface{} `json:"back_transparent"`
					FrontDefault          interface{} `json:"front_default"`
					FrontShiny            interface{} `json:"front_shiny"`
					FrontShinyTransparent interface{} `json:"front_shiny_transparent"`
					FrontTransparent      interface{} `json:"front_transparent"`
				} `json:"crystal"`
				Gold struct {
					BackDefault      interface{} `json:"back_default"`
					BackShiny        interface{} `json:"back_shiny"`
					FrontDefault     interface{} `json:"front_default"`
					FrontShiny       interface{} `json:"front_shiny"`
					FrontTransparent interface{} `json:"front_transparent"`
				} `json:"gold"`
				Silver struct {
					BackDefault      interface{} `json:"back_default"`
					BackShiny        interface{} `json:"back_shiny"`
					FrontDefault     interface{} `json:"front_default"`
					FrontShiny       interface{} `json:"front_shiny"`
					FrontTransparent interface{} `json:"front_transparent"`
				} `json:"silver"`
			} `json:"generation-ii"`
			GenerationIii struct {
				Emerald struct {
					FrontDefault interface{} `json:"front_default"`
					FrontShiny   interface{} `json:"front_shiny"`
				} `json:"emerald"`
				FireredLeafgreen struct {
					BackDefault  interface{} `json:"back_default"`
					BackShiny    interface{} `json:"back_shiny"`
					FrontDefault interface{} `json:"front_default"`
					FrontShiny   interface{} `json:"front_shiny"`
				} `json:"firered-leafgreen"`
				RubySapphire struct {
					BackDefault  interface{} `json:"back_default"`
					BackShiny    interface{} `json:"back_shiny"`
					FrontDefault interface{} `json:"front_default"`
					FrontShiny   interface{} `json:"front_shiny"`
				} `json:"ruby-sapphire"`
			} `json:"generation-iii"`
			GenerationIv struct {
				DiamondPearl struct {
					BackDefault      interface{} `json:"back_default"`
					BackFemale       interface{} `json:"back_female"`
					BackShiny        interface{} `json:"back_shiny"`
					BackShinyFemale  interface{} `json:"back_shiny_female"`
					FrontDefault     interface{} `json:"front_default"`
					FrontFemale      interface{} `json:"front_female"`
					FrontShiny       interface{} `json:"front_shiny"`
					FrontShinyFemale interface{} `json:"front_shiny_female"`
				} `json:"diamond-pearl"`
				HeartgoldSoulsilver struct {
					BackDefault      interface{} `json:"back_default"`
					BackFemale       interface{} `json:"back_female"`
					BackShiny        interface{} `json:"back_shiny"`
					BackShinyFemale  interface{} `json:"back_shiny_female"`
					FrontDefault     interface{} `json:"front_default"`
					FrontFemale      interface{} `json:"front_female"`
					FrontShiny       interface{} `json:"front_shiny"`
					FrontShinyFemale interface{} `json:"front_shiny_female"`
				} `json:"heartgold-soulsilver"`
				Platinum struct {
					BackDefault      interface{} `json:"back_default"`
					BackFemale       interface{} `json:"back_female"`
					BackShiny        interface{} `json:"back_shiny"`
					BackShinyFemale  interface{} `json:"back_shiny_female"`
					FrontDefault     interface{} `json:"front_default"`
					FrontFemale      interface{} `json:"front_female"`
					FrontShiny       interface{} `json:"front_shiny"`
					FrontShinyFemale interface{} `json:"front_shiny_female"`
				} `json:"platinum"`
			} `json:"generation-iv"`
			GenerationV struct {
				BlackWhite struct {
					Animated struct {
						BackDefault      interface{} `json:"back_default"`
						BackFemale       interface{} `json:"back_female"`
						BackShiny        interface{} `json:"back_shiny"`
						BackShinyFemale  interface{} `json:"back_shiny_female"`
						FrontDefault     interface{} `json:"front_default"`
						FrontFemale      interface{} `json:"front_female"`
						FrontShiny       interface{} `json:"front_shiny"`
						FrontShinyFemale interface{} `json:"front_shiny_female"`
					} `json:"animated"`
					BackDefault      interface{} `json:"back_default"`
					BackFemale       interface{} `json:"back_female"`
					BackShiny        interface{} `json:"back_shiny"`
					BackShinyFemale  interface{} `json:"back_shiny_female"`
					FrontDefault     interface{} `json:"front_default"`
					FrontFemale      interface{} `json:"front_female"`
					FrontShiny       interface{} `json:"front_shiny"`
					FrontShinyFemale interface{} `json:"front_shiny_female"`
				} `json:"black-white"`
			} `json:"generation-v"`
			GenerationVi struct {
				OmegarubyAlphasapphire struct {
					FrontDefault     interface{} `json:"front_default"`
					FrontFemale      interface{} `json:"front_female"`
					FrontShiny       interface{} `json:"front_shiny"`
					FrontShinyFemale interface{} `json:"front_shiny_female"`
				} `json:"omegaruby-alphasapphire"`
				XY struct {
					FrontDefault     interface{} `json:"front_default"`
					FrontFemale      interface{} `json:"front_female"`
					FrontShiny       interface{} `json:"front_shiny"`
					FrontShinyFemale interface{} `json:"front_shiny_female"`
				} `json:"x-y"`
			} `json:"generation-vi"`
			GenerationVii struct {
				Icons struct {
					FrontDefault interface{} `json:"front_default"`
					FrontFemale  interface{} `json:"front_female"`
				} `json:"icons"`
				UltraSunUltraMoon struct {
					FrontDefault     interface{} `json:"front_default"`
					FrontFemale      interface{} `json:"front_female"`
					FrontShiny       interface{} `json:"front_shiny"`
					FrontShinyFemale interface{} `json:"front_shiny_female"`
				} `json:"ultra-sun-ultra-moon"`
			} `json:"generation-vii"`
			GenerationViii struct {
				Icons struct {
					FrontDefault interface{} `json:"front_default"`
					FrontFemale  interface{} `json:"front_female"`
				} `json:"icons"`
			} `json:"generation-viii"`
		} `json:"versions"`
	} `json:"sprites"`
	Stats []struct {
		BaseStat int `json:"base_stat"`
		Effort   int `json:"effort"`
		Stat     struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"stat"`
	} `json:"stats"`
	Types []struct {
		Slot int `json:"slot"`
		Type struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"type"`
	} `json:"types"`
	Weight int `json:"weight"`
}
type Pokemon struct {
	Id      int    `json:"id"`
	Name    string `json:"name"`
	Hp      int    `json:"hp"`
	Attack  int    `json:"attack"`
	Defense int    `json:"defense"`
	Speed   int    `json:"speed"`
}

func (f *FullPokemonData) ToSlim() Pokemon {
	return Pokemon{
		Id:      f.ID,
		Name:    f.Name,
		Hp:      f.Stats[0].BaseStat,
		Attack:  f.Stats[1].BaseStat,
		Defense: f.Stats[2].BaseStat,
		Speed:   f.Stats[5].BaseStat,
	}
}
