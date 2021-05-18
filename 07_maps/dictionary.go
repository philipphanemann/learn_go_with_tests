package main

// Dictionary mapping
type Dictionary map[string]string

const (
	errKeyNotFound      = DictionaryErr("could not find the word you were looking for")
	errWordExists       = DictionaryErr("cannot add word because it already exists")
	errWordDoesNotExist = DictionaryErr("cannot update word because it does not exist")
)

type DictionaryErr string

func (e DictionaryErr) Error() string {
	return string(e)
}

// Search method to retrieve entry from Dictionary
func (d Dictionary) Search(word string) (string, error) {
	definition, ok := d[word]

	if !ok {
		return "", errKeyNotFound
	}
	return definition, nil
}

// Add method for adding content to dictionary
func (d Dictionary) Add(key, value string) error {
	_, err := d.Search(key)

	switch err {
	case errKeyNotFound:
		d[key] = value
	case nil:
		return errWordExists
	default:
		return err
	}
	return nil
}

// Update dictionary entry
func (d Dictionary) Update(key, value string) error {
	_, err := d.Search(key)
	if err == errKeyNotFound {
		return errWordDoesNotExist
	}

	d[key] = value

	return nil
}
