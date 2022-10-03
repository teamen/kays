package auth

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEncryptPassword(t *testing.T) {
	password := "foobared"
	hashedPassword, err := Encrypt(password)
	assert.Nil(t, err)
	assert.Nil(t, Compare(hashedPassword, password))
}
