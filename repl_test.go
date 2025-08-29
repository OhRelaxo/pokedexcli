package main

import (
	"testing"
)

func TestCleanInput(t *testing.T) {
	cases := []struct {
		input    string
		expected []string
	}{
		{
			input:    "  hello  world  ",
			expected: []string{"hello", "world"},
		},
		{
			input:    "",
			expected: []string{},
		},
		{
			input:    "  123piKa, ,Chu2",
			expected: []string{"123pika,", ",chu2"},
		},
		{
			input:    " ",
			expected: []string{},
		},
	}

	for _, c := range cases {
		actual := cleanInput(c.input)
		lenghtActual := len(actual)
		lenghtExpected := len(c.expected)
		if lenghtActual != lenghtExpected {
			t.Errorf("error lenght of output: %v not as expected: %v", lenghtActual, lenghtExpected)
		}
		for i := range actual {
			word := actual[i]
			expectedWord := c.expected[i]
			if word != expectedWord {
				t.Errorf("error word of actual: %v doesnt match expected: %v", word, expectedWord)
			}
		}
	}

}
