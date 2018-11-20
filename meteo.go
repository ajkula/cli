package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

const (
	url              = "http://api.openweathermap.org/data/2.5/weather?q="
	appID            = "&appid=7431d386218c6bc0943c880b3c81b868"
	countryByDefault = "Paris,fr"
)

type coord struct {
	Lon float64
	Lat float64
}

type weather struct {
	ID          int
	Main        string
	Description string
	Icon        string
}

type corp struct {
	Temp     float32
	Pressure int
	Humidity int
	Temp_min float32
	Temp_max float32
}

type wind struct {
	Speed float32
	Deg   int
}
type clouds struct {
	All int
}

type sys struct {
	ID      int
	Message float32
	Country string
	Sunrise int64
	Sunset  int64
}

type MeteoCityNow struct {
	Coord      coord
	Weather    []weather
	Base       string
	Main       corp
	Visibility int
	Wind       wind
	Clouds     clouds
	Dt         int64
	Sys        sys
	ID         int
	Name       string
	Cod        int
}

func getMeteoByCity(name string) MeteoCityNow {
	fmt.Println(url + name + appID)
	url := url + name + appID
	// send GET request to GitHub API with the requested user "name"
	resp, err := http.Get(url)
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
	var meteo MeteoCityNow
	json.Unmarshal(body, &meteo)
	// json.Unmarshal(&news.Articles, &news.Articles)
	// fmt.Println(news.Articles)
	return meteo
}
