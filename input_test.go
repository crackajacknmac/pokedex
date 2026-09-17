package main

import (
"testing"
)

func TestUserInput(t *testing.T) {
    cases := []struct {
	input    string
	expected []string
}{
	{
		input:    "  hello  world  ",
		expected: []string{"hello", "world"},
	},
}

	for _, c := range cases {
	actual := cleanInput(c.input)
	// Check the length of the actual slice
	// if they don't match, use t.Errorf and continue to the next case
	if len(actual) != len(c.expected) {
		// error and continue here
		t.Errorf("Actual string length = %v; want %v", len(actual), len(c.expected))
		continue
	}
	for i := range actual {
		word := actual[i]
		expectedWord := c.expected[i]
		// Check each word in the slice
		// if they don't match, use t.Errorf to print an error message
			if word != expectedWord {
				t.Errorf("Actual word: %v; expected word: %v", word, expectedWord)
		// and fail the test
	    }
	}
    }
}
