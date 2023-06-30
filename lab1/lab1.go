package main

//Run the program with test file name as argument

import (
	"fmt"
	"os"
	"strings"

	"github.com/jdkato/prose/v2"
)

func validate() {
	if len(os.Args) < 2 {
		panic("Missing test file name argument, please provide a test file name as the second argument")
	}
	//check if file exists
	if _, err := os.Stat(os.Args[1]); os.IsNotExist(err) {
		panic("Test file does not exist, please provide a valid test file name as the second argument")
	}
}

func readFile(name string) string {
	file, err := os.ReadFile(name)
	if err != nil {
		panic(err)
	}
	str := string(file)
	return str
}

func tokenizer(txt string) []string {
	doc, err := prose.NewDocument(txt)
	if err != nil {
		panic(err)
	}
	tokens := []string{}
	unwanted := map[string]struct{}{
		"(":  {},
		")":  {},
		",":  {},
		":":  {},
		".":  {},
		"''": {},
		"``": {},
		"$":  {},
		"#":  {},
	}
	for _, tok := range doc.Tokens() {
		tok.Text = strings.ToLower(tok.Text)
		if _, ok := unwanted[tok.Tag]; ok {
			continue
		}
		tokens = append(tokens, tok.Text)
	}
	return tokens
}

func count(s []string) ([]string, map[string]int) {
	keys := make(map[string]int)
	list := []string{}
	for _, entry := range s {
		val := keys[entry]
		keys[entry] = val + 1
	}
	return list, keys
}

func main() {
	validate()
	pName := os.Args[1]
	var txt = readFile(pName)
	_, keys := count(tokenizer(txt))
	for k, v := range keys {
		fmt.Println(k, v)
	}
}
