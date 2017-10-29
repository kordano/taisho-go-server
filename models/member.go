package models

type TrelloMember struct {
	ID       string `json:"id"`
	FullName string `json:"fullName"`
	UserName string `json:"username"`
}

type TrelloMemberResponse []TrelloMember
