package lib

import (
	"bytes"
	"io"
	"io/ioutil"

	"offend.me.uk/thiy/common"

	"gopkg.in/yaml.v2"
)

func Parse(r io.Reader) (string, error) {
	input, err := ioutil.ReadAll(r)
	if err != nil {
		return "", err
	}

	var parsed yaml.MapSlice

	err = yaml.Unmarshal(input, &parsed)
	if err != nil {
		return "", err
	}

	var buf bytes.Buffer

	for _, node := range parsed {
		el, err := translateItem(node)
		if err != nil {
			return "", err
		}

		buf.WriteString(el.String())
	}

	return buf.String(), nil
}
