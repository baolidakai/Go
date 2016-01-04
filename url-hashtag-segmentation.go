package main

import (
	"fmt"
	"strings"
	"io/ioutil"
)

func parse(input string) string {
	// Parse the raw input
	if rune(input[0]) == '#' {
		return input[1:]
	} else {
		return strings.Split(input, ".")[1]
	}
}

func tokenize(input string, dictionary map[string]bool) ([]string, bool) {
	/*
	* We use recursion to check if the string could be tokenize
	* Returns the tokens and a boolean indicating whether input is tokenizable
	*/
	// Base case
	if len(input) == 0 {
		return []string{}, true
	}
	// Recursive step
	// Iterate through all possible prefix
	found := false // whether a valid partition is found
	var bestSoFar []string
	for idx := 0; idx < len(input); idx++ {
		prefix := input[:idx + 1]
		if dictionary[prefix] {
			// Check if this constitutes a valid partition
			partition, valid := tokenize(input[idx + 1:], dictionary)
			if valid {
				found = true
				bestSoFar = append([]string{prefix}, partition...)
			}
		}
	}
	if found {
		return bestSoFar, true
	} else {
		return []string{input}, false
	}
}

func main() {
	content, _ := ioutil.ReadFile("words.txt")
	words := strings.Split(string(content), "\n")
	words = words[:len(words) - 1]
	dictionary := make(map[string]bool)
	for _, word := range words {
		dictionary[word] = true
	}
	var N int
	fmt.Scan(&N)
	for i := 0; i < N; i++ {
		var input string
		fmt.Scanln(&input)
		// Parse input
		input = parse(input)
		// Tokenize input
		tokens, _ := tokenize(input, dictionary)
		for i:= 0; i < len(tokens); i++ {
			fmt.Print(tokens[i])
			if i != len(tokens) - 1 {
				fmt.Print(" ")
			} else {
				fmt.Println()
			}
		}
	}
}
