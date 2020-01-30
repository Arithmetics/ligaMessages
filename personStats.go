package main

import (
	"fmt"
	"strings"
)

var personStats = map[string]*statBucket{
	"Jimmy":   &statBucket{Total: 0, RespondingTo: &map[string]int{}, ActionsReceived: &map[string]int{}, ActionsSent: &map[string]int{}},
	"Scott":   &statBucket{Total: 0, RespondingTo: &map[string]int{}, ActionsReceived: &map[string]int{}, ActionsSent: &map[string]int{}},
	"Erik":    &statBucket{Total: 0, RespondingTo: &map[string]int{}, ActionsReceived: &map[string]int{}, ActionsSent: &map[string]int{}},
	"Ian":     &statBucket{Total: 0, RespondingTo: &map[string]int{}, ActionsReceived: &map[string]int{}, ActionsSent: &map[string]int{}},
	"Rengel":  &statBucket{Total: 0, RespondingTo: &map[string]int{}, ActionsReceived: &map[string]int{}, ActionsSent: &map[string]int{}},
	"Kimble":  &statBucket{Total: 0, RespondingTo: &map[string]int{}, ActionsReceived: &map[string]int{}, ActionsSent: &map[string]int{}},
	"Chad":    &statBucket{Total: 0, RespondingTo: &map[string]int{}, ActionsReceived: &map[string]int{}, ActionsSent: &map[string]int{}},
	"Houman":  &statBucket{Total: 0, RespondingTo: &map[string]int{}, ActionsReceived: &map[string]int{}, ActionsSent: &map[string]int{}},
	"Dehaan":  &statBucket{Total: 0, RespondingTo: &map[string]int{}, ActionsReceived: &map[string]int{}, ActionsSent: &map[string]int{}},
	"Cam":     &statBucket{Total: 0, RespondingTo: &map[string]int{}, ActionsReceived: &map[string]int{}, ActionsSent: &map[string]int{}},
	"Mark":    &statBucket{Total: 0, RespondingTo: &map[string]int{}, ActionsReceived: &map[string]int{}, ActionsSent: &map[string]int{}},
	"Dubov":   &statBucket{Total: 0, RespondingTo: &map[string]int{}, ActionsReceived: &map[string]int{}, ActionsSent: &map[string]int{}},
	"Clayton": &statBucket{Total: 0, RespondingTo: &map[string]int{}, ActionsReceived: &map[string]int{}, ActionsSent: &map[string]int{}},
	"Matt":    &statBucket{Total: 0, RespondingTo: &map[string]int{}, ActionsReceived: &map[string]int{}, ActionsSent: &map[string]int{}},
	"Dylan":   &statBucket{Total: 0, RespondingTo: &map[string]int{}, ActionsReceived: &map[string]int{}, ActionsSent: &map[string]int{}},
	"Brock":   &statBucket{Total: 0, RespondingTo: &map[string]int{}, ActionsReceived: &map[string]int{}, ActionsSent: &map[string]int{}},
	"":        &statBucket{Total: 0, RespondingTo: &map[string]int{}, ActionsReceived: &map[string]int{}, ActionsSent: &map[string]int{}},
}

var actionPrefixes = []string{"Emphasized", "Laughed", "Disliked", "Liked", "Questioned", "Loved"}

type statBucket struct {
	Total           int             `json:"total"`
	RespondingTo    *map[string]int `json:"respondingTo"`
	ActionsReceived *map[string]int `json:"actionsReceived"`
	ActionsSent     *map[string]int `json:"actionsSent"`
}

func messagesToPersonStats(messages []*message) {
	lastMessage := &message{}

	for i, message := range messages {
		fmt.Println(i)
		statBucket := personStats[message.Person]
		statBucket.Total++

		if i != 0 {
			markRespondingTo(statBucket, lastMessage)
			markActions(statBucket, *message, messages)
		}
		lastMessage = message
	}
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
