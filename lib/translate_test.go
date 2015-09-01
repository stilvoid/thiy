package lib

import (
	"reflect"
	"testing"

	"offend.me.uk/thiy/common"

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

	expecteds := []common.Node{
		common.TagNode{
			Tag:     "h1",
			Content: []common.Node{common.TextNode{"Hello, world"}},
		},
		common.TagNode{
			Tag:     "div",
			Id:      "content",
			Content: []common.Node{common.TextNode{"some content"}},
		},
		common.TagNode{
			Tag: "div",
			Content: []common.Node{
				common.TagNode{
					Tag:     "p",
					Content: []common.Node{common.TextNode{"Hello?"}},
				},
			},
		},
		common.TagNode{
			Tag: "div",
			Content: []common.Node{
				common.TagNode{
					Tag:     "p",
					Classes: []string{"win"},
					Content: []common.Node{common.TextNode{"First"}},
				},
				common.TagNode{
					Tag:     "p",
					Classes: []string{"lose"},
					Content: []common.Node{common.TextNode{"Second"}},
				},
			},
		},
		common.TagNode{
			Tag: "p",
			Content: []common.Node{
				common.TextNode{"Hello"},
				common.TagNode{
					Tag:     "strong",
					Content: []common.Node{common.TextNode{"big"}},
				},
				common.TextNode{"world"},
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
