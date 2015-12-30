package main

import (
	"fmt"
)

func main() {
	ParseData()
}

var (
	dataFile = "data/brew-dependencies.txt"
)

func ParseData() {
	data, err := Asset(dataFile)
	if err != nil {
		panic(fmt.Sprintf("Data file [%s] not embedded!", dataFile))
	}

	dataAsString := string(data[:])
	fmt.Println("aaa")
	fmt.Println(dataAsString)

}
