package dom

type (
	// MediaStream represents the JavaScript MediaStream object.
	MediaStream struct {
		*DOM
	}
)

// Active returns true if the media stream is active.
func (m *MediaStream) Active() bool {
	return m.Get("active").Bool()
}

// Id returns the media stream id.
func (m *MediaStream) Id() string {
	return m.Get("id").String()
}
