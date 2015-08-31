package lib

import (
	"bytes"
	"fmt"
	"strings"
)

type Node interface {
	String() string
}

type TextNode struct {
	Content string
}

type Attribute struct {
	Name  string
	Value string
}

type Element struct {
	Tag        string
	Id         string
	Classes    []string
	Attributes []Attribute
	Content    []Node
}

func (n TextNode) String() string {
	return n.Content
}

func (n Element) String() string {
	var buf bytes.Buffer

	buf.WriteString(fmt.Sprintf("<%s", n.Tag))

	if n.Id != "" {
		buf.WriteString(fmt.Sprintf(" id=\"%s\"", n.Id))
	}

	if n.Classes != nil || len(n.Classes) > 0 {
		buf.WriteString(fmt.Sprintf(" class=\"%s\"", strings.Join(n.Classes, " ")))
	}

	if n.Attributes != nil || len(n.Attributes) > 0 {
		for _, attr := range n.Attributes {
			buf.WriteString(fmt.Sprintf(" %s=\"%s\"", attr.Name, attr.Value))
		}
	}

	if n.Content == nil || len(n.Content) == 0 {
		buf.WriteString(" />")
		return buf.String()
	}

	buf.WriteString(">")

	for i, child := range n.Content {
		buf.WriteString(child.String())

		if i < len(n.Content)-1 {
			buf.WriteString(" ")
		}
	}

	buf.WriteString(fmt.Sprintf("</%s>", n.Tag))

	return buf.String()
}
