package main

import (
	"flag"
	"fmt"
	"strings"
)

/*
P25 in 2.3.2 end
go的flag类似 python的argparse
go build echo4_flag

./echo4_flag --help
./echo4_flag -n a bc de fg     output: abcdefg
./echo4_flag -s / a bcd e      output: a/bcd/e
*/

var n = flag.Bool("n", false, "omit trailing newline")
var sep = flag.String("s", "", "separator")

func main() {
	flag.Parse()
	fmt.Print(strings.Join(flag.Args(), *sep))

	if !*n {
		fmt.Println()
	}
}
