package lib

import (
	"io"
	"io/ioutil"
	"strings"

	"gopkg.in/yaml.v2"
	"offend.me.uk/thiy/common"
	"offend.me.uk/thiy/dialect"
)

func Parse(r io.Reader, dialectName string, includeWrapper bool) (string, error) {
	input, err := ioutil.ReadAll(r)
	if err != nil {
		return "", err
	}

	var parsed yaml.MapSlice

	err = yaml.Unmarshal(input, &parsed)
	if err != nil {
		return "", err
	}

	var out []string

	if includeWrapper {
		parsed = yaml.MapSlice{
			{"html", parsed},
		}

		out = []string{"<!DOCTYPE html>"}
	}

	for _, node := range parsed {
		el, err := translateItem(node)
		if err != nil {
			return "", err
		}

		switch dialectName {
		case "html":
			el = dialect.HTML(el.(common.TagNode))
		case "bootstrap":
			el = dialect.HTML(el.(common.TagNode))
			el = dialect.Bootstrap(el.(common.TagNode))
		}

		out = append(out, el.String())
	}

	return strings.Join(out, "\n"), nil
}
