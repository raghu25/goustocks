package math

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAdd(t *testing.T) {

	got := Add(4, 6)
	want := 10

	if got != want {
		t.Errorf("got %q, wanted %q", got, want)
	}
}

func TestAddWithAsset(t *testing.T) {
	got := Add(4, 6)
	assert.Equal(t, 10, got)
}

func TestSubstractWithAsset(t *testing.T) {
	got := Subtract(4, 6)
	assert.Equal(t, -2, got)
}

func Test_MultiPlay(t *testing.T) {
	got := Multiplay(4, 6)
	assert.Equal(t, 24, got)
}
