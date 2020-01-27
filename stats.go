package main

var personStats = map[string]*statBucket{
	"Jimmy":   &statBucket{Total: 0, RespondingTo: &map[string]int{}},
	"Scott":   &statBucket{Total: 0, RespondingTo: &map[string]int{}},
	"Erik":    &statBucket{Total: 0, RespondingTo: &map[string]int{}},
	"Ian":     &statBucket{Total: 0, RespondingTo: &map[string]int{}},
	"Rengel":  &statBucket{Total: 0, RespondingTo: &map[string]int{}},
	"Kimble":  &statBucket{Total: 0, RespondingTo: &map[string]int{}},
	"Chad":    &statBucket{Total: 0, RespondingTo: &map[string]int{}},
	"Houman":  &statBucket{Total: 0, RespondingTo: &map[string]int{}},
	"Dehaan":  &statBucket{Total: 0, RespondingTo: &map[string]int{}},
	"Cam":     &statBucket{Total: 0, RespondingTo: &map[string]int{}},
	"Mark":    &statBucket{Total: 0, RespondingTo: &map[string]int{}},
	"Dubov":   &statBucket{Total: 0, RespondingTo: &map[string]int{}},
	"Clayton": &statBucket{Total: 0, RespondingTo: &map[string]int{}},
	"Matt":    &statBucket{Total: 0, RespondingTo: &map[string]int{}},
	"Dylan":   &statBucket{Total: 0, RespondingTo: &map[string]int{}},
	"Brock":   &statBucket{Total: 0, RespondingTo: &map[string]int{}},
	"":        &statBucket{Total: 0, RespondingTo: &map[string]int{}},
}

var actionPrefixes = []string{"Emphasized", "Laughed", "Disliked", "Liked", "Questioned", "Loved"}

type statBucket struct {
	Total        int
	RespondingTo *map[string]int
}

func messagesToStats(messages []*message) {
	lastMessage := &message{}

	for i, message := range messages {
		statBucket := personStats[message.Person]
		statBucket.Total++

		if i != 0 {
			markRespondingTo(statBucket, lastMessage)
		}
		lastMessage = message
	}
}

func markRespondingTo(statBucket *statBucket, lastMessage *message) {
	lastPerson := lastMessage.Person
	respondingTo := *statBucket.RespondingTo
	respondingTo[lastPerson]++
}
