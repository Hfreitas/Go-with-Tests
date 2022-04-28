package main

import "testing"

func TestHello(t *testing.T) {
	/* assign variable to anonymous function which receives
	a testing.TB (test and benchmark interface), a got and a want parmeters */
	assertCorrectMessage := func (t testing.TB, got, want string)  {
		// declares the function as an test helper
		t.Helper()
		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	}

	t.Run("saying a personal hello to people in English", func(t *testing.T) {
		got := Hello("Cris", "English")
		want := "Hello, Cris"

		assertCorrectMessage(t, got, want)
	})

	t.Run("when an empty string is passed as name parameter and language parameter is English", func(t *testing.T) {
		got := Hello("", "English")
		want := "Hello, World"

		assertCorrectMessage(t, got, want)
	})

	t.Run("in Brazilian Portuguese", func(t *testing.T) {
		got := Hello("Hebert", "Brazilian Portuguese")
		want := "E aí, Hebert"

		assertCorrectMessage(t, got, want)
	})

	t.Run("when an empty string is passed as name parameter and language parameter is Brazilian Portuguese", func(t *testing.T) {
		got := Hello("", "Brazilian Portuguese")
		want := "E aí, man"

		assertCorrectMessage(t, got, want)
	})

	t.Run("in Yoruba", func(t *testing.T) {
		got := Hello("Akin", "Yoruba")
		want := "Pẹlẹ o, Akin"

		assertCorrectMessage(t, got, want)
	})

	t.Run("when an empty string is passed as name parameter and language parameter is Yoruba", func(t *testing.T) {
		got := Hello("", "Yoruba")
		want := "Pẹlẹ o, Aye"

		assertCorrectMessage(t, got, want)
	})

	t.Run("in any other language, defaults to English message", func(t *testing.T) {
		got := Hello("Akin", "Japanese")
		want := "Hello, Akin"

		assertCorrectMessage(t, got, want)
	})

	t.Run("when an empty string is passed as name parameter and language parameter is any other language defaults to English message", func(t *testing.T) {
		got := Hello("", "Japanese")
		want := "Hello, World"

		assertCorrectMessage(t, got, want)
	})

	t.Run("when an empty string is passed as name and language parameters defaults to English message", func(t *testing.T) {
		got := Hello("", "")
		want := "Hello, World"

		assertCorrectMessage(t, got, want)
	})
}
