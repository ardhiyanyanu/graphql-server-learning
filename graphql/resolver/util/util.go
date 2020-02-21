package util

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

var BaseURL = "https://dbs-sbackend.sumpahpalapa.com"

type login struct {
	Code           int    `json:"code"`
	Message        string `json:"message"`
	Name           string `json:"name"`
	ProfilePicture string `json:"profilePicture"`
	Token          string `json:"token"`
}

func GetToken() string {
	var err error
	var client = &http.Client{}
	var login login

	requestBody, err := json.Marshal(map[string]interface{}{
		"email":    "yanu@alterra.id",
		"password": "P@ssw0rd",
	})

	request, err := http.NewRequest("POST", BaseURL+"/login", bytes.NewBuffer(requestBody))
	if err != nil {
		fmt.Println("ERROR: " + fmt.Sprintf("%v", err))
		return ""
	}

	response, err := client.Do(request)
	if err != nil {
		fmt.Println("ERROR: " + fmt.Sprintf("%v", err))
		return ""
	}
	defer response.Body.Close()

	err = json.NewDecoder(response.Body).Decode(&login)
	if err != nil {
		fmt.Println("ERROR: " + fmt.Sprintf("%v", err))
		return ""
	}

	return login.Token
}
