package main

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"strings"
	"time"
)

// Create a new type of 'deck'
// which is a slice of strings
type deck []string

// Create a new deck
func newDeck() deck {
	cards := deck{}

	cardSuits := []string{"Spades", "Diamonds", "Hearts", "Clubs"}
	cardValues := []string{"Ace", "Two", "Three", "Four"}

	for _, suit := range cardSuits {
		for _, value := range cardValues {
			cards = append(cards, value+" of "+suit)
		}
	}

	return cards
}

// Print deck
func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}

// Separate deck into two separate slices
func deal(d deck, handSize int) (deck, deck) {
	return d[:handSize], d[handSize:]
}

// Convert a type deck to a string
func (d deck) toString() string {
	return strings.Join([]string(d), ",")
}

// Save a deck to a file on hard drive
func (d deck) saveToFile(filename string) error {
	// 0666 is standard read and write permissions
	return ioutil.WriteFile(filename, []byte(d.toString()), 0666)
}

// Create deck from hard drive file
func newDeckFromFile(filename string) deck {
	// bs is byteslice
	bs, err := ioutil.ReadFile(filename)
	if err != nil {
		// Option #1 - log the error and return a call to newDeck()
		// Option #2 - Log the error and entirely quit the program
		fmt.Println("Error:", err)
		os.Exit(1)
	}

	// s is slice of strings
	s := strings.Split(string(bs), ",")
	return deck(s)
}

// Shuffle deck using random number generator and seed source
func (d deck) shuffle() {
	source := rand.NewSource(time.Now().UnixNano())
	r := rand.New(source)

	for i := range d {
		newPosition := r.Intn(len(d) - 1)

		d[i], d[newPosition] = d[newPosition], d[i]
	}
}
