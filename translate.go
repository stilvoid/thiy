package thiy

import (
	"fmt"

	"gopkg.in/yaml.v2"
)

type Node interface {
	String() string
}

type TextNode struct {
	Content string
}

func (n TextNode) String() string {
	return n.Content
}

func translateItem(in yaml.MapItem) (Node, error) {
	key, ok := in.Key.(string)

	if !ok {
		return nil, fmt.Errorf("Badly formatted tag")
	}

	el, err := NewElement(key)

	if err != nil {
		return nil, err
	}

	switch v := in.Value.(type) {
	case string:
		el.Content = []Node{
			TextNode{v},
		}
	case yaml.MapItem:
		child, err := translateItem(v)

		if err != nil {
			return nil, err
		}

		el.Content = []Node{child}
	case yaml.MapSlice:
		el.Content = make([]Node, 0)

		for _, item := range v {
			child, err := translateItem(item)

			if err != nil {
				return nil, err
			}

			el.Content = append(el.Content, child)
		}
	default:
		return nil, fmt.Errorf("Badly formatted content")
	}

	return el, nil
}
