package maps

import "testing"

func TestSearch(t *testing.T) {
	dictionary := Dictionary{
		"test":   "this is just a test",
		"test02": "this is the second test",
	}

	t.Run("known word", func(t *testing.T) {
		key := "test"
		got, _ := dictionary.Search(key)
		want := "this is just a test"
		assertMapStrings(t, got, want, key)
	})

	t.Run("unknown word", func(t *testing.T) {
		key := "unknown"
		_, err := dictionary.Search(key)
		assertError(t, err, ErrNotFound)
	})
}

func TestAdd(t *testing.T) {
	t.Run("new word", func(t *testing.T) {
		dictionary := Dictionary{}
		key := "test"
		value := "this is just a test"
		err := dictionary.Add(key, value)
		assertNoError(t, err)

		assertMapContent(t, dictionary, key, value)
	})

	t.Run("existing word", func(t *testing.T) {
		key := "test"
		value := "this is just a test"
		dictionary := Dictionary{key: value}
		err := dictionary.Add(key, "new test")

		assertError(t, err, ErrWordExists)
		assertMapContent(t, dictionary, key, value)
	})
}

func TestUpdate(t *testing.T) {
	t.Run("existing word", func(t *testing.T) {
		word := "test"
		definition := "this is just a test"
		dictionary := Dictionary{word: definition}
		newDefinition := "new definition"
		err := dictionary.Update(word, newDefinition)
		assertNoError(t, err)
		assertMapContent(t, dictionary, word, newDefinition)
	})

	t.Run("new word", func(t *testing.T) {
		word := "test"
		definition := "this is just a test"
		dictionary := Dictionary{}
		err := dictionary.Update(word, definition)
		assertError(t, err, ErrWordNotExist)
	})
}

func TestDelete(t *testing.T) {
	t.Run("existing word", func(t *testing.T) {
		word := "test"
		definition := "this is just a test"
		dictionary := Dictionary{word: definition}
		dictionary.Delete(word)
		assertMapNoContent(t, dictionary, word)
		assertMapEmpty(t, dictionary)
	})

	t.Run("word does not exist", func(t *testing.T) {
		word := "test"
		dictionary := Dictionary{}
		dictionary.Delete(word)
		assertMapNoContent(t, dictionary, word)
		assertMapEmpty(t, dictionary)
	})
}

func assertMapStrings(t *testing.T, got, want, given string) {
	t.Helper()
	if got != want {
		t.Errorf("got '%s' want '%s' given '%s'", got, want, given)
	}
}

func assertError(t *testing.T, got, want error) {
	t.Helper()
	if got != want {
		t.Errorf("got error '%s' want '%s'", got, want)
	}
}

func assertNoError(t *testing.T, got error) {
	t.Helper()
	if got != nil {
		t.Fatal("got an error but didn't want one")
	}
}

func assertMapContent(t *testing.T, dictionary Dictionary, key, value string) {
	t.Helper()
	got, err := dictionary.Search(key)
	if err != nil {
		t.Fatal("should find added word:", key)
	}

	assertMapStrings(t, got, value, key)
}

func assertMapNoContent(t *testing.T, dictionary Dictionary, key string) {
	t.Helper()
	got, err := dictionary.Search(key)
	if err == nil {
		t.Errorf("should not find value: '%s' by key: '%s' but got", got, key)
	}
}

func assertMapEmpty(t *testing.T, dictionary Dictionary) {
	t.Helper()
	size := len(dictionary)
	if size != 0 {
		t.Error("the dictionary should be empty but not")
	}
}
