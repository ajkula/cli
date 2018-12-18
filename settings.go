package main

import (
	"fmt"
	"io/ioutil"
	"os"

	"gopkg.in/yaml.v2"
)

type Settings struct {
	User struct {
		Country  string
		UserName string
	}
}

func noFileError(err error, UserSettings *Settings) {
	if err != nil {
		UserSettings.set("country", "france,fr")
		fmt.Printf("Error opening settings file: %s\n", err)
		us, err := yaml.Marshal(&UserSettings)
		fmt.Println("Writing default settings file:", string(us))
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

func (s *Settings) set(key, value string) {
	// if key == "country" {
	// 	s.User.Country = value
	// }
	switch key {
	case "country":
		s.User.Country = value
		break
	case "user":
		s.User.UserName = value
		break
	}
	us, err := yaml.Marshal(s)
	check(err)
	ioutil.WriteFile("./settings.yml", us, 0644)
}
