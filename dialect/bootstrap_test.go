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
		{
			Tag: "input",
			Content: []common.Node{
				common.TextNode{"Some text"},
			},
		},
		{
			Tag: "input",
			Attributes: []common.Attribute{
				{"label", "Field label"},
			},
			Content: []common.Node{
				common.TextNode{"Some text"},
			},
		},
		{
			Tag: "input",
			Id:  "my-field",
			Attributes: []common.Attribute{
				{"label", "Field label"},
			},
			Content: []common.Node{
				common.TextNode{"Some text"},
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
		{
			Tag:     "div",
			Classes: []string{"form-group"},
			Content: []common.Node{
				common.TagNode{
					Tag:     "input",
					Classes: []string{"form-control"},
					Attributes: []common.Attribute{
						{"placeholder", "Some text"},
					},
				},
			},
		},
		{
			Tag:     "div",
			Classes: []string{"form-group"},
			Content: []common.Node{
				common.TagNode{
					Tag:     "label",
					Classes: []string{"control-label"},
					Content: []common.Node{
						common.TextNode{"Field label"},
					},
				},
				common.TagNode{
					Tag:     "input",
					Classes: []string{"form-control"},
					Attributes: []common.Attribute{
						{"placeholder", "Some text"},
					},
				},
			},
		},
		{
			Tag:     "div",
			Classes: []string{"form-group"},
			Content: []common.Node{
				common.TagNode{
					Tag: "label",
					Attributes: []common.Attribute{
						{"for", "my-field"},
					},
					Classes: []string{"control-label"},
					Content: []common.Node{
						common.TextNode{"Field label"},
					},
				},
				common.TagNode{
					Tag:     "input",
					Id:      "my-field",
					Classes: []string{"form-control"},
					Attributes: []common.Attribute{
						{"placeholder", "Some text"},
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
