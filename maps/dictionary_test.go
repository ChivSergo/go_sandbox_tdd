package maps

import "testing"

func TestSearch(t *testing.T) {
	dictionary := Dictionary{"test": "This is just a test"}
	t.Run("Known word", func(t *testing.T) {
		got, _ := dictionary.Search("test")
		want := "This is just a test"

		assertingString(t, got, want)

	})
	t.Run("Unkown word", func(t *testing.T) {
		_, err := dictionary.Search("Unkown")
		want := "could not find the word you were looking for"

		if err == nil {
			t.Fatal("expected to get an error.")
		}

		assertingString(t, err.Error(), want)

	})
}
func assertingString(t testing.TB, got, want string) {
	t.Helper()

	if got != want {
		t.Errorf("We got %s but want %s", got, want)
	}
}
