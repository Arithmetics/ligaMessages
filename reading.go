package main

import (
	"encoding/csv"
	"os"
	"regexp"
	"strconv"
	"time"
)

var phoneLookup = map[string]string{
	"16505208371":             "Jimmy",
	"17144032278":             "Scott",
	"15035513466":             "Erik",
	"14252837314":             "Ian",
	"15035043582":             "Rengel",
	"15034598925":             "Kimble",
	"15035806566":             "Chad",
	"15037895186":             "Houman",
	"15039292572":             "Dehaan",
	"19714099773":             "Cam",
	"14153064486":             "Mark",
	"15039972480":             "Dubov",
	"15035107079":             "Clayton",
	"15038833302":             "Matt",
	"15039100235":             "Dylan",
	"dylanbranch9@icloud.com": "Dylan",
	"":                        "Brock",
	"+16505208371":            "Jimmy",
	"+17144032278":            "Scott",
	"+15035513466":            "Erik",
	"+14252837314":            "Ian",
	"+15035043582":            "Rengel",
	"+15034598925":            "Kimble",
	"+15035806566":            "Chad",
	"+15037895186":            "Houman",
	"+15039292572":            "Dehaan",
	"+19714099773":            "Cam",
	"+14153064486":            "Mark",
	"+15039972480":            "Dubov",
	"+15035107079":            "Clayton",
	"+15038833302":            "Matt",
	"+15039100235":            "Dylan",
}

type message struct {
	Text      string
	Timestamp time.Time
	Person    string
}

func readCSV(filename string) ([][]string, error) {
	// Open CSV file
	f, err := os.Open(filename)
	if err != nil {
		return [][]string{}, err
	}
	defer f.Close()
	// Read File into a Variable
	lines, err := csv.NewReader(f).ReadAll()
	if err != nil {
		return [][]string{}, err
	}

	return lines, nil
}

func convertCSV(filename string) []*message {
	allMessages := []*message{}
	lines, err := readCSV(filename)
	if err != nil {
		panic(err)
	}
	re := regexp.MustCompile("[[:^ascii:]]")
	// Loop through lines & turn into object
	for _, line := range lines {
		ts, _ := strconv.Atoi(line[2])
		person := phoneLookup[line[3]]
		cleanText := re.ReplaceAllLiteralString(line[0], "")
		data := message{
			Text:      cleanText,
			Timestamp: getAppleDate(int64(ts)),
			Person:    person,
		}
		allMessages = append(allMessages, &data)
	}

	return allMessages
}

func getAppleDate(appleNanos int64) time.Time {
	raw := int64(appleNanos)
	dividedRaw := raw / 1000000000

	layout := "2006-01-02"
	str := "2001-01-01"
	appleZeroDate, _ := time.Parse(layout, str)

	secsSinceEpoch := appleZeroDate.Unix()

	totalSecs := secsSinceEpoch + dividedRaw

	finalTime := time.Unix(totalSecs, 0)
	return finalTime
}
