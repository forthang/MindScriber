package storage

import "errors"

type Storage struct {
	data map[string]string
}

func NewStorage() *Storage {
	return &Storage{
		data: make(map[string]string),
	}
}

func (rs *Storage) Get(key string) (string, error) {
	value, exists := rs.data[key]
	if !exists {
		return "", errors.New("key not found")
	}
	return value, nil
}

func (rs *Storage) Put(key string, value string) error {
	rs.data[key] = value
	return nil
}

func (rs *Storage) Post(key string, value string) error {
	if _, exists := rs.data[key]; exists {
		return errors.New("key already exists")
	}
	rs.data[key] = value
	return nil
}

func (rs *Storage) Delete(key string) error {
	if _, exists := rs.data[key]; !exists {
		return errors.New("key not found")
	}
	delete(rs.data, key)
	return nil
}
