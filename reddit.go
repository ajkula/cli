// "Package main" is the namespace declaration
package main

// importing standard libraries
import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

// constants
const (
	redditApiURL     = "https://www.reddit.com"
	postsEndPoint    = "/r/"
	commentsEndPoint = "/comments/"
	limit            = "/.json?limit=10"
	UserAgent        = "script:reddit.reader:v0.14 (by /u/Ptk7l2)"
)

// User struct represents the JSON data from GitHub API: https://api.github.com/users/defunct
// This struct was generated via a JSON-to-GO utility by Matt Holt: https://mholt.github.io/json-to-go/
type Post struct {
	Data struct {
		Selftext   string `json:"Selftext"`
		ID         string `json:"id"`
		CreatedUTC int64  `json:"Created_utc"`
		Author     string `json:"Author"`
	} `json:"data"`
}

type Comment struct {
	Data struct {
		Selftext   string `json:"Selftext"`
		ID         string `json:"id"`
		CreatedUTC int64  `json:"Created_utc"`
		Author     string `json:"Author"`
		Body       string `json:"Body"`
	} `json:"data"`
}

type Posts struct {
	Data struct {
		Children []Post `json:"children"`
	} `json:"data"`
}

type Comments struct {
	Data struct {
		Children []Comment `json:"children"`
	} `json:"data"`
}

// getUsers queries GitHub API for a given user
func getRedditPosts(name string) Posts {
	fmt.Println(redditApiURL + postsEndPoint + name + limit)
	url := redditApiURL + postsEndPoint + name + limit
	// send GET request to GitHub API with the requested user "name"
	req, err := http.NewRequest("GET", url, nil)
	// if err occurs during GET request, then throw error and quit application
	if err != nil {
		log.Fatalf("Error retrieving data: %s\n", err)
	}
	req.Header.Set("User-Agent", UserAgent)

	// Handle the request
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Fatalf("Error reading data: %s\n", err)
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
	var POSTS Posts
	json.Unmarshal(body, &POSTS)
	return POSTS
}

// getUsers queries GitHub API for a given user
func getRedditComments(name string) []Comments {
	// send GET request to GitHub API with the requested user "name"
	url := redditApiURL + commentsEndPoint + name + limit
	req, err := http.NewRequest("GET", url, nil)
	// if err occurs during GET request, then throw error and quit application
	if err != nil {
		log.Fatalf("Error retrieving data: %s\n", err)
	}
	req.Header.Set("User-Agent", UserAgent)

	// Handle the request
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Fatalf("Error reading data: %s\n", err)
	}

	// Always good practice to defer closing the response body.
	// If application crashes or function finishes successfully, GO will always execute this "defer" statement
	defer resp.Body.Close()

	// read the response body and handle any errors during reading.
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalf("Error reading data: %s\n", err)
	}
	// fmt.Println(string(body))

	var coms []Comments
	// create a user variable of type "User" struct to store the "Unmarshal"-ed (aka parsed JSON) data, then return the user
	json.Unmarshal(body, &coms)
	return coms
}
