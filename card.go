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

// Card ranks Ace, Two,...,Ten, J, Q, K.
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

// minRank and maxRank facilitate looping over Rank consts.
// They are also used for getting
// absolute rank values (for sorting).
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

// New returns a new deck as a slice of cards.
// With no options specified, it will be a standard
// 52-card deck sorted (via DefaultSort)
// Spade, Diamond, Club, Hearts
// with ranks in each suit sorted in ascending order
// (A, 2,...,10, J, Q, K).
// options are called in the order they were passed.
func New(options ...func([]Card) []Card) []Card {
	deck := make([]Card, deckSize)

	var i int
	for _, s := range suits {
		for r := minRank; r <= maxRank; r++ {
			deck[i] = Card{Suit: s, Rank: r}
			i++
		}
	}

	deck = DefaultSort(deck)

	for _, opt := range options {
		deck = opt(deck)
	}

	return deck
}

// getAbsRank returns an absolute rank value for c.
// It is used by DefaultSort.
func getAbsRank(c Card) int {
	return int(c.Suit)*int(maxRank) + int(c.Rank)
}

// WithJokers wraps in closure a func
// (which is called by New)
// that appends j-many jokers to deck.
func WithJokers(j int) func([]Card) []Card {
	return func(deck []Card) []Card {
		for i := 0; i < j; i++ {
			deck = append(deck, Card{Suit: Joker, Rank: 0}) // What should be Joker's Rank???
		}
		return deck
	}
}

// WithExtraDecks appends k-many
// copies of the deck to the deck.
// Panics if k < 0.
func WithExtraDecks(k int) func([]Card) []Card {
	if k < 0 {
		panic("deck.WithExtraDecks: k cannot be negative.")
	}
	return func(d []Card) []Card {
		cpy := make([]Card, len(d))
		copy(cpy, d)

		for i := 0; i < k; i++ {
			d = append(d, cpy...)
		}
		return d
	}
}

// WithFilter returns a function
// (which is called by New) that calls
// filter on each card in the deck.
// The filter func must return false
// for cards that are to be filtered out.
func WithFilter(filter func(Card) bool) func([]Card) []Card {
	return func(d []Card) []Card {
		filteredDeck := make([]Card, 0)
		for _, c := range d {
			if ok := filter(c); ok {
				filteredDeck = append(filteredDeck, c)
			}
		}
		return filteredDeck
	}
}

// DefaultSort sorts cards as described in New.
func DefaultSort(cards []Card) []Card {
	sort.Slice(cards, func(i, j int) bool {
		return getAbsRank(cards[i]) < getAbsRank(cards[j])
	})
	return cards
}

// DeckSorter defines a user-provided function
// that wraps a less func in closure whose signature
// must match that of [sort.Interface.Less].
type DeckSorter func([]Card) func(i, j int) bool

// WithSorter wraps in closure a func
// (which is called by New)
// that sorts the deck using the user-defined
// less func that sorter returns.
func WithSorter(sorter DeckSorter) func([]Card) []Card {
	return func(deck []Card) []Card {
		lessFn := sorter(deck)
		sort.Slice(deck, lessFn)
		return deck
	}
}

// Shuffle shuffles d using [rand.Shuffle].
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
