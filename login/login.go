package login

import (
	"fmt"
	"time"

	"github.com/raghu25/goustocks/cli"
)

type User struct {
	Username      string
	Password      string
	IsLoggedIn    bool
	LastLoggedIOn string
	IsNew         bool
}

//Private function
func login() *User {
	username := cli.ReadString("Enter user username:")
	password := cli.ReadString("Enter user password:")
	if username == "admin" && password == "123" {
		return &User{
			Username:      username,
			Password:      password,
			IsLoggedIn:    true,
			LastLoggedIOn: time.Now().String(),
			IsNew:         false,
		}
	} else {
		return &User{
			Username:      username,
			Password:      password,
			IsLoggedIn:    false,
			LastLoggedIOn: "",
			IsNew:         false,
		}
	}
}

func register() *User {
	username := cli.ReadString("Enter user username:")
	password := cli.ReadString("Enter user password:")

	return &User{
		Username:      username,
		Password:      password,
		IsLoggedIn:    true,
		LastLoggedIOn: time.Now().String(),
		IsNew:         true,
	}

}

//Public function
func Login() *User {
	option := Menu()

	switch option {
	case 1:
		return login()
	case 2:
		return register()
	}

	return nil
}

func Menu() int {

	fmt.Println("Press 1 to Login")
	fmt.Println("Press 2 to Register")
	return cli.ReadInt("")
}
