package main

import (
	"strings"
)

// SplitStringByComma This function expect a string value and return a split array
func SplitStringByComma(str string) []string {
	cleaned := strings.Replace(str, ",", " ", -1)
	return strings.Fields(cleaned)
}
