package dialect

import "offend.me.uk/thiy/common"

var bsSizes = []string{"xs", "sm", "md", "lg"}

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

func makeFormControl(in common.TagNode) (out common.TagNode) {
	if in.Tag != "input" && in.Tag != "select" {
		panic("attempt to make a bootstrap from control out of a non-input")
	}

	out.Tag = "div"

	out.Classes = []string{"form-group"}

	in.Classes = append(in.Classes, "form-control")

	var attrs []common.Attribute

	for _, attr := range in.Attributes {
		include := true

		if attr.Name == "label" {
			labelNode := common.TagNode{
				Tag:     "label",
				Classes: []string{"control-label"},
				Content: []common.Node{
					common.TextNode{attr.Value},
				},
			}

			if in.Id != "" {
				labelNode.Attributes = []common.Attribute{
					{"for", in.Id},
				}
			}

			out.Content = append(out.Content, labelNode)

			include = false
		}

		if include {
			attrs = append(attrs, attr)
		}
	}

	in.Attributes = attrs

	if len(in.Content) == 1 {
		if textNode, ok := in.Content[0].(common.TextNode); ok {
			in.Attributes = append(in.Attributes, common.Attribute{"placeholder", textNode.Content})

			in.Content = nil
		}
	}

	out.Content = append(out.Content, in)

	return
}

func makeRow(in common.TagNode) common.TagNode {
	if in.Tag != "row" {
		panic("attempt to make a bootstrap row out of a non-row")
	}

	in.Tag = "div"

	in.Classes = append(in.Classes, "row")

	return in
}

func makeCol(in common.TagNode) common.TagNode {
	if in.Tag != "col" {
		panic("attempt to make a bootstrap col out of a non-col")
	}

	in.Tag = "div"

	var attrs []common.Attribute

	for _, attr := range in.Attributes {
		include := true

		for _, size := range bsSizes {
			if attr.Name == size {
				in.Classes = append(in.Classes, "col-"+attr.Name+"-"+attr.Value)
				include = false
			}
		}

		if include {
			attrs = append(attrs, attr)
		}
	}

	in.Attributes = attrs

	return in
}

func makeIcon(in common.TagNode) common.TagNode {
	if in.Tag != "icon" {
		panic("attempt to make a bootstrap icon out of a non-icon")
	}

	if len(in.Content) != 1 {
		panic("bootstrap icon requires an icon name")
	}

	textNode, ok := in.Content[0].(common.TextNode)

	if !ok {
		panic("bootstrap icon requires an icon name")
	}

	in.Tag = "span"

	in.Classes = append(in.Classes, "glyphicon", "glyphicon-"+textNode.Content)

	in.Content = nil

	return in
}

func Bootstrap(in common.TagNode) common.TagNode {
	in.Content = translateNodes(in.Content)

	switch in.Tag {
	case "panel":
		return makePanel(in)
	case "input", "select":
		return makeFormControl(in)
	case "row":
		return makeRow(in)
	case "col":
		return makeCol(in)
	case "icon":
		return makeIcon(in)
	default:
		return in
	}
}
