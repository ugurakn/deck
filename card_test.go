package deck

import (
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

// func TestShuffleModifyOrig(t *testing.T) {
// 	d := New()
// 	shf := make([]Card, len(d))
// 	copy(shf, d)
// 	shf = Shuffle(shf)

// 	for i, c := range d {
// 		if c.String() != shf[i].String() {
// 			fmt.Printf("card mismatch at [%v]: \"%v\" \"%v\"\n", i, c.String(), shf[i].String())
// 			t.FailNow()
// 		}
// 	}
// }

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
