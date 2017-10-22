package main

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
)

type TrelloMember struct {
	Id       string `json:"id"`
	FullName string `json:"fullName"`
	UserName string `json:"username"`
}

type TrelloMemberResponse []TrelloMember

func perror(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func GetTrelloMembers(w http.ResponseWriter, r *http.Request) {
	const (
		trelloURL   = "https://api.trello.com/1/boards/wF8wnpha/members"
		trelloKey   = "d5eb90af629773d627ad6449bd733318"
		trelloToken = "580440ae3857fd9f0cd0dbcf950d6cf1beeccbd6ea64c31e064efb9219e67584"
	)

	req, err := http.NewRequest("GET", trelloURL, nil)
	perror(err)

	q := req.URL.Query()
	q.Add("key", trelloKey)
	q.Add("token", trelloToken)
	req.URL.RawQuery = q.Encode()

	req.Close = true
	fmt.Println(req)
	resp, err := http.DefaultClient.Do(req)
	perror(err)

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	perror(err)

	var data TrelloMemberResponse
	err = json.Unmarshal(body, &data)
	trelloMembers, err := json.Marshal(&data)
	perror(err)

	io.WriteString(w, string(trelloMembers))
}

var mux map[string]func(http.ResponseWriter, *http.Request)

func main() {
	server := http.Server{
		Addr:    ":8000",
		Handler: &myHandler{},
	}
	mux = make(map[string]func(http.ResponseWriter, *http.Request))
	mux["/"] = GetTrelloMembers

	log.Print("Server started: localhost:8000")
	server.ListenAndServe()
}

type myHandler struct{}

func (*myHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if h, ok := mux[r.URL.String()]; ok {
		h(w, r)
		return
	}
	io.WriteString(w, "My server: "+r.URL.String())
}
