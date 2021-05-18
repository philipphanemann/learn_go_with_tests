package main

import "testing"

func TestSearch(t *testing.T) {
	dictionary := Dictionary{"test": "this is just a test"}

	t.Run("known word", func(t *testing.T) {
		got, _ := dictionary.Search("test")
		want := "this is just a test"
		assertStrings(t, got, want)
	})

	t.Run("unknown word", func(t *testing.T) {
		_, got := dictionary.Search("unknown")
		assertError(t, got, errKeyNotFound)
	})
}

func TestAdd(t *testing.T) {
	t.Run("new key", func(t *testing.T) {
		dictionary := Dictionary{}
		key := "test"
		definition := "this is just a test"
		err := dictionary.Add(key, definition)

		assertError(t, err, nil)
		assertDefinition(t, dictionary, key, definition)
	})

	t.Run("existing key", func(t *testing.T) {

		key := "test"
		definition := "this is just a test"
		dictionary := Dictionary{key: definition}

		err := dictionary.Add(key, "new test")

		assertError(t, err, errWordExists)
		assertDefinition(t, dictionary, key, definition)
	})
}

func TestUpdate(t *testing.T) {
	t.Run("update existing entry", func(t *testing.T) {
		key := "test"
		definition := "this is just a test"
		dictionary := Dictionary{key: definition}
		newDefinition := "new definition"

		err := dictionary.Update(key, newDefinition)

		assertError(t, err, nil)
		assertDefinition(t, dictionary, key, newDefinition)
	})

	t.Run("new word", func(t *testing.T) {
		key := "test"
		definition := "this is just a test"
		dictionary := Dictionary{}

		err := dictionary.Update(key, definition)

		assertError(t, err, errWordDoesNotExist)
	})

}

func TestDelete(t *testing.T) {
	key := "test"
	definition := "this is just a test"
	dictionary := Dictionary{key: definition}

	dictionary.Delete(key)

	_, err := dictionary.Search(key)
	if err != errKeyNotFound {
		t.Errorf("Expected %q to be deleted", key)
	}
}

func assertDefinition(t testing.TB, dictionary Dictionary, key, definition string) {
	t.Helper()

	got, err := dictionary.Search(key)
	if err != nil {
		t.Fatal("should find added word:", err)
	}

	if got != definition {
		t.Errorf("got %q want %q", got, definition)
	}
}

func assertError(t testing.TB, got, want error) {
	t.Helper()
	if got != want {
		t.Errorf("got error %q want %q", got, want)
	}
}

func assertStrings(t testing.TB, got string, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got %q want %q given, %q", got, want, "test")
	}
}
