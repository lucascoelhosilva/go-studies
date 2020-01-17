package hello

import "testing"

func TestHello(t *testing.T)  {
	verificaMensagemCorreta := func(t *testing.T, resultado, esperado string) {
		t.Helper()
		if resultado != esperado {
			t.Errorf("resultado '%s', esperado '%s'", resultado, esperado)
		}
	}

	t.Run("diz olá para as pessoas", func(t *testing.T) {
		resultado := Hello("Lucas", "english")
		esperado := "Hello Lucas"
		verificaMensagemCorreta(t, resultado, esperado)
	})

	t.Run("'Mundo' como padrão para 'string' vazia", func(t *testing.T) {
		resultado := Hello("", "english")
		esperado := "Hello World"
		verificaMensagemCorreta(t, resultado, esperado)
	})

	t.Run("em frances", func(t *testing.T) {
		resultado := Hello("Elodie", "frances")
		esperado := "Bonjour Elodie"
		verificaMensagemCorreta(t, resultado, esperado)
	})

	t.Run("diz olá para as pessoas em portugues", func(t *testing.T) {
		resultado := Hello("Lucas", "portugues")
		esperado := "Ola Lucas"
		verificaMensagemCorreta(t, resultado, esperado)
	})
}
