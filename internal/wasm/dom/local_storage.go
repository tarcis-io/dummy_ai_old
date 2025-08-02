package dom

type (
	// LocalStorage represents the JavaScript LocalStorage object.
	LocalStorage struct {
		*DOM
	}
)

// SetItem sets the value of the specified key in the *LocalStorage object.
func (l *LocalStorage) SetItem(key, value string) {
	l.Call("setItem", key, value)
}

// GetItem retrieves the value of the specified key from the *LocalStorage object.
func (l *LocalStorage) GetItem(key string) (string, bool) {
	v := l.Call("getItem", key)
	if v.Truthy() {
		return v.String(), true
	}
	return "", false
}

// GetLocalStorage returns the web page's current *LocalStorage object.
func GetLocalStorage() *LocalStorage {
	return GetWindow().LocalStorage()
}
