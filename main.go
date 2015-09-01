package main

import (
	"fmt"
	"io"
	"os"

	"github.com/andrew-d/go-termutil"
	"offend.me.uk/thiy/lib"
)

func printHelp() {
	fmt.Println(`Usage: thiy [FILE]

Reads FILE as YAML, converts it into HTML, and outputs the result.

If FILE is missing or "-", input will be read from stdin.

The input document should be a map where each element represent an HTML node. The contents of each element should be a string, an array of strings and maps, or another map representing the contents of the HTML node.

The element keys take the form:

  <TAG>[#<ID>][.CLASS...][(ATTRIBUTE...)]

Example:

  head:
    title: My first page
  body:
    h1: Hello, world
    div#content:
      p:
      - Generated by
      - a(href=https://github.com/stilvoid/thiy): Thiy

Result:

  <!DOCTYPE html>
  <html>
    <head>
      <title>My first page</title>
    </head>
    <body>
      <h1>Hello, world</h1>
        <p class="title">Part 1</p>
        <p>
          Generated by
          <a href="https://github.com/stilvoid/thiy">Thiy</a>
        </p>
      </div>
    </body>
  </html>
`)
}

func main() {
	var r io.Reader

	if len(os.Args) < 2 || os.Args[1] == "-" {
		if len(os.Args) < 2 && termutil.Isatty(os.Stdin.Fd()) {
			printHelp()
			os.Exit(1)
		}

		r = os.Stdin
	} else {
		file, err := os.Open(os.Args[1])
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		r = file
	}

	output, err := lib.Parse(r)

	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	fmt.Println("<!DOCTYPE html>")
	fmt.Println("<html>")
	fmt.Println(output)
	fmt.Println("</html>")
}
