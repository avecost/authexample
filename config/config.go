package config

import (
	"os"
	"fmt"
	"encoding/json"
)

type AppUser struct {
	User 		string		`json:"user"`
	Password 	string		`json:"password"`
	Url 		[]string 	`json:"url"`
}

func LoadConfiguration(file string) []AppUser {
	var users []AppUser
	configFile, err := os.Open(file)
	defer configFile.Close()

	if err != nil {
		fmt.Println(err.Error())
	}
	jsonParser := json.NewDecoder(configFile)
	jsonParser.Decode(&users)
	return users
}