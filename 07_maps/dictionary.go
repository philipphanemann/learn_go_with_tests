package main

import "errors"

// Dictionary mapping
type Dictionary map[string]string

var errKeyNotFound = errors.New("could not find the word you were looking for")

// Search method to retrieve entry from Dictionary
func (d Dictionary) Search(word string) (string, error) {
	definition, ok := d[word]

	if !ok {
		return "", errKeyNotFound
	}
	return definition, nil
}

// Add method for adding content to dictionary
func (d Dictionary) Add(key, value string) {
	d[key] = value
}
