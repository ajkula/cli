package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

// constants
const (
	newsURL = "https://newsapi.org/v2/top-headlines?apiKey=f99aa135983b46be95358b8d9da1018e"
)

type News struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	Url         string `json:"url"`
	UrlToImage  string `json:"urlToImage"`
	PublishedAt string `json:"publishedAt"`
	Content     string `json:"content"`
	Source      Source `json:"source"`
}

type Source struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type Articles struct {
	Status       string
	TotalResults int
	Articles     []News
}

func getNews(name string) Articles {
	// send GET request to GitHub API with the requested user "name"
	resp, err := http.Get(newsURL + "&country=" + name)
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
	// fmt.Println(string(body))

	// create a user variable of type "User" struct to store the "Unmarshal"-ed (aka parsed JSON) data, then return the user
	var news Articles
	json.Unmarshal(body, &news)
	// json.Unmarshal(&news.Articles, &news.Articles)
	// fmt.Println(news.Articles)
	return news
}

func DisplayNews(name string) {
	fmt.Printf("Getting news: %s\n", news)
	results := getNews(news)
	for _, res := range results.Articles {
		fmt.Println("**********************************************************")
		fmt.Println(`Source:             `, res.Source.Name)
		fmt.Println(`Publishing date:    `, res.PublishedAt)
		fmt.Println(`Title:              `, res.Title)
		// fmt.Println(`Description:        `, res.Description)
		fmt.Println(`Content:            `, res.Content)
		fmt.Println(`Url:                `, res.Url)
		fmt.Println(`UrlToImage:         `, res.UrlToImage)
	}
}
