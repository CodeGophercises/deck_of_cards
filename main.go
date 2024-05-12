package main

import (
	"fmt"

	"github.com/CodeGophercises/deck_of_cards/deck"
)

func main() {

	cardDeck := deck.NewMultiDeck(2, deck.DefaultSort, deck.Shuffle, deck.AddJoker(2), deck.FilterRanks(deck.Ace, deck.Three))
	fmt.Printf("Got deck with %d cards\n", len(cardDeck))
	for _, card := range cardDeck {
		fmt.Println(card.Name())
	}
}
