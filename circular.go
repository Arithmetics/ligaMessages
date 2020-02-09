package main

import (
	"fmt"
	"os"
	"strconv"
)

var personNumKey = map[string]int{
	"Jimmy":   0,
	"Scott":   1,
	"Erik":    2,
	"Ian":     3,
	"Rengel":  4,
	"Kimble":  5,
	"Chad":    6,
	"Houman":  7,
	"Dehaan":  8,
	"Cam":     9,
	"Mark":    10,
	"Dubov":   11,
	"Clayton": 12,
	"Matt":    13,
	"Dylan":   14,
	"Brock":   15,
	"":        16,
}

// var personStats = map[string]*statBucket{

// type statBucket struct {
// 	Total           int             `json:"total"`
// 	RespondingTo    *map[string]int `json:"respondingTo"`
// 	ActionsReceived *map[string]int `json:"actionsReceived"`
// 	ActionsSent     *map[string]int `json:"actionsSent"`
// 	AllGoodWords    *[]string       `json:"allGoodWords"`
// 	AllWords        *[]string       `json:"allWords"`
// }

func makeCircularCSV() {
	filename := fmt.Sprintf("circular2017.csv")
	f, _ := os.OpenFile(filename, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)

	for originPerson, bucket := range personStats {
		for destPerson, v := range *bucket.RespondingTo {
			originID := strconv.Itoa(personNumKey[originPerson])
			destID := strconv.Itoa(personNumKey[destPerson])
			flow := strconv.Itoa(v)

			fullString := fmt.Sprintf("%s,%s,%s,%s,%s,%s,%s,%s,%s,%s,%s,%s,%s,%s,%s,%s,%s,%s,%s,", originID, originPerson, destID, destPerson, "0", "0", "0", flow, "", originPerson, originPerson, destPerson, destPerson, "0", "0", "0", flow, "1", "1")
			f.WriteString(fullString + "\n")
		}
	}
	f.Close()
}
