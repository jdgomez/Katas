package main

import "testing"

func TestHello(t *testing.T) {
	t.Run("saying hello to people", func(t *testing.T) {
		got := Hello("Chris", "")
		want := "Hello, Chris"
		assertCorrectMessage(t, got, want)
	})

	t.Run("saying hello to people in Spanish", func(t *testing.T) {
		got := Hello("Cristina", "Spanish")
		want := "Hola, Cristina"
		assertCorrectMessage(t, got, want)
	})

	t.Run("saying hello to people in French", func(t *testing.T) {
		got := Hello("Emily", "French")
		want := "Bonjour, Emily"
		assertCorrectMessage(t, got, want)
	})

	t.Run("saying hello to people in Deutsch", func(t *testing.T) {
		got := Hello("Lily", "Deutsch")
		want := "Halo, Lily"
		assertCorrectMessage(t, got, want)
	})

	t.Run("say 'Hello, World' when an empty string is supplied", func(t *testing.T) {
		got := Hello("", "")
		want := "Hello, World"
		assertCorrectMessage(t, got, want)
	})
}

func assertCorrectMessage(t *testing.T, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
