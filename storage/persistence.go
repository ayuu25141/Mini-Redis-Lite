package storage

import (
	"encoding/gob"
	"os"
)

func SaveToFile(store *Store, filename string) error {
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	return gob.NewEncoder(file).Encode(store.All())
}

func LoadFromFile(store *Store, filename string) error {
	if _, err := os.Stat(filename); os.IsNotExist(err) {
		return nil // file not exists
	}

	file, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	data := make(map[string]string)
	err = gob.NewDecoder(file).Decode(&data)
	if err != nil {
		return err
	}

	for k, v := range data {
		store.Set(k, v)
	}
	return nil
}
