package login

import (
	"fmt"
	"time"

	"github.com/raghu25/goustocks/cli"
)

type User struct {
	username      string
	password      string
	isLoggedIn    bool
	lastLoggedIOn string
}

//Private function
func login() *User {
	username := cli.ReadString("Enter user username:")
	password := cli.ReadString("Enter user password:")
	if username == "admin" && password == "123" {
		return &User{
			username:      username,
			password:      password,
			isLoggedIn:    true,
			lastLoggedIOn: time.Now().String(),
		}
	} else {
		return &User{
			username:      username,
			password:      password,
			isLoggedIn:    false,
			lastLoggedIOn: "",
		}
	}
}

func register() *User {
	username := cli.ReadString("Enter user username:")
	password := cli.ReadString("Enter user password:")

	return &User{
		username:      username,
		password:      password,
		isLoggedIn:    true,
		lastLoggedIOn: time.Now().String(),
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
