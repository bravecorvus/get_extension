package main

import (
	"fmt"
	"io/ioutil"
	"os"

	"github.com/h2non/filetype"
)

func main() {
	if len(os.Args) < 2 {
		panic("Please pass full filepath")
	}

	fileData, err := ioutil.ReadFile(os.Args[1])
	if err != nil {
		panic("ioutil.ReadFile(\"" + os.Args[1] + "\") Exception: " + err.Error())
	}

	kind, err := filetype.Match(fileData)
	if err != nil {
		panic("filetype.Match(fileData) Exception: " + err.Error())
	}

	fmt.Println(kind.Extension)
}
