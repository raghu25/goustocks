package login

import (
	"fmt"
	"time"

	"github.com/raghu25/goustocks/cli"
)

type User struct {
	Username      string `json:"_id"`
	Password      string
	IsLoggedIn    bool
	LastLoggedIOn string
	IsNew         bool
}

type LoginStruct struct {
	Cli cli.IInputOutput
}

// Private function
func (l *LoginStruct) login() *User {
	username := l.Cli.ReadString("Enter user username:")
	password := l.Cli.ReadString("Enter user password:")

	return &User{
		Username:      username,
		Password:      password,
		IsLoggedIn:    false,
		LastLoggedIOn: "",
		IsNew:         false,
	}
}

func (l *LoginStruct) Login2(username string, password string) *User {
	return &User{
		Username:      username,
		Password:      password,
		IsLoggedIn:    false,
		LastLoggedIOn: "",
		IsNew:         false,
	}
}

func (l *LoginStruct) register() *User {
	username := l.Cli.ReadString("Enter user username:")
	password := l.Cli.ReadString("Enter user password:")

	return &User{
		Username:      username,
		Password:      password,
		IsLoggedIn:    true,
		LastLoggedIOn: time.Now().String(),
		IsNew:         true,
	}

}

// Public function
func (l *LoginStruct) Login() *User {
	option := l.Menu()

	switch option {
	case 1:
		return l.login()
	case 2:
		return l.register()
	}

	return nil
}

func (l *LoginStruct) Menu() int {

	fmt.Println("Press 1 to Login")
	fmt.Println("Press 2 to Register")
	return l.Cli.ReadInt("")
}
