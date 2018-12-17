package main

import (
	"fmt"
	"io/ioutil"
	"os"

	"gopkg.in/yaml.v2"
)

type Settings struct {
	Country string
}

func noFileError(err error, UserSettings *Settings) {
	UserSettings.set("france,fr")
	if err != nil {
		fmt.Printf("Error retrieving data: %s\n", err)
		us, err := yaml.Marshal(&UserSettings)
		check(err)
		ioutil.WriteFile("./settings.yml", us, 0644)
	}
}

func ReadSettingsFile() {
	var UserSettings Settings
	file, err := os.Open("./settings.yml")
	noFileError(err, &UserSettings)
	defer file.Close()

	if file != nil {
		dec := yaml.NewDecoder(file)
		dec.Decode(&UserSettings)

		fmt.Println("UserSettings", UserSettings)
	}
}

func (s Settings) set(value string) {
	s.Country = value
	us, err := yaml.Marshal(s)
	check(err)
	ioutil.WriteFile("./settings.yml", us, 0644)
}
