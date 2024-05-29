//go:generate stringer -type=Suit,Rank

package deck

import (
	"fmt"
	"sort"
)

const deckSize = 52

// Card suits (in default order) Spade, Diamond, Club, Heart.
type Suit uint8

const (
	Spade Suit = iota
	Diamond
	Club
	Heart
	Joker // special
)

// card ranks Ace, Two,...,Ten, J, Q, K.
type Rank uint8

const (
	_ Rank = iota
	Ace
	Two
	Three
	Four
	Five
	Six
	Seven
	Eight
	Nine
	Ten
	J
	Q
	K
)

var suits = [...]Suit{Spade, Diamond, Club, Heart}

// var ranks = [...]Rank{Ace, Two, Three, Four, Five, Six, Seven, Eight, Nine, Ten, J, Q, K}

const (
	minRank = Ace
	maxRank = K
)

type Card struct {
	Suit
	Rank
}

func (c Card) String() string {
	if c.Suit == Joker {
		return "Joker"
	}
	return fmt.Sprintf("%s of %ss", c.Rank, c.Suit)
}

// type SortByRank []Card

// func (a SortByRank) Len() int      { return len(a) }
// func (a SortByRank) Swap(i, j int) { a[i], a[j] = a[j], a[i] }
// func (a SortByRank) Less(i, j int) bool {
// 	if a[i].Suit == a[j].Suit {
// 		return a[i].Rank < a[j].Rank
// 	}
// 	return a[i].Suit < a[j].Suit
// }

func New(options ...func([]Card) []Card) []Card {
	deck := make([]Card, deckSize)

	var i int
	for _, s := range suits {
		for r := minRank; r <= maxRank; r++ {
			deck[i] = Card{Suit: s, Rank: r}
			i++
		}
	}

	for _, opt := range options {
		deck = opt(deck)
	}

	return deck
}

// WithJokers returns a func that adds j-many jokers to deck.
func WithJokers(j int) func([]Card) []Card {
	return func(deck []Card) []Card {
		for i := 0; i < j; i++ {
			deck = append(deck, Card{Suit: Joker, Rank: 0}) // What should be Joker's Rank???
		}
		return deck
	}
}

// DeckSorter is a user-defined function
// that wraps a less func in closure
// whose signature matches that of sort.Interface Less function
type DeckSorter func([]Card) func(i, j int) bool

func WithSort(sorter DeckSorter) func([]Card) []Card {
	return func(deck []Card) []Card {
		lessFn := sorter(deck)
		sort.Slice(deck, lessFn)
		return deck
	}
}
