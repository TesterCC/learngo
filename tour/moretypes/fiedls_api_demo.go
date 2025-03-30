package main

import (
	"fmt"
	"strings"
)

// https://go-zh.org/pkg/strings/#Fields

func main() {
	fmt.Printf("Fields are: %q", strings.Fields("  foo bar  baz   "))
	fmt.Println("---------------------------------------")
	fmt.Printf("Second test: %q\n", strings.Fields(" "))
	fmt.Printf("Third test: %q\n", strings.Fields("alice            bob  chris     david"))
}
