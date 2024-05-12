package deck

import (
	"math/rand"
	"sort"
)

// Sort sorts the deck in an order determined by the provided comparator function
func Sort(less func(deck []Card) func(i, j int) bool) func([]Card) []Card {
	return func(deck []Card) []Card {
		sort.Slice(deck, less(deck))
		return deck
	}
}

func AddJoker(n int) func([]Card) []Card {
	return func(cards []Card) []Card {
		for i := 0; i < n; i++ {
			cards = append(cards, Card{
				Suit: Joker,
			})
		}
		return cards
	}
}

func Shuffle(cards []Card) []Card {
	rand.Shuffle(len(cards), func(i, j int) {
		cards[i], cards[j] = cards[j], cards[i]
	})
	return cards
}

func DefaultSort(cards []Card) []Card {
	sort.Slice(cards, Less(cards))
	return cards
}

// Removes the cards with the given ranks from the deck.
func FilterRanks(ranks ...Rank) func([]Card) []Card {
	return func(cards []Card) []Card {
		m := make(map[Rank]struct{})
		for _, rank := range ranks {
			m[rank] = struct{}{}
		}
		var cardsNew []Card
		for _, c := range cards {
			if _, ok := m[c.Rank]; !ok {
				cardsNew = append(cardsNew, c)
			}
		}
		return cardsNew
	}
}

func Less(cards []Card) func(i, j int) bool {
	return func(i, j int) bool {
		return absRank(cards[i]) < absRank(cards[j])
	}
}

func absRank(card Card) int {
	return int(card.Suit)*20 + int(card.Rank)
}

func NewMultiDeck(n int, options ...func([]Card) []Card) []Card {
	var multiDeck []Card
	for i := 0; i < n; i++ {
		multiDeck = append(multiDeck, generateDeck()...)
	}
	for _, option := range options {
		multiDeck = option(multiDeck)
	}
	return multiDeck
}

func generateDeck() []Card {
	// Generate deck of cards in the default order
	var deck []Card
	for i := 0; i < 4; i++ {
		suit := Suit(i)
		for r := Ace; r <= King; r++ {
			deck = append(deck, Card{
				Rank: r,
				Suit: suit,
			})
		}
	}
	return deck
}

// Generate a deck of cards
func New(options ...func([]Card) []Card) []Card {
	deck := generateDeck()
	for _, option := range options {
		deck = option(deck)
	}

	return deck
}
