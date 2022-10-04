package token

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSignToken(t *testing.T) {

	signingSecret := "zNQK5AOOf0uyAp35OnWTMEOpbRohBxOng7AcSfroxs4y7wOP15gC82WjwcnojhTD"
	Init(signingSecret)
	id := 1
	username := "Wayne Tse"

	tokenString, err := Sign(id, username)

	if err != nil {
		t.Fatal(err.Error())
	}
	t.Log(tokenString)
	assert.Nil(t, err)
	assert.NotEmpty(t, tokenString)

	assert.Equal(t, signingSecret, GetSignSecret())

}

func TestParseToken(t *testing.T) {
	id := 1
	username := "Wayne Tse"
	tokenString, _ := Sign(id, username)
	id, username, err := Parse(tokenString)
	if err != nil {
		t.Fatalf("failed to parse token:%s", err.Error())
	}

	assert.Equal(t, 1, id)
	assert.Equal(t, "Wayne Tse", username)
}
