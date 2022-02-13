package models

type Swapi struct {
	Count   int       `json:"count"`
	Next    string    `json:"next"`
	Results []Results `json:"results"`
}

type Results struct {
	Name  string   `json:"name"`
	Films []string `json:"films"`
}
