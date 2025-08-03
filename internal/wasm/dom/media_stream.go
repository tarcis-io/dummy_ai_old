package dom

type (
	// MediaStream represents the JavaScript MediaStream object.
	MediaStream struct {
		*DOM
	}
)

// GetUserMedia requests access to the user's media devices.
func GetUserMedia(constraints map[string]any) (*MediaStream, error) {
	return GetMediaDevices().GetUserMedia(constraints)
}
