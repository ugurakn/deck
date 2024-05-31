package deck

import (
	"math/rand"
	"slices"
	"testing"
)

func TestCardStringer(t *testing.T) {
	testCases := []struct {
		c      Card
		expect string
	}{
		{Card{Suit: Heart, Rank: Ace}, "Ace of Hearts"},
		{Card{Suit: Spade, Rank: Nine}, "Nine of Spades"},
		{Card{Suit: Diamond, Rank: K}, "K of Diamonds"},
	}

	for _, tc := range testCases {
		t.Run(tc.expect, func(t *testing.T) {
			if r := tc.c.String(); r != tc.expect {
				t.Fatalf("expected: \"%v\", got: \"%v\"", tc.expect, r)
			}
		})
	}
}

func TestShuffle(t *testing.T) {
	shuffleRand = rand.New(rand.NewSource(0))
	d1 := New()
	d1 = Shuffle(d1)

	shuffleRand = rand.New(rand.NewSource(0))
	d2 := New()
	d2 = Shuffle(d2)

	for i, c1 := range d1 {
		if c2 := d2[i]; c1 != c2 {
			t.Fatalf(
				"decks shuffled with the same seed are different (index %v): expected %v, got %v",
				i, c1, c2,
			)
		}
	}
}

func TestJokers(t *testing.T) {
	numOfJokers := 4
	d := New(WithJokers(numOfJokers))

	count := 0
	for _, c := range d {
		if c.Suit == Joker {
			count++
		}
	}
	if count != numOfJokers {
		t.Fatalf("expected %v jokers, got %v", numOfJokers, count)
	}
}

func TestDefaultSort(t *testing.T) {
	// default order => S, D, C, H. ranks asc
	testCases := []struct {
		i      int
		expect Card
	}{
		{0, Card{Suit: Spade, Rank: Ace}},
		{1, Card{Suit: Spade, Rank: Two}},
		{12, Card{Suit: Spade, Rank: K}},
		{13, Card{Suit: Diamond, Rank: Ace}},
		{25, Card{Suit: Diamond, Rank: K}},
		{26, Card{Suit: Club, Rank: Ace}},
		{38, Card{Suit: Club, Rank: K}},
		{39, Card{Suit: Heart, Rank: Ace}},
		{51, Card{Suit: Heart, Rank: K}},
	}

	// shuffle, then resort as default
	d := New(Shuffle)
	d = DefaultSort(d)

	for _, tc := range testCases {
		if got := d[tc.i]; got != tc.expect {
			t.Fatalf("expected %v at index [%v], got %v.", tc.expect, tc.i, got)
		}
	}
}

func TestWithFilter(t *testing.T) {
	filt234 := func(c Card) bool {
		if c.Rank == 2 || c.Rank == 3 || c.Rank == 4 {
			return false
		}
		return true
	}

	filtSpade := func(c Card) bool {
		return c.Suit != Spade
	}

	filtD_Ace := func(c Card) bool {
		return c.Suit != Diamond && c.Rank != Ace
	}

	testCases := []struct {
		fRank []Rank
		fSuit []Suit
		fFn   func(c Card) bool
		name  string
	}{
		{
			fRank: []Rank{Two, Three, Four},
			fSuit: []Suit{},
			fFn:   filt234,
			name:  "filter out ranks 2, 3, 4",
		},
		{
			fRank: []Rank{},
			fSuit: []Suit{Spade},
			fFn:   filtSpade,
			name:  "filter out spades",
		},
		{
			fRank: []Rank{Ace},
			fSuit: []Suit{Diamond},
			fFn:   filtD_Ace,
			name:  "filter out diamonds and aces",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			d := New(WithFilter(tc.fFn))
			for _, c := range d {
				if slices.Contains(tc.fRank, c.Rank) || slices.Contains(tc.fSuit, c.Suit) {
					t.Fatalf("expected rank(s) %v & suit(s) %v filtered out, found %v", tc.fRank, tc.fSuit, c)
				}
			}
		})
	}
}

func TestWithExtraDecks(t *testing.T) {
	testCases := []struct {
		deck      []Card
		expectLen int
		name      string
	}{
		{New(WithExtraDecks(0)), deckSize, "extra_decks_0"},
		{New(WithExtraDecks(1)), deckSize * 2, "extra_decks_1"},
		{New(WithExtraDecks(2)), deckSize * 3, "extra_decks_2"},
		{New(WithJokers(2), WithExtraDecks(1)), (deckSize + 2) * 2, "jokers_2_extra_decks_1"},
		{New(WithExtraDecks(1), WithJokers(4)), deckSize*2 + 4, "extra_decks_1_jokers_4"},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			if len(tc.deck) != tc.expectLen {
				t.Errorf("expected deck length to be %v, got %v", len(tc.deck), tc.expectLen)
			}
		})
	}
}
