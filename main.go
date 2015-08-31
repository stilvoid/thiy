package main

import (
	"fmt"
	"os"

	"offend.me.uk/thiy/lib"

	"github.com/andrew-d/go-termutil"
)

func main() {
	if termutil.Isatty(os.Stdin.Fd()) {
		fmt.Fprintln(os.Stderr, "Could not read from stdin")
		os.Exit(1)
	}

	output, err := lib.Parse(os.Stdin)

	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	fmt.Println("<!DOCTYPE html>")
	fmt.Println("<html>")
	fmt.Println(output)
	fmt.Println("</html>")
}
