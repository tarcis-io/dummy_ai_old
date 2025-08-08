package dom

type (
	// MediaDevices represents the JavaScript MediaDevices object.
	MediaDevices struct {
		*DOM
	}
)

// GetUserMedia requests access to the user's media devices.
func (m *MediaDevices) GetUserMedia(constraints map[string]any) (*MediaStream, error) {
	v, err := m.Call("getUserMedia", constraints).Await()
	if err != nil {
		return nil, err
	}
	return &MediaStream{
		DOM: v,
	}, nil
}
