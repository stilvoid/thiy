package lib

import (
	"strings"
	"testing"
)

func TestParse(t *testing.T) {
	cases := map[string]string{
		"p: hello":                            "<p>hello</p>",
		"p:\n- hello\n- strong: big\n- world": "<p>hello <strong>big</strong> world</p>",
	}

	for input, expected := range cases {
		actual, err := Parse(strings.NewReader(input))

		if err != nil {
			t.Errorf("unexpected error: %v", err)
		}

		if actual != expected {
			t.Errorf("unexpected '%v', want '%v'", actual, expected)
		}
	}
}
