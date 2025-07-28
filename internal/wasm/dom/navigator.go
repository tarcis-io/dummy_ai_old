package dom

type (
	Navigator struct {
		*DOM
	}
)

func (n *Navigator) MediaDevices() *MediaDevices {
	return &MediaDevices{
		DOM: n.Get("mediaDevices"),
	}
}

func (n *Navigator) Language() (string, bool) {
	v := n.Get("language")
	if v.Truthy() {
		return v.String(), true
	}
	return "", false
}

func GetNavigator() *Navigator {
	return GetWindow().Navigator()
}
