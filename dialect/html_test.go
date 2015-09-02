package dialect

import (
	"reflect"
	"testing"

	"offend.me.uk/thiy/common"
)

func TestHTML(t *testing.T) {
	inputs := []common.TagNode{
		{
			Tag: "meta",
			Attributes: []common.Attribute{
				{"turnip", ""},
			},
			Content: []common.Node{
				common.TextNode{"juice"},
			},
		},
		{
			Tag: "meta",
			Attributes: []common.Attribute{
				{"turnip", "juice"},
			},
		},
		{
			Tag: "include",
			Content: []common.Node{
				common.TextNode{"foo.js"},
			},
		},
		{
			Tag: "include",
			Content: []common.Node{
				common.TextNode{"foo.css"},
			},
		},
		{
			Tag: "include",
			Content: []common.Node{
				common.TextNode{"foo.ico"},
			},
		},
		{
			Tag: "include",
			Content: []common.Node{
				common.TextNode{"foo.xml"},
			},
		},
		{
			Tag: "include",
			Content: []common.Node{
				common.TextNode{"foo.atom"},
			},
		},
	}

	expecteds := []common.TagNode{
		{
			Tag: "meta",
			Attributes: []common.Attribute{
				{"name", "turnip"},
				{"content", "juice"},
			},
		},
		{
			Tag: "meta",
			Attributes: []common.Attribute{
				{"turnip", "juice"},
			},
		},
		{
			Tag: "script",
			Attributes: []common.Attribute{
				{"src", "foo.js"},
			},
		},
		{
			Tag: "link",
			Attributes: []common.Attribute{
				{"rel", "stylesheet"},
				{"href", "foo.css"},
			},
		},
		{
			Tag: "link",
			Attributes: []common.Attribute{
				{"rel", "shortcut icon"},
				{"href", "foo.ico"},
			},
		},
		{
			Tag: "link",
			Attributes: []common.Attribute{
				{"rel", "alternate"},
				{"type", "application/rss+xml"},
				{"href", "foo.xml"},
			},
		},
		{
			Tag: "link",
			Attributes: []common.Attribute{
				{"rel", "alternate"},
				{"type", "application/atom+xml"},
				{"href", "foo.atom"},
			},
		},
	}

	for i, input := range inputs {
		actual := HTML(input)

		expected := expecteds[i]

		if !reflect.DeepEqual(actual, expected) {
			t.Errorf("unexpected '%#v', want '%#v'", actual, expected)
		}
	}
}
