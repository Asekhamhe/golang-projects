package main

import "fmt"

func main() {

	// cards := newDeck()

	// // cards.print()

	// hand, remainingDecks := deal(cards, 5)

	// hand.print()
	// remainingDecks.print()

	// Type conversion
	// greeting := "Hi there!"
	// fmt.Println([]byte(greeting))

	// cards := newDeck()
	// fmt.Println(cards.toString())

	// cards.saveToFile("my_cards")
	newDecks := newDeckFromFile("my_cards")

	fmt.Println("******************** Shuffle Card  ***************")
	newDecks.shuffleCard()
	newDecks.print()
	fmt.Println("******************** End of Shuffle Card  ***************")

}
