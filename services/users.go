package services

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"

	"github.com/raghu25/goustocks/login"
)

func GetUser(username string) *login.User {
	url := "https://nodeapi.collabmantra.com/gousers/" + username
	method := "GET"
	payload := strings.NewReader(``)

	client := &http.Client{}
	req, err := http.NewRequest(method, url, payload)

	if err != nil {
		fmt.Println(err)
		return nil
	}
	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	if res.StatusCode == 404 {
		return nil
	}

	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return nil
	}

	user := &login.User{}
	errparse := json.Unmarshal(body, user)
	if errparse != nil {
		log.Fatalf("Error occured during unmarshaling. GetUsers Error: %s", errparse.Error())
	}

	return user
}

func LoginUser(user *login.User) *login.User {
	nuser := GetUser(user.Username)
	if nuser == nil {
		return nil
	}
	if nuser.Password == user.Password {
		return nuser
	}
	return nil

}

//Create ....
func RegisterUser(user *login.User) *login.User {

	checkUser := GetUser(user.Username)

	if checkUser != nil {
		fmt.Println("User already exist")
		return nil
	}

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

	errparse := json.Unmarshal(body, user)
	if errparse != nil {
		log.Fatalf("Error occured during unmarshaling. CreateUser Error: %s", errparse.Error())
	}

	return user
}
