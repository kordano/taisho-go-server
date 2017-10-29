package models

type TrelloShortListEntry struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type TrelloBoard struct {
	ID           string                 `json:"id"`
	LastActivity string                 `json:"dateLastActivity"`
	Lists        []TrelloShortListEntry `json:"lists"`
}

type TrelloBoardList struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type TrelloListsResponse []TrelloBoardList

type TrelloCard struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	IsClosed    bool   `json:"closed"`
	DueComplete bool   `json:"dueComplete"`
	DueDate     string `json:"due"`
}

type TrelloBadges struct {
	CheckItemsCount        int `json:"checkItems"`
	CheckitemsCheckedCount int `json:"checkItemsChecked"`
}
