package actions

import (
	"bytes"
	"testing"
)

func TestFlipCoin(t *testing.T) {
	n := 0
	heads := 0
	tails := 0

	for n < 200 {
		n++
		buf := bytes.Buffer{}
		err := flipCoin(nil, &buf)

		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}

		results := []string{"heads\n", "tails\n"}

		switch got := buf.String(); got {
		case "heads\n":
			heads++
		case "tails\n":
			tails++
		default:
			t.Errorf("got %q, wanted any of %q", got, results)
		}
	}

	if heads > 150 || tails > 150 {
		t.Errorf("coin looks biased: got %d heads and %d tails out of %d flips, want neither side above 150", heads, tails, n)
	}
}
