package dom

import (
	"fmt"
)

type (
	// Navigator represents the JavaScript Navigator object.
	Navigator struct {
		*DOM
	}
)

// Language returns the user's preferred language.
func (n *Navigator) Language() (string, error) {
	v := n.Get("language")
	if v.Truthy() {
		return v.String(), nil
	}
	return "", fmt.Errorf("Language not found")
}

// MediaDevices returns a [*MediaDevices] object for accessing media devices.
func (n *Navigator) MediaDevices() *MediaDevices {
	return &MediaDevices{
		DOM: n.Get("mediaDevices"),
	}
}
