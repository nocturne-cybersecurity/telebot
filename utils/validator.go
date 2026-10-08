package utils

import (
	"regexp"
)

func IsValidTelegramUsername(username string) bool {
	if username == "" {
		return false
	}
	
	matched, _ := regexp.MatchString(`^[a-zA-Z0-9_]{5,32}$`, username)
	return matched
}

func IsValidCommand(command string) bool {
	if command == "" {
		return false
	}
	
	matched, _ := regexp.MatchString(`^/[a-zA-Z0-9_]+$`, command)
	return matched
}

func IsValidEmail(email string) bool {
	if email == "" {
		return false
	}
	
	matched, _ := regexp.MatchString(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`, email)
	return matched
}
