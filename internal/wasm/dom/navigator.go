package dom

type (
	Navigator struct {
		*DOM
	}
)

func (n *Navigator) Language() (string, bool) {
	l := n.Get("language")
	if l.Truthy() {
		return l.String(), true
	}
	return "", false
}

func GetNavigator() *Navigator {
	return GetWindow().Navigator()
}
