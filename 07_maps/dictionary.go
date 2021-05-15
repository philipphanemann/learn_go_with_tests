package main

// Dictionary mapping
type Dictionary map[string]string

// Search method to retrieve entry from Dictionary
func (d Dictionary) Search(word string) string {
	return d[word]
}
