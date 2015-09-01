package lib

import (
	"fmt"

	"offend.me.uk/thiy/common"

	"gopkg.in/yaml.v2"
)

func translateValue(in interface{}) ([]common.Node, error) {
	switch v := in.(type) {
	case nil:
		return []common.Node{}, nil
	case string:
		return []common.Node{common.TextNode{v}}, nil
	case yaml.MapItem:
		child, err := translateItem(v)

		return []common.Node{child}, err
	case yaml.MapSlice:
		out := make([]common.Node, len(v))

		for i, item := range v {
			child, err := translateItem(item)

			if err != nil {
				return nil, err
			}

			out[i] = child
		}

		return out, nil
	case []interface{}:
		out := make([]common.Node, 0, len(v))

		for _, item := range v {
			child, err := translateValue(item)

			if err != nil {
				return nil, err
			}

			out = append(out, child...)
		}

		return out, nil
	}

	return nil, fmt.Errorf("Badly formatted content: %#v", in)
}

func translateItem(in yaml.MapItem) (common.Node, error) {
	key, ok := in.Key.(string)

	if !ok {
		return nil, fmt.Errorf("Badly formatted tag")
	}

	el, err := newElement(key)

	if err != nil {
		return nil, err
	}

	el.Content, err = translateValue(in.Value)

	if err != nil {
		return nil, err
	}

	return el, nil
}
