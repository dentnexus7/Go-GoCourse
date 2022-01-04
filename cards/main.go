package main

func main() {
	// var card string = "Ace of Spades"

	// Go will infer the type
	// Only use := when defining a new variable
	//card := "Ace of Spades"

	// type deck ==== []string
	//cards := deck{"Ace of Spades", newCard()}
	//cards = append(cards, "Six of Spades")

	// Create a new deck example
	//cards := newDeck()
	// Separate deck into two slices
	//hand, remainingCards := deal(cards, 5)
	//hand.print()
	//remainingCards.print()

	// Type conversion example
	//greeting := "Hi There"
	//fmt.Println([]byte(greeting))

	// Write to File example
	//cards := newDeck()
	//cards.saveToFile("my_cards")

	// Read from File Example
	//cards := newDeckFromFile("my_cards")
	//cards.print()

	// Shuffle Deck example
	cards := newDeck()
	cards.shuffle()
	cards.print()
}
