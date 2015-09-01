package dialect

import "offend.me.uk/thiy/common"

func translateNodes(in []common.Node) []common.Node {
	var out []common.Node

	for _, node := range in {
		switch v := node.(type) {
		case common.TagNode:
			node = Bootstrap(v)
		}

		out = append(out, node)
	}

	return out
}

func makePanel(in common.TagNode) (out common.TagNode) {
	if in.Tag != "panel" {
		panic("attempt to make a bootstrap panel out of a non-panel")
	}

	out.Tag = "div"

	out.Classes = []string{"panel"}

	if len(in.Classes) == 0 {
		out.Classes = append(out.Classes, "panel-default")
	} else {
		for _, class := range in.Classes {
			out.Classes = append(out.Classes, "panel-"+class)
		}
	}

	for _, attribute := range in.Attributes {
		if attribute.Name == "title" {
			out.Content = append(out.Content, common.TagNode{
				Tag:     "div",
				Classes: []string{"panel-heading", "panel-title"},
				Content: []common.Node{
					common.TextNode{attribute.Value},
				},
			})
		}
	}

	out.Content = append(out.Content, common.TagNode{
		Tag:     "div",
		Classes: []string{"panel-body"},
		Content: in.Content,
	})

	return
}

func Bootstrap(in common.TagNode) common.TagNode {
	in.Content = translateNodes(in.Content)

	switch in.Tag {
	case "panel":
		return makePanel(in)
	default:
		return in
	}
}
