package main

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/kordano/taisho-go-server/models"
)

var configuration models.Configuration

func perror(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func setConfig(configuration *models.Configuration) {
	configFile, err := os.Open("config.json")
	perror(err)

	jsonParser := json.NewDecoder(configFile)
	if err = jsonParser.Decode(configuration); err != nil {
		fmt.Println("parsing error!")
	}
	return
}

func fetchURL(endpoint string) []byte {
	trelloURL := configuration.Trello.BaseURL + endpoint

	req, err := http.NewRequest("GET", trelloURL, nil)
	perror(err)

	q := req.URL.Query()
	q.Add("key", configuration.Trello.Key)
	q.Add("token", configuration.Trello.Token)
	req.URL.RawQuery = q.Encode()

	req.Close = true
	resp, err := http.DefaultClient.Do(req)
	perror(err)

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	perror(err)

	return body
}

// GetTrelloMembers fetches Trello Members of Development Board
func GetTrelloMembers(w http.ResponseWriter, r *http.Request) {
	body := fetchURL("/boards/wF8wnpha/members")

	var data models.TrelloMemberResponse
	err := json.Unmarshal(body, &data)
	perror(err)

	trelloMembers, err := json.Marshal(&data)
	perror(err)

	io.WriteString(w, string(trelloMembers))
}

// GetTrelloBoardLists retrieves all Trello Lists on Development Board
func GetTrelloBoardLists(w http.ResponseWriter, r *http.Request) {
	body := fetchURL("/boards/wF8wnpha/lists")

	var data models.TrelloListsResponse
	err := json.Unmarshal(body, &data)
	perror(err)

	trelloBoard, err := json.Marshal(&data)
	perror(err)

	io.WriteString(w, string(trelloBoard))
}

var mux map[string]func(http.ResponseWriter, *http.Request)

func main() {
	setConfig(&configuration)

	server := http.Server{
		Addr:    ":8000",
		Handler: &myHandler{},
	}
	mux = make(map[string]func(http.ResponseWriter, *http.Request))
	mux["/"] = GetTrelloMembers
	mux["/board"] = GetTrelloBoardLists

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
