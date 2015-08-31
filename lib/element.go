package lib

import (
	"fmt"
	"regexp"
	"strings"
)

const spaceReText = `\s+`
const attributeReText = `^(?P<name>[\w-]+)(?:=(?:\"(?P<qvalue>[^"]*)\"|(?P<value>\S+)))?`
const elementReText = `^(?P<tag>[\w-]+)(?:#(?P<id>[\w-]+))?(?:\.(?P<classes>.*?))?(?:\((?P<attributes>.*?)\))?$`

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

func parseAttributes(input string) ([]attribute, error) {
	attributes := make([]attribute, 0)

	for len(input) > 0 {
		match := attributeRe.FindStringSubmatch(input)

		if match == nil {
			return nil, fmt.Errorf("Badly-formatted attribute: %v", input)
		}

		var attribute attribute

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

func newElement(input string) (element, error) {
	var el element

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
