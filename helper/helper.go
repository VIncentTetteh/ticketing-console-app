package helper

import "strings"

func InputValidator(firstName string,lastName string, email string) (bool,bool){
	isValidName := len(firstName) >= 2 && len(lastName) >= 2
	isEmailValid := strings.Contains(email,"@")
	return isEmailValid,isValidName
}
