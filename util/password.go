package util

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

//return hashed password
func HashePassword(password string) (string,error) {
  hashed,err := bcrypt.GenerateFromPassword([]byte(password),bcrypt.DefaultCost)
  if err != nil {
	return "",fmt.Errorf("failed to hashed pass: %v",err)
  }
  return string(hashed),nil
}

func CheckPassWord(password string,hashedPassword string) error {
   return bcrypt.CompareHashAndPassword([]byte(hashedPassword),[]byte(password))
}