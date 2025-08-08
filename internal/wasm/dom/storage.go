package dom

import (
	"fmt"
)

type (
	// Storage represents the JavaScript Storage object.
	Storage struct {
		*DOM
	}
)

// SetItem sets the value for the specified key in the Storage object.
func (s *Storage) SetItem(key, value string) {
	s.Call("setItem", key, value)
}

// GetItem returns the value for the specified key in the Storage object.
func (s *Storage) GetItem(key string) (string, error) {
	v := s.Call("getItem", key)
	if v.Truthy() {
		return v.String(), nil
	}
	return "", fmt.Errorf("Item with key %q not found", key)
}

// GetLocalStorage returns the current local Storage object.
func GetLocalStorage() *Storage {
	return GetWindow().LocalStorage()
}

// GetSessionStorage returns the current session Storage object.
func GetSessionStorage() *Storage {
	return GetWindow().SessionStorage()
}
