package main

import (
	"fmt"
	"sort"
	"strings"
)

func longestCommonPrefix(words []string) string {
	stringSlice := ""
	sort.Strings(words)
	firstWord := words[0]
	lastWord := words[len(words)-1]

	for i, char := range strings.Split(firstWord, "") {
		if firstWord[i] != lastWord[i] {
			break
		} else {
			stringSlice += char
		}
	}

	return stringSlice
}

func main() {
	fmt.Println(longestCommonPrefix([]string{"flower", "flowers"}))
}
