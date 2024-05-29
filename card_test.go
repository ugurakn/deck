package deck

import "testing"

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
