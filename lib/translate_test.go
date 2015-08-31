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
			yaml.MapItem{"strong", "big"},
			"world",
		}},
	}

	expecteds := []Node{
		Element{
			Tag:     "h1",
			Content: []Node{TextNode{"Hello, world"}},
		},
		Element{
			Tag:     "div",
			Id:      "content",
			Content: []Node{TextNode{"some content"}},
		},
		Element{
			Tag: "div",
			Content: []Node{
				Element{
					Tag:     "p",
					Content: []Node{TextNode{"Hello?"}},
				},
			},
		},
		Element{
			Tag: "div",
			Content: []Node{
				Element{
					Tag:     "p",
					Classes: []string{"win"},
					Content: []Node{TextNode{"First"}},
				},
				Element{
					Tag:     "p",
					Classes: []string{"lose"},
					Content: []Node{TextNode{"Second"}},
				},
			},
		},
		Element{
			Tag: "p",
			Content: []Node{
				TextNode{"Hello"},
				Element{
					Tag:     "strong",
					Content: []Node{TextNode{"big"}},
				},
				TextNode{"world"},
			},
		},
	}

	for i, input := range inputs {
		actual, err := TranslateItem(input)

		if err != nil {
			t.Errorf("unexpected error: %v", err)
		}

		expected := expecteds[i]

		if !reflect.DeepEqual(actual, expected) {
			t.Errorf("unexpected:\n%#v\n\nwant:\n%#v\n", actual, expected)
		}
	}
}
