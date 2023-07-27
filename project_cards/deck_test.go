package main

import (
	"os"
	"testing"
)

func TestNewDeck(t *testing.T) {
	d := newDeck()

	if len(d) != 40 {
		t.Errorf("Esperamos um deck de tamanho 40, mas temos %v", len(d))
	}

	if d[0] != "As de Espada" {
		t.Errorf("Esperamos a primeira carta As de Espadas, mas temos %v", d[0])
	}

	if d[len(d)-1] != "Dez de Baloes" {
		t.Errorf("Esperamos a ultima carta Dez de Baloes, mas temos %v", d[len(d)-1])
	}
}

func TestSaveToDeckAndNewDeckFromFile(t *testing.T) {
	os.Remove("_decktesting")

	deck := newDeck()
	deck.saveToFile("_decktesting")

	loadedDeck := newDeckFromFile("_decktesting")

	if len(loadedDeck) != 40 {
		t.Errorf("Esperamos um deck de tamanho 40, mas temos %v", len(loadedDeck))
	}

	os.Remove("_decktesting")

}
