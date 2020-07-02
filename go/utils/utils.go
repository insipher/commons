// Package utils provides utilities to handle common activities such as logging, error handling, etc.
//
package utils

import "encoding/json"

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

func IsEditor(userID string, editors json.RawMessage) (bool, error) {
	hasAccess := false
	var editorIDs []string
	err := json.Unmarshal(editors, &editorIDs)
	if err != nil {
		return false, err
	}

	// check to see if user already is in favorties list, if so just return
	if IndexOf(userID, editorIDs) > 0 {
		return true, nil
	}
	return hasAccess, nil
}
