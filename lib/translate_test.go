package lib

import (
	"reflect"
	"testing"

	"gopkg.in/yaml.v2"
)

func TestTranslate(t *testing.T) {
	inputs := []yaml.MapItem{
		{"h1", "Hello, world"},
		{"div#content", "some content"},
		{"div", yaml.MapItem{"p", "Hello?"}},
		{"div", yaml.MapSlice{
			{"p.win", "First"},
			{"p.lose", "Second"},
		}},
		{"p", []interface{}{
			"Hello",
			yaml.MapSlice{yaml.MapItem{"strong", "big"}},
			"world",
		}},
	}

	expecteds := []node{
		element{
			Tag:     "h1",
			Content: []node{textNode{"Hello, world"}},
		},
		element{
			Tag:     "div",
			Id:      "content",
			Content: []node{textNode{"some content"}},
		},
		element{
			Tag: "div",
			Content: []node{
				element{
					Tag:     "p",
					Content: []node{textNode{"Hello?"}},
				},
			},
		},
		element{
			Tag: "div",
			Content: []node{
				element{
					Tag:     "p",
					Classes: []string{"win"},
					Content: []node{textNode{"First"}},
				},
				element{
					Tag:     "p",
					Classes: []string{"lose"},
					Content: []node{textNode{"Second"}},
				},
			},
		},
		element{
			Tag: "p",
			Content: []node{
				textNode{"Hello"},
				element{
					Tag:     "strong",
					Content: []node{textNode{"big"}},
				},
				textNode{"world"},
			},
		},
	}

	for i, input := range inputs {
		actual, err := translateItem(input)

		if err != nil {
			t.Errorf("unexpected error: %v", err)
		}

		expected := expecteds[i]

		if !reflect.DeepEqual(actual, expected) {
			t.Errorf("unexpected:\n%#v\n\nwant:\n%#v\n", actual, expected)
		}
	}
}
