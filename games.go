package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"

	"github.com/julienschmidt/httprouter"
)

type gamesjson struct {
	ID        int    `json:"id"`
	Game      string `json:"game"`
	Type      string `json:"type"`
	System    string `json:"system"`
	Version   string `json:"version"`
	Downloads struct {
		Market  string `json:"market"`
		Slideme string `json:"slideme"`
		Getjar  string `json:"getjar"`
		Amazon  string `json:"amazon"`
	} `json:"downloads"`
	Localurl    string `json:"localurl"`
	Iconimage   string `json:"iconimage"`
	Featurelogo string `json:"featurelogo"`
	Htmldesc    string `json:"htmldesc"`
	Appbrain    []struct {
		ID    string `json:"id"`
		Title string `json:"title"`
		URL   string `json:"url"`
	} `json:"appbrain"`
	Googleplay []struct {
		URL    string `json:"url"`
		Qrcode string `json:"qrcode"`
	} `json:"googleplay"`
	Slideme []struct {
		URL    string `json:"url"`
		Qrcode string `json:"qrcode"`
	} `json:"slideme"`
	Getjar []struct {
		URL    string `json:"url"`
		Qrcode string `json:"qrcode"`
	} `json:"getjar"`
	Amazon      []interface{} `json:"amazon"`
	Video       string        `json:"video"`
	Screenshots struct {
		Base      string `json:"base"`
		Thumbbase string `json:"thumbbase"`
		Images    []struct {
			Image string `json:"image"`
			Thumb string `json:"thumb"`
		} `json:"images"`
	} `json:"screenshots"`
}

var gameData []byte

func init() {
	var e error
	gameData, e = ioutil.ReadFile("./data/games.json")
	if e != nil {
		LogThis(e.Error())
		os.Exit(1)
	}
}

//GetGameRoutes Assign all of the Game routes to the Router
func GetGameRoutes(r *httprouter.Router) {

	r.GET("/api/games", GameList)
	r.GET("/api/games/:name", GameByName)
}

//GameList returns the list of Games to the server
func GameList(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {

	fmt.Fprintf(w, "%s\n", string(gameData))

}

//GameByName returns a single Game by Name to the server
func GameByName(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	gameName := ps.ByName("name")

	var gamelist []gamesjson
	json.Unmarshal(gameData, &gamelist)

	for _, game := range gamelist {
		if game.Localurl == gameName {
			gameItem, err := json.Marshal(game)
			if err != nil {
				fmt.Println(err)
				return
			}
			fmt.Fprintf(w, "%s\n", string(gameItem))
			return
		}
	}

}
