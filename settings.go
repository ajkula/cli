package main

import (
	"fmt"
	"os"

	"gopkg.in/yaml.v2"
)

type Settings struct {
	Country string
}

func noFileError(e error) {}

func ReadSettingsFile() {
	file, err := os.Open("./settings.yml")
	check(err)
	defer file.Close()

	var UserSettings Settings
	if file != nil {
		dec := yaml.NewDecoder(file)
		dec.Decode(&UserSettings)

		fmt.Println(file)
		fmt.Println(UserSettings)
	}
}
