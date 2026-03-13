package main
import "testing"

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
			input:    "hello world",
			expected: []string{"hello", "world"},
		},
		{
			input:    "hello   world   again",
			expected: []string{"hello", "world", "again"},
		},
		{
			input:    "   spaced   out   words   ",
			expected: []string{"spaced", "out", "words"},
		},
		{
			input:    "",
			expected: []string{},
		},
	}

	for _, c := range cases {
		actual := cleanInput(c.input)

		// Check slice length
		if len(actual) != len(c.expected) {
			t.Errorf(
				"input: %q - expected length %d, got %d",
				c.input,
				len(c.expected),
				len(actual),
			)
			continue
		}

		// Check each word
		for i := range actual {
			word := actual[i]
			expectedWord := c.expected[i]

			if word != expectedWord {
				t.Errorf(
					"input: %q - expected word %q at index %d, got %q",
					c.input,
					expectedWord,
					i,
					word,
				)
			}
		}
	}
}