package dom

type (
	// MediaQueryList represents the JavaScript MediaQueryList object.
	MediaQueryList struct {
		*DOM
	}
)

// Matches returns a boolean indicating
// whether the current media query matches.
func (m *MediaQueryList) Matches() bool {
	return m.Get("matches").Bool()
}

// Media returns a string representing the serialized media query.
func (m *MediaQueryList) Media() string {
	return m.Get("media").String()
}
