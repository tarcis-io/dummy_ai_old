package dom

type (
	MediaDevices struct {
		*DOM
	}
)

func (m *MediaDevices) GetUserMedia(constraints map[string]any) (*MediaStream, error) {
	v, err := m.Call("getUserMedia", constraints).Await()
	if err != nil {
		return nil, err
	}
	return &MediaStream{
		DOM: v,
	}, nil
}

func GetMediaDevices() *MediaDevices {
	return GetNavigator().MediaDevices()
}
