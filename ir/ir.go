package ir

import (
	"fmt"
	"io/ioutil"
	"strings"
	"regexp"
)

//Example
//test := ir.NewInformationRetrieval()
//test.Initialize("./data/")
//var input string
//for {
//	fmt.Print("Bowen Search (! to end): ")
//	fmt.Scanln(&input)
//	test.Query(input)
//	if string(input[0]) == "!" {
//		fmt.Println("Thank you!")
//		break
//	}
//}
//

type InformationRetrieval struct {
	dict map[string][]int // Map from word to document ids
	name map[int]string // Map from id to filename
}

func (ir *InformationRetrieval) PushBack(key string, value int) {
	curr := ir.dict[key]
	if len(curr) == 0 || curr[len(curr) - 1] != value {
		ir.dict[key] = append(curr, value)
	}
}

func (ir *InformationRetrieval) Initialize(dir string) {
	// Initialize the inverted index system under dir
	files, _ := ioutil.ReadDir(dir)
	pattern, _ := regexp.Compile("[A-Za-z]*")
	for id, file := range files {
		filename := file.Name()
		ir.name[id] = filename
		content, _ := ioutil.ReadFile(dir + filename)
		// Parse the file
		result := pattern.FindAllString(string(content), -1)
		for _, word := range result {
			word = strings.ToLower(word)
			if len(word) > 0 {
				ir.PushBack(word, id)
			}
		}
	}
}

func (ir *InformationRetrieval) Query(query string) {
	// Translate the query into words
	pattern, _ := regexp.Compile("[A-Za-z]*")
	words := pattern.FindAllString(query, -1)
	locked := true
	// Merge recursively
	var rtn []int
	for _, word := range words {
		if len(word) > 0 {
			if locked {
				rtn = ir.dict[word]
				locked = false
			} else {
				rtn = Merge(rtn, ir.dict[word])
			}
		}
	}
	// Translate into filenames
	for _, id := range rtn {
		fmt.Println(ir.name[id])
	}
}

func Merge(ids1, ids2 []int) []int {
	var rtn []int
	n1 := len(ids1)
	n2 := len(ids2)
	i := 0
	j := 0
	for i < n1 && j < n2 {
		if ids1[i] > ids2[j] {
			j = j + 1
		} else if ids2[j] > ids1[i] {
			i = i + 1
		} else {
			rtn = append(rtn, ids1[i])
			i = i + 1
			j = j + 1
		}
	}
	return rtn
}

func NewInformationRetrieval() *InformationRetrieval {
	return &InformationRetrieval{
		dict: make(map[string][]int),
		name: make(map[int]string),
	}
}
