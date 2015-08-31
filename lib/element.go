package thiy

import (
	"bytes"
	"fmt"
	"regexp"
	"strings"
)

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

const spaceReText = `\s+`
const attributeReText = `^(?P<name>\w+)(?:=(?:(?P<value>\w+)|\"(?P<qvalue>[^"]*)\"))?`
const elementReText = `^(?P<tag>\w+)(?:#(?P<id>\w+))?(?:\.(?P<classes>.*?))?(?:\((?P<attributes>.*?)\))?$`

var spaceRe *regexp.Regexp
var attributeRe *regexp.Regexp
var elementRe *regexp.Regexp

var attributeNames []string
var elementNames []string

func init() {
	spaceRe = regexp.MustCompile(spaceReText)
	attributeRe = regexp.MustCompile(attributeReText)
	elementRe = regexp.MustCompile(elementReText)

	attributeNames = attributeRe.SubexpNames()
	elementNames = elementRe.SubexpNames()
}

func parseAttributes(input string) ([]Attribute, error) {
	attributes := make([]Attribute, 0)

	for len(input) > 0 {
		match := attributeRe.FindStringSubmatch(input)

		if match == nil {
			return nil, fmt.Errorf("Badly-formatted thingy")
		}

		var attribute Attribute

		for i, name := range attributeNames {
			switch name {
			case "name":
				attribute.Name = match[i]
			case "value", "qvalue":
				if match[i] != "" {
					attribute.Value = match[i]
				}
			}
		}

		attributes = append(attributes, attribute)

		input = input[len(match[0]):]

		input = strings.TrimSpace(input)
	}

	return attributes, nil
}

func NewElement(input string) (Element, error) {
	var el Element

	match := elementRe.FindStringSubmatch(input)

	if match == nil {
		return el, fmt.Errorf("Badly-formatted wossname")
	}

	for i, name := range elementNames {
		switch name {
		case "tag":
			el.Tag = match[i]
		case "id":
			el.Id = match[i]
		case "classes":
			if match[i] == "" {
				continue
			}

			el.Classes = strings.Fields(match[i])
		case "attributes":
			if match[i] == "" {
				continue
			}

			var err error

			el.Attributes, err = parseAttributes(match[i])

			if err != nil {
				return el, err
			}
		}

	}

	return el, nil
}
