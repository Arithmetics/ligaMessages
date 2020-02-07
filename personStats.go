package main

import (
	"fmt"
	"github.com/bbalet/stopwords"
	"strings"
)

var personStats = map[string]*statBucket{
	"Jimmy":   &statBucket{AllWords: &[]string{}, AllGoodWords: &[]string{}, Total: 0, RespondingTo: &map[string]int{}, ActionsReceived: &map[string]int{}, ActionsSent: &map[string]int{}},
	"Scott":   &statBucket{AllWords: &[]string{}, AllGoodWords: &[]string{}, Total: 0, RespondingTo: &map[string]int{}, ActionsReceived: &map[string]int{}, ActionsSent: &map[string]int{}},
	"Erik":    &statBucket{AllWords: &[]string{}, AllGoodWords: &[]string{}, Total: 0, RespondingTo: &map[string]int{}, ActionsReceived: &map[string]int{}, ActionsSent: &map[string]int{}},
	"Ian":     &statBucket{AllWords: &[]string{}, AllGoodWords: &[]string{}, Total: 0, RespondingTo: &map[string]int{}, ActionsReceived: &map[string]int{}, ActionsSent: &map[string]int{}},
	"Rengel":  &statBucket{AllWords: &[]string{}, AllGoodWords: &[]string{}, Total: 0, RespondingTo: &map[string]int{}, ActionsReceived: &map[string]int{}, ActionsSent: &map[string]int{}},
	"Kimble":  &statBucket{AllWords: &[]string{}, AllGoodWords: &[]string{}, Total: 0, RespondingTo: &map[string]int{}, ActionsReceived: &map[string]int{}, ActionsSent: &map[string]int{}},
	"Chad":    &statBucket{AllWords: &[]string{}, AllGoodWords: &[]string{}, Total: 0, RespondingTo: &map[string]int{}, ActionsReceived: &map[string]int{}, ActionsSent: &map[string]int{}},
	"Houman":  &statBucket{AllWords: &[]string{}, AllGoodWords: &[]string{}, Total: 0, RespondingTo: &map[string]int{}, ActionsReceived: &map[string]int{}, ActionsSent: &map[string]int{}},
	"Dehaan":  &statBucket{AllWords: &[]string{}, AllGoodWords: &[]string{}, Total: 0, RespondingTo: &map[string]int{}, ActionsReceived: &map[string]int{}, ActionsSent: &map[string]int{}},
	"Cam":     &statBucket{AllWords: &[]string{}, AllGoodWords: &[]string{}, Total: 0, RespondingTo: &map[string]int{}, ActionsReceived: &map[string]int{}, ActionsSent: &map[string]int{}},
	"Mark":    &statBucket{AllWords: &[]string{}, AllGoodWords: &[]string{}, Total: 0, RespondingTo: &map[string]int{}, ActionsReceived: &map[string]int{}, ActionsSent: &map[string]int{}},
	"Dubov":   &statBucket{AllWords: &[]string{}, AllGoodWords: &[]string{}, Total: 0, RespondingTo: &map[string]int{}, ActionsReceived: &map[string]int{}, ActionsSent: &map[string]int{}},
	"Clayton": &statBucket{AllWords: &[]string{}, AllGoodWords: &[]string{}, Total: 0, RespondingTo: &map[string]int{}, ActionsReceived: &map[string]int{}, ActionsSent: &map[string]int{}},
	"Matt":    &statBucket{AllWords: &[]string{}, AllGoodWords: &[]string{}, Total: 0, RespondingTo: &map[string]int{}, ActionsReceived: &map[string]int{}, ActionsSent: &map[string]int{}},
	"Dylan":   &statBucket{AllWords: &[]string{}, AllGoodWords: &[]string{}, Total: 0, RespondingTo: &map[string]int{}, ActionsReceived: &map[string]int{}, ActionsSent: &map[string]int{}},
	"Brock":   &statBucket{AllWords: &[]string{}, AllGoodWords: &[]string{}, Total: 0, RespondingTo: &map[string]int{}, ActionsReceived: &map[string]int{}, ActionsSent: &map[string]int{}},
	"":        &statBucket{AllWords: &[]string{}, AllGoodWords: &[]string{}, Total: 0, RespondingTo: &map[string]int{}, ActionsReceived: &map[string]int{}, ActionsSent: &map[string]int{}},
}

var actionPrefixes = []string{"Emphasized", "Laughed", "Disliked", "Liked", "Questioned", "Loved"}

type statBucket struct {
	Total           int             `json:"total"`
	RespondingTo    *map[string]int `json:"respondingTo"`
	ActionsReceived *map[string]int `json:"actionsReceived"`
	ActionsSent     *map[string]int `json:"actionsSent"`
	AllGoodWords    *[]string       `json:"allGoodWords"`
	AllWords        *[]string       `json:"allWords"`
}

func messagesToPersonStats(messages []*message) {
	lastMessage := &message{}

	for i, message := range messages {
		fmt.Println(i)
		statBucket := personStats[message.Person]
		statBucket.Total++

		addMessageToWords := true
		for _, prefix := range actionPrefixes {
			if strings.HasPrefix(message.Text, prefix) {
				addMessageToWords = false
			}
		}

		cloudifiedMessages := cloudifyMessage(message.Text)
		regularMessages := regularCleanMessage(message.Text)

		if addMessageToWords {
			*statBucket.AllGoodWords = append(*statBucket.AllGoodWords, cloudifiedMessages...)
			*statBucket.AllWords = append(*statBucket.AllWords, regularMessages...)
		}

		if i != 0 {
			markRespondingTo(statBucket, lastMessage)
			markActions(statBucket, *message, messages)
		}
		lastMessage = message
	}
}

func cloudifyMessage(text string) []string {
	cleanContent := stopwords.CleanString(text, "en", true)
	splitContent := strings.Split(cleanContent, " ")

	filtered := []string{}
	for _, word := range splitContent {
		if word != "" && word != "I" && word != "i" {
			filtered = append(filtered, word)
		}
	}

	return filtered
}

func regularCleanMessage(text string) []string {
	splitContent := strings.Split(text, " ")

	filtered := []string{}
	for _, word := range splitContent {
		if word != "" {
			filtered = append(filtered, word)
		}
	}

	return filtered
}

func markRespondingTo(statBucket *statBucket, lastMessage *message) {
	lastPerson := lastMessage.Person
	respondingTo := *statBucket.RespondingTo
	respondingTo[lastPerson]++
}

func markActions(statBucket *statBucket, message message, allMessages []*message) {
	for _, prefix := range actionPrefixes {
		if strings.HasPrefix(message.Text, prefix) {
			actionsSent := *statBucket.ActionsSent
			actionsSent[prefix]++

			for i, originalMessage := range allMessages {
				matchingText := fmt.Sprintf("%s %s", prefix, originalMessage.Text)
				if matchingText == message.Text {
					senderPerson := allMessages[i+1].Person
					senderStatBucket := personStats[senderPerson]
					senderActionsReceived := *senderStatBucket.ActionsReceived
					senderActionsReceived[prefix]++
				}
			}
		}
	}
}
