package main

func main() {
	allMessages := convertCSV("LigaTexts.csv")
	// messagesToTotalStats(allMessages)
	messagesToPersonStats(allMessages)
	// saveTotalStatsFiles()
	savePersonFiles()
}
