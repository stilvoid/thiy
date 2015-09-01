package dialect

import (
	"reflect"
	"testing"

	"offend.me.uk/thiy/common"
)

func TestBootstrap(t *testing.T) {
	inputs := []common.TagNode{
		{
			Tag: "panel",
			Content: []common.Node{
				common.TextNode{"Hello"},
			},
		},
		{
			Tag:     "panel",
			Classes: []string{"primary"},
			Content: []common.Node{
				common.TextNode{"Hello"},
			},
		},
		{
			Tag: "panel",
			Attributes: []common.Attribute{
				{"title", "My panel title"},
			},
			Content: []common.Node{
				common.TextNode{"Hello"},
			},
		},
		{
			Tag:     "div",
			Classes: []string{"not-bootstrap"},
			Content: []common.Node{
				common.TagNode{
					Tag: "panel",
				},
			},
		},
	}

	expecteds := []common.TagNode{
		{
			Tag:     "div",
			Classes: []string{"panel", "panel-default"},
			Content: []common.Node{
				common.TagNode{
					Tag:     "div",
					Classes: []string{"panel-body"},
					Content: []common.Node{
						common.TextNode{"Hello"},
					},
				},
			},
		},
		{
			Tag:     "div",
			Classes: []string{"panel", "panel-primary"},
			Content: []common.Node{
				common.TagNode{
					Tag:     "div",
					Classes: []string{"panel-body"},
					Content: []common.Node{
						common.TextNode{"Hello"},
					},
				},
			},
		},
		{
			Tag:     "div",
			Classes: []string{"panel", "panel-default"},
			Content: []common.Node{
				common.TagNode{
					Tag:     "div",
					Classes: []string{"panel-heading", "panel-title"},
					Content: []common.Node{
						common.TextNode{"My panel title"},
					},
				},
				common.TagNode{
					Tag:     "div",
					Classes: []string{"panel-body"},
					Content: []common.Node{
						common.TextNode{"Hello"},
					},
				},
			},
		},
		{
			Tag:     "div",
			Classes: []string{"not-bootstrap"},
			Content: []common.Node{
				common.TagNode{
					Tag:     "div",
					Classes: []string{"panel", "panel-default"},
					Content: []common.Node{
						common.TagNode{
							Tag:     "div",
							Classes: []string{"panel-body"},
						},
					},
				},
			},
		},
	}

	for i, input := range inputs {
		actual := Bootstrap(input)

		expected := expecteds[i]

		if !reflect.DeepEqual(actual, expected) {
			t.Errorf("unexpected '%#v', want '%#v'", actual, expected)
		}
	}
}
