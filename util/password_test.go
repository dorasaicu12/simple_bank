package util

import (
	"testing"

	"github.com/stretchr/testify/require"
	"golang.org/x/crypto/bcrypt"
)

func TestPassword(t *testing.T){
  password := RandomString(6)

  hashed,err := HashePassword(password)

  require.NoError(t,err)
  require.NotEmpty(t,hashed)

  err = CheckPassWord(password,hashed)

  require.NoError(t,err)

  wrong_password := RandomString(6)

  err = CheckPassWord(wrong_password,hashed)

  require.EqualError(t,err,bcrypt.ErrMismatchedHashAndPassword.Error())
}