package main

import (
	"fmt"
)

func main() {
	fmt.Println("yooooo")
	allMessages := convertCSV("LigaTexts.csv")
	fmt.Println(len(allMessages))
}
