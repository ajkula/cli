package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

// constants
const (
	imdbApiURL = "http://www.omdbapi.com/?apikey=9c7cec43"
)

// User struct represents the JSON data from GitHub API: https://api.github.com/users/defunct
// This struct was generated via a JSON-to-GO utility by Matt Holt: https://mholt.github.io/json-to-go/
type Movie struct {
	Title      string `json:"Title"`
	Year       string `json:"Year"`
	Type       string `json:"type"`
	Rated      string `json:"Rated"`
	Released   string `json:"Released"`
	Runtime    string `json:"Runtime"`
	Genre      string `json:"Genre"`
	Director   string `json:"Director"`
	Writer     string `json:"Writer"`
	Actors     string `json:"Actors"`
	Plot       string `json:"Plot"`
	Language   string `json:"Language"`
	Country    string `json:"Country"`
	Awards     string `json:"Awards"`
	Poster     string `json:"Poster"`
	ImdbRating string `json:"imdbRating"`
	ImdbVotes  string `json:"imdbVotes"`
	DVD        string `json:"DVD"`
	ID         string `json:"imdbID"`
}

// getUsers queries GitHub API for a given user
func getMovie(name string) Movie {
	// send GET request to GitHub API with the requested user "name"
	resp, err := http.Get(imdbApiURL + "&t=" + name)
	// if err occurs during GET request, then throw error and quit application
	if err != nil {
		log.Fatalf("Error retrieving data: %s\n", err)
	}

	// Always good practice to defer closing the response body.
	// If application crashes or function finishes successfully, GO will always execute this "defer" statement
	defer resp.Body.Close()

	// read the response body and handle any errors during reading.
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalf("Error reading data: %s\n", err)
	}

	// create a user variable of type "User" struct to store the "Unmarshal"-ed (aka parsed JSON) data, then return the user
	var movie Movie
	json.Unmarshal(body, &movie)
	return movie
}

func cleanQuotes(s string) string {
	var result string
	for _, c := range s {
		if string(c) == " " {
			result += "%20"
		} else if string(c) != "\"" {
			result += string(c)
		}
	}
	return result
}
