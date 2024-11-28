package util

const (
	App = iota
)

var (
	allMessages = map[string]map[int]string{
		English: {
			App: "DummyAI",
		},
		Spanish: {
			App: "DummyAI",
		},
		Portuguese: {
			App: "DummyAI",
		},
	}
	messages = allMessages[language]
)

func AllMessages() map[string]map[int]string {

	return allMessages
}

func Messages() map[int]string {

	return messages
}

func Message(key int) string {

	return messages[key]
}
