package main

import (
	"fmt"
	"io/ioutil"
	"os"

	"offend.me.uk/thiy/lib"

	"github.com/andrew-d/go-termutil"
	"gopkg.in/yaml.v2"
)

func main() {
	if termutil.Isatty(os.Stdin.Fd()) {
		fmt.Fprintln(os.Stderr, "Could not read from stdin")
		os.Exit(1)
	}

	var err error

	input, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error reading input")
		os.Exit(1)
	}

	var parsed yaml.MapSlice

	err = yaml.Unmarshal(input, &parsed)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error parsing yaml: %v\n", err)
	}

	for _, node := range parsed {
		el, err := thiy.TranslateItem(node)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error converting: %v\n", err)
			os.Exit(1)
		}

		fmt.Println(el.String())
	}
}
