package main

func main() {

	// cards := newDeck()
	// cards.saveToFile("my_cards")

	// cards.print()

	// hand, remain := deal(cards, 5)

	// hand.print()
	// remain.print()
	// fmt.Println(cards.toString())

	cards := newDeckFromFile("my_cards")
	cards.shuffle()
	cards.print()

}
