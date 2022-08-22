package login

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_That_UserName_works(t *testing.T) {
	t.Log("TestConvert")
	user := Login2("raghu", "password")
	assert.Equal(t, "raghu", user.Username)
}
