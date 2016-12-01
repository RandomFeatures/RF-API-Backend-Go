package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"

	"github.com/julienschmidt/httprouter"
)

type newsjson struct {
	ID        int    `json:"id"`
	Title     string `json:"title"`
	Date      string `json:"date"`
	Localurl  string `json:"localurl"`
	Image     string `json:"image"`
	Text      string `json:"text"`
	Remoteurl string `json:"remoteurl"`
}

var newsData []byte

func init() {
	var e error
	newsData, e = ioutil.ReadFile("./data/news.json")
	if e != nil {
		LogThis(e.Error())
		os.Exit(1)
	}
}

//GetNewsRoutes assign all of the news routes to the Router
func GetNewsRoutes(router *httprouter.Router) {

	router.GET("/api/news", NewsList)
	router.GET("/api/news/:id", NewsByID)
}

//NewsList returns the list of news items to the server
func NewsList(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	fmt.Fprintf(w, "%s\n", string(newsData))
}

//NewsByID returns a single news item to the server
func NewsByID(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	index, _ := strconv.Atoi(ps.ByName("id"))

	var newslist []newsjson
	json.Unmarshal(newsData, &newslist)

	for _, item := range newslist {
		if item.ID == index {
			newsItem, err := json.Marshal(item)
			if err != nil {
				fmt.Println(err)
				return
			}
			fmt.Fprintf(w, "%s\n", string(newsItem))
			return
		}
	}

}
