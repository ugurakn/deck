//go:generate stringer -type=Suit,Rank

package deck

import (
	"fmt"
	"math/rand"
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

// New returns a new deck as a slice of cards.
// With no options specified, it will be a standard
// 52-card deck sorted Spade, Diamond, Club, Hearts
// with ranks in each suit sorted in ascending order
// (A, 2,...,10, J, Q, K).
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

// WithJokers wraps in closure a func that appends j-many jokers to deck.
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
// whose signature must match that of sort.Interface Less function
type DeckSorter func([]Card) func(i, j int) bool

// WithSorter wraps in closure a func that sorts the deck
// using the user-defined less func that sorter returns.
// Use WithSorter to implement a custom sorting that
// WithSortBy can't provide.
func WithSorter(sorter DeckSorter) func([]Card) []Card {
	return func(deck []Card) []Card {
		lessFn := sorter(deck)
		sort.Slice(deck, lessFn)
		return deck
	}
}

// Shuffle shuffles the deck d (or any slice of cards).
func Shuffle(d []Card) []Card {
	rand.Shuffle(len(d), func(i, j int) {
		d[i], d[j] = d[j], d[i]
	})
	return d
}

// ShuffleP shuffles the deck d like Shuffle
// but instead uses rand.Perm.
// func ShuffleP(d []Card, rnd *rand.Rand) []Card {
// 	shfDeck := make([]Card, len(d))

// 	for i, c := range rnd.Perm(len(d)) {
// 		shfDeck[i] = d[c]
// 	}

// 	return shfDeck
// }

// type SortConfig struct {
// 	// suits are sorted in the order they are found in Suits
// 	// e.g. the suit at Suits[0] is sorted before the one at Suits[1]
// 	// if Suits == nil, the default order is used.
// 	Suits []Suit
// 	// sort ranks in descending order if set to true.
// 	// default is sort in ascending order.
// 	RanksDesc bool
// 	// the deck is sorted by rank if set to true. default is sort by suit.
// 	ByRank bool
// }

// WithSortBy allows sorting by a user-defined order.
// Check out SortConfig for details.
// func WithSortBy(sc SortConfig) func([]Card) []Card {
// 	type cOrder struct {
// 		Card
// 		order int
// 	}

// 	setOrder := func(c cOrder) cOrder {
// 		if sc.Suits == nil || len(sc.Suits) != 5 {
// 			sc.Suits = suits[:]
// 			sc.Suits = append(sc.Suits, Joker)
// 		}

// 		switch c.Suit {
// 		case sc.Suits[0]:
// 			c.order = 0
// 		case sc.Suits[1]:
// 			c.order = 1
// 		case sc.Suits[2]:
// 			c.order = 2
// 		case sc.Suits[3]:
// 			c.order = 3
// 		case sc.Suits[4]:
// 			c.order = 4
// 		}
// 		return c
// 	}

// 	lessWrapper := func(d []Card) func(i int, j int) bool {
// 		return func(i int, j int) bool {
// 			ci, cj := cOrder{d[i], 0}, cOrder{d[j], 0}
// 			ci, cj = setOrder(ci), setOrder(cj)

// 			// sort by rank
// 			if sc.ByRank {
// 				if ci.Rank == cj.Rank {
// 					return ci.order < cj.order
// 				}
// 				if sc.RanksDesc {
// 					return ci.Rank > cj.Rank
// 				} else {
// 					return ci.Rank < cj.Rank
// 				}
// 			}

// 			// sort by suit
// 			if ci.Suit == cj.Suit {
// 				if sc.RanksDesc {
// 					return ci.Rank > cj.Rank
// 				} else {
// 					return ci.Rank < cj.Rank
// 				}
// 			}
// 			return ci.order < cj.order
// 		}
// 	}

// 	return func(d []Card) []Card {
// 		sort.Slice(d, lessWrapper(d))
// 		return d
// 	}
// }
