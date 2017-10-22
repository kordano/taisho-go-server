package models

type TrelloCredentials struct {
	Key     string `json:"key"`
	Token   string `json:"token"`
	BaseURL string `json:"baseURL"`
}

type Configuration struct {
	Trello TrelloCredentials `json:"trello"`
}
