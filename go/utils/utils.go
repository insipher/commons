// Package utils provides utilities to handle common activities such as logging, error handling, etc.
//
package utils

func ArrayContains(arr []string, str string) bool {
	for _, a := range arr {
		if a == str {
			return true
		}
	}
	return false
}
