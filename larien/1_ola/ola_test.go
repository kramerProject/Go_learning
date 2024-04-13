package main

import "testing"

func TestOla(t *testing.T) {

	verifyRightMessage := func(t *testing.T, result, expected string) {
		t.Helper()
		if result != expected {
			t.Errorf("resultado '%s', esperado '%s'", result, expected)
		}
	}

	t.Run("say hello to people", func(t *testing.T) {
		result := Ola("Kramer", "")
		expected := "Olá, Kramer."
		verifyRightMessage(t, result, expected)
	})

	t.Run("say hello world when empty string", func(t *testing.T) {
		result := Ola("", "spanish")
		expected := "Hola, Mundo."
		verifyRightMessage(t, result, expected)
	})

	t.Run("say hello in spanish", func(t *testing.T) {
		result := Ola("Elodie", "spanish")
		expected := "Hola, Elodie."
		verifyRightMessage(t, result, expected)
	})

	t.Run("say hello in french", func(t *testing.T) {
		result := Ola("Elodie", "french")
		expected := "Bonjour, Elodie."
		verifyRightMessage(t, result, expected)
	})

}
