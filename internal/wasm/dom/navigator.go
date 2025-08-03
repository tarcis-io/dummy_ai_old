package dom

type (
	// Navigator represents the JavaScript Navigator object.
	Navigator struct {
		*DOM
	}
)

// Language returns the user's preferred language.
func (n *Navigator) Language() (string, bool) {
	v := n.Get("language")
	if v.Truthy() {
		return v.String(), true
	}
	return "", false
}

// MediaDevices returns a [*MediaDevices] object for accessing media devices.
func (n *Navigator) MediaDevices() *MediaDevices {
	return &MediaDevices{
		DOM: n.Get("mediaDevices"),
	}
}

// GetNavigator returns the current Navigator object.
func GetNavigator() *Navigator {
	return GetWindow().Navigator()
}
