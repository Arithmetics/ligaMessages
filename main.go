package main

import (
	"encoding/json"
	"io/ioutil"
)

func main() {
	allMessages := convertCSV("LigaTexts.csv")
	messagesToStats(allMessages)

	file, _ := json.MarshalIndent(personStats, "", " ")
	_ = ioutil.WriteFile("test.json", file, 0644)
}
