package helpers

import "strings"

func CheckUserPass(username, password string) bool {
	userpass := make(map[string]string)
	userpass["hello"] = "itsme"
	userpass["john"] = "doe"

	if val, ok := userpass[username]; ok {
		return val == password
	} else {
		return false
	}
}

func EmptyUserPass(username, password string) bool {
	return strings.Trim(username, " ") == "" ||
		strings.Trim(password, " ") == ""
}
