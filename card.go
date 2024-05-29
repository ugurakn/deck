//go:generate stringer -type=Suit,Rank

package deck

import (
	"fmt"
)

const deckSize = 52

// Card suits (in default order) Spade, Diamond, Club, Heart.
type Suit uint8

const (
	Spade Suit = iota
	Diamond
	Club
	Heart
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
	return fmt.Sprintf("%s of %ss", c.Rank, c.Suit)
}

type SortByRank []Card

func (a SortByRank) Len() int      { return len(a) }
func (a SortByRank) Swap(i, j int) { a[i], a[j] = a[j], a[i] }
func (a SortByRank) Less(i, j int) bool {
	if a[i].Suit == a[j].Suit {
		return a[i].Rank < a[j].Rank
	}
	return a[i].Suit < a[j].Suit
}

func New() []Card {
	deck := make([]Card, deckSize)

	var i int
	for _, s := range suits {
		for r := minRank; r <= maxRank; r++ {
			deck[i] = Card{Suit: s, Rank: r}
			i++
		}
	}

	return deck
}
