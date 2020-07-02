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

func IndexOf(element string, arr []string) int {
	for k, v := range arr {
		if element == v {
			return k
		}
	}
	return -1 //not found.
}

func RemoveIndex(arr []string, index int) []string {
	return append(arr[:index], arr[index+1:]...)
}
