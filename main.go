package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	allMessages := convertCSV("LigaTexts.csv")
	// messagesToTotalStats(allMessages)
	messagesToPersonStats(allMessages)
	// saveTotalStatsFiles(allMessages)
	// savePersonFiles()
	// makeWorkCloud()
	makeCircularCSV()
}

func makeWorkCloud() {
	countMap := map[string]int{}
	for _, word := range *totalStats.AllCleanWords {
		countMap[word]++
	}

	filename := fmt.Sprintf("cloud.txt")
	f, _ := os.OpenFile(filename, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)

	sep := ", "
	for k, v := range countMap {
		if _, err := f.WriteString(k + sep + strconv.Itoa(v) + "\n"); err != nil {
			panic(err)
		}
	}
	f.Close()

	// w := wordclouds.NewWordcloud(
	// 	countMap,
	// 	wordclouds.FontFile("./Roboto-Black.ttf"),
	// 	wordclouds.Height(2048),
	// 	wordclouds.Width(2048),
	// )
	// fmt.Println(countMap)
	// img := w.Draw()

	// outputFile, err := os.Create("./output.png")
	// if err != nil {
	// 	// Handle error
	// }

	// // Encode takes a writer interface and an image interface
	// // We pass it the File and the RGBA
	// png.Encode(outputFile, img)

	// // Don't forget to close files
	// outputFile.Close()

}
