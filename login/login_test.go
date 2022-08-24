package login

import (
	"fmt"
	"testing"

	"github.com/raghu25/goustocks/cli"
	"github.com/stretchr/testify/assert"
)

type mockCLI struct {
	cli.IInputOutput
	t   *testing.T
	err error
}

var returnValue string

func (c *mockCLI) ReadString(msg string) string {
	return returnValue

}

func (c *mockCLI) Clear() {
	fmt.Println("")

}

func newMockCLI(t *testing.T) cli.IInputOutput {
	return &mockCLI{}
}

func newLogin(t *testing.T) *LoginStruct {
	return &LoginStruct{
		Cli: newMockCLI(t),
	}

}

func Test_That_UserName_works(t *testing.T) {
	t.Log("TestConvert")
	login := newLogin(t)
	user := login.Login2("raghu", "password")
	assert.Equal(t, "raghu", user.Username)
}

func Test_That_UserName_works_org(t *testing.T) {
	t.Log("TestConvert")
	login := newLogin(t)
	returnValue = "Hello"
	user := login.login()
	assert.Equal(t, "Hello", user.Username)
	assert.Equal(t, "Hello", user.Password)
}

func Test_That_Register(t *testing.T) {
	t.Log("TestConvert")
	login := newLogin(t)
	returnValue = "Darshan"
	user := login.register()
	assert.Equal(t, "Darshan", user.Username)
	assert.Equal(t, "Darshan", user.Password)
}
