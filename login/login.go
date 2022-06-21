package login

import (
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

//Public function
func Login() *User {
	return login()
}
