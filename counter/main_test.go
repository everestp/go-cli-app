package main

import (
	"testing"

	"github.com/go-rod/rod/lib/input"
)

func TestCountWord(t *testing.T) {

	testCases := []struct {
		name  string
		input string
		wants int
	}{
		{
			name: "5 words",

			input: "one two three four five",
			wants: 5,
		},
		{
			name:  "empty input ",
			input: "",
			wants: 0,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := CountWords([]byte(tc.input))
			if result != tc.wants {
				t.Logf("expected : %d got :%d", tc.wants, result)
				t.Fail()
			}
		})

	}
}
