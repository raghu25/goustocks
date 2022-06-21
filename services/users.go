package services

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/raghu25/goustocks/login"
)

// func GetUser(string username) {

// 	fmt.Println(username)

// }
func RegisterUser(user *login.User) *login.User {

	url := "https://nodeapi.collabmantra.com/gousers"
	method := "POST"

	payload, err := json.MarshalIndent(user, "", " ")
	if err != nil {
		return nil
	}

	payloadToSend := strings.NewReader(string(payload))

	fmt.Println(string(payload))

	client := &http.Client{}
	req, err := http.NewRequest(method, url, payloadToSend)

	if err != nil {
		fmt.Println(err)
		return nil
	}
	req.Header.Add("Content-Type", "application/json")

	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return nil
	}

	fmt.Println(body)
	return user
}
