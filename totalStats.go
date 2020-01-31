package main

import (
	"fmt"
	"strings"
	"time"
)

type cumulativeStat struct {
	TotalCount     int
	TotalByHour    *map[int]int
	TotalByDoW     *map[int]int
	IntervalTotals *map[time.Time]int
	AllWords       *[]string
}

var layout = "2006-01-02"
var start = "2019-01-01"
var end = "2020-01-01"
var startingDate, _ = time.Parse(layout, start)
var endingDate, _ = time.Parse(layout, end)

var totalStats = cumulativeStat{TotalByHour: &map[int]int{}, TotalByDoW: &map[int]int{}, IntervalTotals: &map[time.Time]int{}, AllWords: &[]string{}}

var yearsIntervals = fillIntervals(startingDate, endingDate)

func fillIntervals(startDate time.Time, endDate time.Time) []time.Time {
	var allIntervals []time.Time
	secondsInInterval := time.Second * time.Duration(1800) // 30 mins
	rollingDate := startDate
	for {
		if rollingDate.After(endDate) {
			break
		}
		allIntervals = append(allIntervals, rollingDate)
		rollingDate = rollingDate.Add(secondsInInterval)
	}
	return allIntervals
}

func messagesToTotalStats(messages []*message) {
	for m, message := range messages {
		fmt.Printf("%+v\n", m)
		// allText
		addMessageToWords := true
		for _, prefix := range actionPrefixes {
			if strings.HasPrefix(message.Text, prefix) {
				addMessageToWords = false
			}
		}
		cloudifiedMessages := cloudifyMessage(message.Text)

		if addMessageToWords {
			*totalStats.AllWords = append(*totalStats.AllWords, cloudifiedMessages...)
		}
		// intervalTotals
		for i, interval := range yearsIntervals {
			if len(yearsIntervals) > (i+2) && message.Timestamp.After(yearsIntervals[i]) && message.Timestamp.Before(yearsIntervals[i+1]) {
				intervalTotals := *totalStats.IntervalTotals
				intervalTotals[interval]++
				break
			}
		}
		//totalByHour
		totalByHour := *totalStats.TotalByHour
		totalByHour[message.Timestamp.Hour()]++
		//totalByDayOfWeek
		totalByDoW := *totalStats.TotalByDoW
		totalByDoW[int(message.Timestamp.Weekday())]++
		// total
		totalStats.TotalCount++
	}
}
