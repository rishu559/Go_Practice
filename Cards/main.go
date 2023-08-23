package main

var deckSize int

func main() {
	// cards := newDeck()

	cards := newDeckFromFile("DeckOnDist.txt")
	// cards.saveToFile("DeckOnDist.txt")
	// cards = append(cards, newCards())

	// for i,card:=range cards {
	// 	fmt.Println(i,card)
	// }

	// print(cards)
	cards.Shuffle()
	cards.print()
	// fmt.Print(cards);

	// hand , remain := deal(cards,5)

	// hand.print()
	// fmt.Println("---------------")
	// remain.print()
}

