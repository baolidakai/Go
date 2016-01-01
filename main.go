package main

import (
	"fmt"
	"github.com/baolidakai/hello/ir"
)

func main() {
	test := ir.NewInformationRetrieval()
	test.Initialize("./data/")
	var input string
	for {
		fmt.Print("Bowen Search (! to end): ")
		fmt.Scanln(&input)
		test.Query(input)
		if string(input[0]) == "!" {
			fmt.Println("Thank you!")
			break
		}
	}
}
