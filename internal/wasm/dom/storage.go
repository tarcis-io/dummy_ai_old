package dom

type (
	Storage struct {
		*DOM
	}
)

func (s *Storage) SetItem(key, value string) {
	s.Call("setItem", key, value)
}

func (s *Storage) GetItem(key string) (string, bool) {
	v := s.Call("getItem", key)
	if v.Truthy() {
		return v.String(), true
	}
	return "", false
}

func GetLocalStorage() *Storage {
	return GetWindow().LocalStorage()
}

func GetSessionStorage() *Storage {
	return GetWindow().SessionStorage()
}
