package main

import (
	"fmt"
	"sort"
)

func longestCommonPrefix(words []string) string {
	if len(words) == 1 {
		return words[0]
	}

	stringSlice := ""
	sort.Strings(words)
	firstWord := words[0]
	lastWord := words[len(words)-1]

	for i, char := range firstWord {
		if byte(char) != lastWord[i] {
			break
		} else {
			stringSlice += string(char)
		}
	}

	return stringSlice
}

func main() {
	fmt.Println(longestCommonPrefix([]string{"flower", "flowers"}))
}
