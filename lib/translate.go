package lib

import (
	"fmt"

	"gopkg.in/yaml.v2"
)

func TranslateItem(in yaml.MapItem) (Node, error) {
	key, ok := in.Key.(string)

	if !ok {
		return nil, fmt.Errorf("Badly formatted tag")
	}

	el, err := NewElement(key)

	if err != nil {
		return nil, err
	}

	switch v := in.Value.(type) {
	case nil:
	case string:
		el.Content = []Node{
			TextNode{v},
		}
	case yaml.MapItem:
		child, err := TranslateItem(v)

		if err != nil {
			return nil, err
		}

		el.Content = []Node{child}
	case yaml.MapSlice:
		el.Content = make([]Node, 0)

		for _, item := range v {
			child, err := TranslateItem(item)

			if err != nil {
				return nil, err
			}

			el.Content = append(el.Content, child)
		}
	case []interface{}:
		el.Content = make([]Node, 0)

		for _, item := range v {
			mapItem, ok := item.(yaml.MapItem)
			if !ok {
				el.Content = append(el.Content, TextNode{fmt.Sprint(item)})
				continue
			}

			child, err := TranslateItem(mapItem)

			if err != nil {
				return nil, err
			}

			el.Content = append(el.Content, child)
		}
	default:
		return nil, fmt.Errorf("Badly formatted content: %#v", v)
	}

	return el, nil
}
