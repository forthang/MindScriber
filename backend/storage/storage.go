package storage

import "errors"

type RamStorage struct {
	data map[string]string
}

func NewRamStorage() *RamStorage {
	return &RamStorage{
		data: make(map[string]string),
	}
}

func (rs *RamStorage) Get(key string) (string, error) {
	value, exists := rs.data[key]
	if !exists {
		return "", errors.New("key not found")
	}
	return value, nil
}

func (rs *RamStorage) Put(key string, value string) error {
	rs.data[key] = value
	return nil
}

func (rs *RamStorage) Post(key string, value string) error {
	if _, exists := rs.data[key]; exists {
		return errors.New("key already exists")
	}
	rs.data[key] = value
	return nil
}

func (rs *RamStorage) Delete(key string) error {
	if _, exists := rs.data[key]; !exists {
		return errors.New("key not found")
	}
	delete(rs.data, key)
	return nil
}
