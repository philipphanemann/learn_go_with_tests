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
	key := "test"
	definition := "this is just a test"
	dictionary := Dictionary{key: definition}
	newDefinition := "new definition"

	dictionary.Update(key, newDefinition)

	assertDefinition(t, dictionary, key, newDefinition)

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
