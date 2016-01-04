package main

import (
	"fmt"
	"github.com/baolidakai/hello/ir"
	"bufio"
	"os"
	"io"
	"strconv"
	"strings"
	"github.com/baolidakai/hello/compute"
	"golang.org/x/crypto/ssh/terminal"
)

// Stores the state of the terminal before making it raw
var regularState *terminal.State

func main() {
	test := ir.NewInformationRetrieval()
	test.Initialize("./data/")
	sc := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Bowen Search (! to end): ")
		sc.Scan()
		input := sc.Text()
		test.Query(input)
		if len(input) > 0 && string(input[0]) == "!" {
			fmt.Println("Thank you!")
			break
		}
	}
	if len(os.Args) > 1 {
		input := strings.Replace(strings.Join(os.Args[1:], ""), " ", "", -1)
		res, err := compute.Evaluate(input)
		if err != nil {
			fmt.Println("Error: " + err.Error())
			return
		}
		fmt.Printf("%s\n", strconv.FormatFloat(res, 'G', -1, 64))
		return
	}

	var err error
	regularState, err = terminal.MakeRaw(0)
	if err != nil {
		panic(err)
	}
	defer terminal.Restore(0, regularState)

	term := terminal.NewTerminal(os.Stdin, "> ")
	term.AutoCompleteCallback = handleKey
	for {
		text, err := term.ReadLine()
		if err != nil {
			if err == io.EOF {
				// Quit without error on Ctrl^D
				fmt.Println()
				break
			}
			panic(err)
		}

		text = strings.Replace(text, " ", "", -1)
		if text == "exit" || text == "quit" {
			break
		}

		res, err := compute.Evaluate(text)
		if err != nil {
			fmt.Println("Error: " + err.Error())
			continue
		}
		fmt.Printf("%s\n", strconv.FormatFloat(res, 'G', -1, 64))
	}
}

func handleKey(line string, pos int, key rune) (newLine string, newPos int, ok bool) {
	if key == '\x03' {
		fmt.Println()
		terminal.Restore(0, regularState)
		os.Exit(0)
	}
	return "", 0, false
}
