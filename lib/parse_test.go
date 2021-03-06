package lib

import (
	"strings"
	"testing"
)

func TestParse(t *testing.T) {
	cases := map[string]string{
		"p: hello":                            "<!DOCTYPE html>\n<html>\n    <p>hello</p>\n</html>",
		"p:\n  span: hello":                   "<!DOCTYPE html>\n<html>\n    <p>\n        <span>hello</span>\n    </p>\n</html>",
		"p:\n- hello\n- strong: big\n- world": "<!DOCTYPE html>\n<html>\n    <p>\n        hello\n        <strong>big</strong>\n        world\n    </p>\n</html>",
		"p: hello\np: goodbye":                "<!DOCTYPE html>\n<html>\n    <p>hello</p>\n    <p>goodbye</p>\n</html>",
	}

	for input, expected := range cases {
		actual, err := Parse(strings.NewReader(input), "html", true)

		if err != nil {
			t.Errorf("unexpected error: %v", err)
		}

		if actual != expected {
			t.Errorf("unexpected '%v', want '%v'", actual, expected)
		}
	}
}

func TestParseNoWrap(t *testing.T) {
	cases := map[string]string{
		"p: hello":                            "<p>hello</p>",
		"p:\n  span: hello":                   "<p>\n    <span>hello</span>\n</p>",
		"p:\n- hello\n- strong: big\n- world": "<p>\n    hello\n    <strong>big</strong>\n    world\n</p>",
		"p: hello\np: goodbye":                "<p>hello</p>\n<p>goodbye</p>",
	}

	for input, expected := range cases {
		actual, err := Parse(strings.NewReader(input), "html", false)

		if err != nil {
			t.Errorf("unexpected error: %v", err)
		}

		if actual != expected {
			t.Errorf("unexpected '%v', want '%v'", actual, expected)
		}
	}
}
