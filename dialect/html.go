package dialect

import (
	"strings"

	"offend.me.uk/thiy/common"
)

func translateHTMLNodes(in []common.Node) []common.Node {
	var out []common.Node

	for _, node := range in {
		switch v := node.(type) {
		case common.TagNode:
			node = HTML(v)
		}

		out = append(out, node)
	}

	return out
}

func makeMeta(in common.TagNode) common.TagNode {
	if len(in.Attributes) != 1 {
		return in
	}

	if in.Attributes[0].Value != "" {
		return in
	}

	if len(in.Content) != 1 {
		return in
	}

	textNode, ok := in.Content[0].(common.TextNode)

	if !ok {
		return in
	}

	in.Attributes = []common.Attribute{
		{"name", in.Attributes[0].Name},
		{"content", textNode.Content},
	}

	in.Content = nil

	return in
}

func makeInclude(in common.TagNode) common.TagNode {
	if len(in.Content) != 1 {
		return in
	}

	textNode, ok := in.Content[0].(common.TextNode)

	if !ok {
		return in
	}

	if strings.HasSuffix(textNode.Content, ".js") {
		in.Tag = "script"
		in.Attributes = append(in.Attributes, common.Attribute{"src", textNode.Content})
		in.Content = nil
	} else if strings.HasSuffix(textNode.Content, ".css") {
		in.Tag = "link"
		in.Attributes = append(in.Attributes, []common.Attribute{
			{"rel", "stylesheet"},
			{"href", textNode.Content},
		}...)
		in.Content = nil
	} else if strings.HasSuffix(textNode.Content, ".ico") {
		in.Tag = "link"
		in.Attributes = append(in.Attributes, []common.Attribute{
			{"rel", "shortcut icon"},
			{"href", textNode.Content},
		}...)
		in.Content = nil
	} else if strings.HasSuffix(textNode.Content, ".xml") {
		in.Tag = "link"
		in.Attributes = append(in.Attributes, []common.Attribute{
			{"rel", "alternate"},
			{"type", "application/rss+xml"},
			{"href", textNode.Content},
		}...)
		in.Content = nil
	} else if strings.HasSuffix(textNode.Content, ".atom") {
		in.Tag = "link"
		in.Attributes = append(in.Attributes, []common.Attribute{
			{"rel", "alternate"},
			{"type", "application/atom+xml"},
			{"href", textNode.Content},
		}...)
		in.Content = nil
	}

	return in
}

/*
HTML adds some convenience elements to save typing.

Examples:

	meta(a): b -> meta(name=a value=b)
	import: a -> link(rel=stylesheet href=a) or script(src=a): ""
*/
func HTML(in common.TagNode) common.TagNode {
	in.Content = translateHTMLNodes(in.Content)

	switch in.Tag {
	case "meta":
		return makeMeta(in)
	case "include":
		return makeInclude(in)
	default:
		return in
	}
}
