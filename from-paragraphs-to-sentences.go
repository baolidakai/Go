package main
/*
 * sentence tokenize a paragraph
 * Usage: sent_tokenize(input)
 * input - a string
 */

import (
	"fmt"
	"bufio"
	"os"
)

func sent_tokenize(input string) []string {
	// prev stores the previous delimiter
	// curr stores the current delimiter
	var rtn []string
	n := len(input)
	prev := 0
	curr := 0
	for curr < n {
		currChar := rune(input[curr])
		isDelimiter := false
		if curr == n - 1 {
			isDelimiter = true
		} else if rune(input[curr + 1]) != '"' {
			if currChar == '?' || currChar == '!' {
				isDelimiter = true
			} else if currChar == '.' {
				runner := curr - 1
				for runner > 0 && rune(input[runner - 1]) != ' ' {
					runner--
				}
				runnerChar := rune(input[runner])
				if runnerChar < 'A' || runnerChar > 'Z' {
					isDelimiter = true
				}
			}
		}
		if isDelimiter {
			// Print the current sentence
			fmt.Println(input[prev:curr + 1])
			rtn = append(rtn, input[prev:curr + 1])
			// Update prev
			prev = curr + 1
			for prev < n && rune(input[prev]) == ' ' {
				prev++
			}
		}
		curr++
	}
	return rtn
}


func main() {
	// Split by ?_, !_, or ._, and of course EOL
	// Read in the sentence
	sc := bufio.NewScanner(os.Stdin)
	sc.Scan()
	input := sc.Text()
	sent_tokenize(input)
}
