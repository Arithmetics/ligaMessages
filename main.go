package main

import "fmt"

func main() {
	allMessages := convertCSV("LigaTexts.csv")
	messagesToTotalStats(allMessages)
	// messagesToPersonStats(allMessages)
	// file, _ := json.MarshalIndent(personStats, "", " ")
	// _ = ioutil.WriteFile("test.json", file, 0644)
	fmt.Printf("%+v\n", totalStats)
}
