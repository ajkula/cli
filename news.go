package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
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

func getNews(name, category string) Articles {
	var cat string = ""
	// send GET request to GitHub API with the requested user "name"
	if category != "" {
		cat = "&category=" + category
	}
	resp, err := http.Get(newsURL + "&country=" + name + cat)
	// if err occurs during GET request, then throw error and quit application
	check(err)

	// Always good practice to defer closing the response body.
	// If application crashes or function finishes successfully, GO will always execute this "defer" statement
	defer resp.Body.Close()

	// read the response body and handle any errors during reading.
	body, err := ioutil.ReadAll(resp.Body)
	check(err)

	// create a user variable of type "User" struct to store the "Unmarshal"-ed (aka parsed JSON) data, then return the user
	var news Articles
	json.Unmarshal(body, &news)
	return news
}

func DisplayNews(news, category, x string) {
	fmt.Printf("Getting %s news: %s\n", category, news)
	results := getNews(news, category)

	var size int
	var err error
	if x != "" {
		if size, err = strconv.Atoi(x); err != nil {
			check(err)
		}
	} else {
		size = 80
	}

	for _, res := range results.Articles {
		res := res
		ch := make(chan string)
		go func() {
			concat := ""
			concat += makeLines("**********************************************************")
			concat += makeLines(`Source:             `, res.Source.Name)
			concat += makeLines(`Publishing date:    `, res.PublishedAt)
			concat += makeLines(`Title:              `, res.Title)
			// fmt.Println(`Description:        `, res.Description)
			concat += makeLines(`Content:            `, res.Content)
			concat += makeLines(`Url:                `, res.Url)
			// fmt.Println(`UrlToImage:         `, res.UrlToImage)
			concat += makeLines("")
			if res.UrlToImage != "" {
				// asciiArt := Convert2Ascii(res.UrlToImage, size)
				concat += string(Convert2Ascii(res.UrlToImage, size))
			}
			ch <- concat
		}()

		select {
		case str := <-ch:
			fmt.Println(str)
		}
	}
}

func makeLines(str ...interface{}) string {
	s := fmt.Sprint(str...)
	return s + "\n"
}
