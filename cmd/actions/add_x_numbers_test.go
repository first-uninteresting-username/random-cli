package actions

import (
	"bytes"
	"errors"
	"io"
	"strconv"
	"strings"
	"testing"
)

func TestAddNums(t *testing.T) {
	cases := []struct {
		name    string
		input   string
		sum     int
		wantErr error
	}{
		{name: "sum three numbers", input: "3\n10\n20\n30\n", sum: 60, wantErr: nil},
		{name: "zero numbers", input: "0\n", sum: 0, wantErr: nil},
		{name: "negative count is positive", input: "-2\n5\n6\n", sum: 11, wantErr: nil},
		{name: "negative numbers are added", input: "4\n1\n7\n-2\n-5\n", sum: 1, wantErr: nil},
		{name: "count is not a number", input: "hii\n", wantErr: strconv.ErrSyntax},
		{name: "a number is not a number", input: "2\n4\nx\n", wantErr: strconv.ErrSyntax},
		{name: "input ends early", input: "3\n1\n", wantErr: io.EOF},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			buf := bytes.Buffer{}
			err := addNums(strings.NewReader(c.input), &buf)
			got := buf.String()

			if !errors.Is(err, c.wantErr) {
				t.Fatalf("err = %v, want %v", err, c.wantErr)
			}

			if c.wantErr != nil {
				if strings.Contains(got, "The sum is") {
					t.Errorf("printed a sum despite error: %q", got)
				}
				return
			}

			want := "The sum is " + strconv.Itoa(c.sum) + "\n"
			if !strings.HasSuffix(got, want) {
				t.Errorf("got %q, want it to end with %q", got, want)
			}
		})
	}
}

func TestAddNumsSuffixes(t *testing.T) {
	const count = 23

	var input strings.Builder
	input.WriteString(strconv.Itoa(count))
	input.WriteString("\n")
	for range count {
		input.WriteString("1\n")
	}

	buf := bytes.Buffer{}
	if err := addNums(strings.NewReader(input.String()), &buf); err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	got := buf.String()

	wantPrompts := []string{
		"Enter your 1st number: ",
		"Enter your 2nd number: ",
		"Enter your 3rd number: ",
		"Enter your 4th number: ",
		"Enter your 11th number: ",
		"Enter your 12th number: ",
		"Enter your 13th number: ",
		"Enter your 21st number: ",
		"Enter your 22nd number: ",
		"Enter your 23rd number: ",
	}
	for _, p := range wantPrompts {
		if !strings.Contains(got, p) {
			t.Errorf("output is missing prompt %q", p)
		}
	}

	if want := "The sum is 23\n"; !strings.HasSuffix(got, want) {
		t.Errorf("got %q, want it to end with %q", got, want)
	}
}
