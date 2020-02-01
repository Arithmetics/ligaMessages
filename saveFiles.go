package main

import (
	"fmt"
	"os"
)

func savePersonFiles() {
	for k, v := range personStats {
		filename := fmt.Sprintf("%s_clean.txt", k)
		f, _ := os.OpenFile(filename, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)

		sep := " "
		for _, word := range *v.AllGoodWords {
			if _, err := f.WriteString(word + sep); err != nil {
				panic(err)
			}
		}
		f.Close()
	}

	for k, v := range personStats {
		filename := fmt.Sprintf("%s_all.txt", k)
		f, _ := os.OpenFile(filename, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)

		sep := " "
		for _, word := range *v.AllWords {
			if _, err := f.WriteString(word + sep); err != nil {
				panic(err)
			}
		}
		f.Close()
	}
}

func saveTotalStatsFiles() {
	// write all words to file
	// If the file doesn't exist, create it, or append to the file
	f, _ := os.OpenFile("allWords.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)

	sep := " "
	for _, word := range *totalStats.AllWords {
		if _, err := f.WriteString(word + sep); err != nil {
			panic(err)
		}
	}
	f.Close()

	// TotalCount
	fmt.Println("Total Text Count")
	fmt.Println(totalStats.TotalCount)

	// Seperated by hour
	f, _ = os.OpenFile("hourlyInterval.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)

	for k, v := range *totalStats.TotalByHour {
		if _, err := f.WriteString(fmt.Sprintf("%d,%d\n", k, v)); err != nil {
			panic(err)
		}
	}
	f.Close()

	// Seperated by hour
	f, _ = os.OpenFile("dayInterval.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)

	for k, v := range *totalStats.TotalByDoW {
		if _, err := f.WriteString(fmt.Sprintf("%d,%d\n", k, v)); err != nil {
			panic(err)
		}
	}
	f.Close()

	// Seperated by 30 min over year
	f, _ = os.OpenFile("yearIntervals.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)

	for k, v := range *totalStats.IntervalTotals {
		if _, err := f.WriteString(fmt.Sprintf("%v,%d\n", k, v)); err != nil {
			panic(err)
		}
	}
	f.Close()

}
