package helper

import "strings"

func Isvalid(firstname string, lastname string, usertickets int, email string, remainingtickets uint) (bool, bool, bool) {
	isValidname := len(firstname) >= 2 && len(lastname) >= 2
	isValidemail := strings.Contains(email, "@")
	isValidticketNumber := usertickets > 0 && remainingtickets >= uint(usertickets)

	return isValidname, isValidemail, isValidticketNumber
}
