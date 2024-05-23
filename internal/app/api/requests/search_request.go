package requests

type SearchRequest struct {
	Name    string `json:"name"`
	Hp      int    `json:"hp"`
	Attack  int    `json:"attack"`
	Defense int    `json:"defense"`
	Speed   int    `json:"speed"`
}
