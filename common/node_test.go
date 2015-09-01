package common

import (
	"reflect"
	"testing"
)

func TestOutput(t *testing.T) {
	inputs := []Node{
		TagNode{
			Tag: "div",
		},
		TagNode{
			Tag: "div",
			Id:  "content",
		},
		TagNode{
			Tag:     "div",
			Classes: []string{"red", "blue"},
		},
		TagNode{
			Tag:     "div",
			Id:      "content",
			Classes: []string{"red", "blue"},
		},
		TagNode{
			Tag:     "div",
			Id:      "content",
			Classes: []string{"red", "blue"},
			Attributes: []Attribute{
				{"href", "#"},
				{"name", "link"},
			},
		},
		TextNode{"Hello, world"},
		TagNode{
			Tag:     "p",
			Content: []Node{TextNode{"Hello, world"}},
		},
		TagNode{
			Tag: "div",
			Content: []Node{
				TagNode{
					Tag:     "p",
					Content: []Node{TextNode{"Hello, world"}},
				},
			},
		},
		TagNode{
			Tag: "p",
			Content: []Node{
				TextNode{"Goodbye"},
				TagNode{
					Tag:     "em",
					Content: []Node{TextNode{"cruel"}},
				},
				TextNode{"world"},
			},
		},
		TagNode{
			Tag: "p",
			Attributes: []Attribute{
				{"empty", ""},
			},
			Content: []Node{
				TextNode{"Hello, world"},
			},
		},
	}

	expecteds := []string{
		"<div />",
		"<div id=\"content\" />",
		"<div class=\"red blue\" />",
		"<div id=\"content\" class=\"red blue\" />",
		"<div id=\"content\" class=\"red blue\" href=\"#\" name=\"link\" />",
		"Hello, world",
		"<p>Hello, world</p>",
		"<div><p>Hello, world</p></div>",
		"<p>Goodbye <em>cruel</em> world</p>",
		"<p empty>Hello, world</p>",
	}

	for i, input := range inputs {
		actual := input.String()

		expected := expecteds[i]

		if !reflect.DeepEqual(actual, expected) {
			t.Errorf("unexpected '%v', want '%v'", actual, expected)
		}
	}
}
