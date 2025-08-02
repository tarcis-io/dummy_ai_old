package dom

type (
	// MediaDevices represents the JavaScript MediaDevices object.
	MediaDevices struct {
		*DOM
	}
)

// GetUserMedia requests access to use a media input based on the specified constraints
// and returns a *MediaStream object.
func (m *MediaDevices) GetUserMedia(constraints map[string]any) (*MediaStream, error) {
	v, err := m.Call("getUserMedia", constraints).Await()
	if err != nil {
		return nil, err
	}
	return &MediaStream{
		DOM: v,
	}, nil
}

// GetMediaDevices returns the web page's current MediaDevices object.
func GetMediaDevices() *MediaDevices {
	return GetNavigator().MediaDevices()
}
