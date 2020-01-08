package main

import "testing"

func TestHello(t *testing.T)  {
	resultado := Hello("Lucas")
	esperado := "Hello World Lucas"

	if resultado != esperado {
		t.Errorf("resultado '%s', esperado '%s'", resultado, esperado)
	}
}