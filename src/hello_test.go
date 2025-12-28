package main

import "testing"

func TestHello(t *testing.T) {
	t.Run("hello greg test", func(test *testing.T) {
		got := HelloWorld("greg", "")
		want := "hello greg"

		assertString(test, got, want)
	})

	t.Run("hello empty string test", func(test *testing.T) {
		got := HelloWorld("", "")
		want := "hello world"

		assertString(test, got, want)
	})

	t.Run("in spanish", func(test *testing.T) {
		got := HelloWorld("elodie", "spanish")
		want := "hola elodie"

		assertString(test, got, want)
	})

	t.Run("in french", func(test *testing.T) {
		got := HelloWorld("baguette", "french")
		want := "bonjour baguette"

		assertString(test, got, want)
	})
}

func assertString(t testing.TB, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
