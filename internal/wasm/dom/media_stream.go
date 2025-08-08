package dom

type (
	// MediaStream represents the JavaScript MediaStream object.
	MediaStream struct {
		*DOM
	}
)

func (m *MediaStream) Active() bool {
	return m.Get("active").Bool()
}

func (m *MediaStream) Id() string {
	return m.Get("id").String()
}
