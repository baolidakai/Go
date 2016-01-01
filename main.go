package main

import (
	"fmt"
	"github.com/baolidakai/hello/ir"
	"bufio"
	"os"
)

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
}
