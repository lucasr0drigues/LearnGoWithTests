package main

import "testing"

func TestHello(t *testing.T) {
	t.Run("saying hello to people", func(t *testing.T){
		got := Hello("Lucas","")
		want := "Hello, Lucas"

		assertCorrectMessage(t,got,want)
	})

	t.Run("say 'Hello, World' when an empty string is supplied", func(t *testing.T){
		got := Hello("","")
		want := "Hello, World"

		assertCorrectMessage(t,got,want)
	})

	t.Run("in Spanish", func(t *testing.T){
		got := Hello("Elodie", "Spanish")
		want := "Hola, Elodie"
		assertCorrectMessage(t,got,want)
	})

	t.Run("in French", func(t *testing.T){
		got := Hello("Ana","French")
		want := "Bonjour, Ana"
		assertCorrectMessage(t,got,want)
	})

	t.Run("in Portuguese", func(t *testing.T){
		got := Hello("Lucas", "Portuguese")
		want := "Bom dia, Lucas"
		assertCorrectMessage(t,got,want)
	})
}

func assertCorrectMessage(t testing.TB, got, want string){
	t.Helper()
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}