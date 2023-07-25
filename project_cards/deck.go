package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

// Criar um tipo de Deck
// Slice de Strings

type deck []string

func newDeck() deck {
	cards := deck{}

	cardSuits := []string{"Espada", "Coracoes", "Paus", "Baloes"}
	cardValues := []string{"As", "Dois", "Tres", "Quatro", "Cinco", "Seis", "Sete", "Oito", "Nove", "Dez"}

	for _, suit := range cardSuits {
		for _, value := range cardValues {
			cards = append(cards, value+" de "+suit)
		}
	}

	return cards
}

func deal(d deck, handsize int) (deck, deck) {
	return d[:handsize], d[handsize:]
}

func (d deck) print() {
	for _, card := range d {
		fmt.Println(card)
	}
}

func (d deck) toString() string {
	return strings.Join([]string(d), ",")

}

func (d deck) saveToFile(filename string) error {
	return ioutil.WriteFile(filename, []byte(d.toString()), 0666)
}

func newDeckFromFile(filename string) deck {
	bs, err := ioutil.ReadFile("my_cards")
	if err != nil {
		// opcao 1 - guardar o erro em log e retornar um deck , newDeck()
		// opcao 2 - guardar o erro em log e sair do programa
		fmt.Println("Error: ", err)
	}
}
