package util

const (
	App = iota
)

var (
	allMessages = map[string]map[int]string{
		English: {
			App: "DummyAI",
		},
	}
	messages = allMessages[language]
)
