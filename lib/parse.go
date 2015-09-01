package lib

import (
	"io"
	"io/ioutil"
	"strings"

	"gopkg.in/yaml.v2"
	"offend.me.uk/thiy/common"
	"offend.me.uk/thiy/dialect"
)

func Parse(r io.Reader, dialectName string) (string, error) {
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

	for _, node := range parsed {
		el, err := translateItem(node)
		if err != nil {
			return "", err
		}

		if dialectName == "bootstrap" {
			el = dialect.Bootstrap(el.(common.TagNode))
		}

		out = append(out, el.String())
	}

	return strings.Join(out, "\n"), nil
}
