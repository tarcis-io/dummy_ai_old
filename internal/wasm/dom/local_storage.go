package dom

type (
	LocalStorage struct {
		*DOM
	}
)

func (l *LocalStorage) SetItem(key, value string) {
	l.Call("setItem", key, value)
}

func (l *LocalStorage) GetItem(key string) (string, bool) {
	v := l.Call("getItem", key)
	if v.Truthy() {
		return v.String(), true
	}
	return "", false
}

func GetLocalStorage() *LocalStorage {
	return GetWindow().LocalStorage()
}
