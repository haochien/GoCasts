package main

func main() {
	cards := newDeck()
	cards.shuffle()
	// cards.print()

	hand, rest := deal(cards, 5)
	hand.print()
	rest.print()

}
