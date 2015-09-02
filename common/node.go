package common

import (
	"bytes"
	"fmt"
	"strings"
)

var voidElements = []string{"area", "base", "br", "col", "embed", "hr", "img", "input", "keygen", "link", "meta", "param", "source", "track", "wbr"}

type Node interface {
	String() string
	render(int) string
}

type TextNode struct {
	Content string
}

type Attribute struct {
	Name  string
	Value string
}

type TagNode struct {
	Tag        string
	Id         string
	Classes    []string
	Attributes []Attribute
	Content    []Node
}

const INDENT = "    "

func (n TextNode) render(indent int) string {
	var buf bytes.Buffer

	indentString := strings.Repeat(INDENT, indent)

	buf.WriteString(indentString)

	buf.WriteString(strings.Replace(n.Content, "\n", "\n"+indentString, -1))

	return buf.String()
}

func (n TextNode) String() string {
	return n.render(0)
}

func (n TagNode) render(indent int) string {
	var buf bytes.Buffer

	indentString := strings.Repeat(INDENT, indent)

	buf.WriteString(indentString)

	buf.WriteString(fmt.Sprintf("<%s", n.Tag))

	if n.Id != "" {
		buf.WriteString(fmt.Sprintf(" id=\"%s\"", n.Id))
	}

	if n.Classes != nil || len(n.Classes) > 0 {
		buf.WriteString(fmt.Sprintf(" class=\"%s\"", strings.Join(n.Classes, " ")))
	}

	if n.Attributes != nil || len(n.Attributes) > 0 {
		for _, attr := range n.Attributes {
			if attr.Value == "" {
				buf.WriteString(fmt.Sprintf(" %s", attr.Name))
			} else {
				buf.WriteString(fmt.Sprintf(" %s=\"%s\"", attr.Name, attr.Value))
			}
		}
	}

	if n.Content == nil || len(n.Content) == 0 {
		for _, tag := range voidElements {
			if n.Tag == tag {
				buf.WriteString(" />")
				return buf.String()
			}
		}

		buf.WriteString("></" + n.Tag + ">")
		return buf.String()
	}

	buf.WriteString(">")

	if len(n.Content) == 1 {
		switch node := n.Content[0].(type) {
		case TagNode:
			buf.WriteString("\n")
			buf.WriteString(node.render(indent + 1))
			buf.WriteString("\n")

			buf.WriteString(indentString)
		case TextNode:
			if strings.ContainsRune(node.Content, '\n') {
				buf.WriteString("\n")
				buf.WriteString(node.render(indent + 1))
				buf.WriteString("\n")

				buf.WriteString(indentString)
			} else {
				buf.WriteString(node.String())
			}
		}
	} else {
		for _, child := range n.Content {
			buf.WriteString("\n")
			buf.WriteString(child.render(indent + 1))
		}

		buf.WriteString("\n")
		buf.WriteString(indentString)
	}

	buf.WriteString(fmt.Sprintf("</%s>", n.Tag))

	return buf.String()
}

func (n TagNode) String() string {
	return n.render(0)
}
