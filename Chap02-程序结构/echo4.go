package main

import (
	"flag"
	"fmt"
	"strings"
)

var n = flag.Bool("n", false, "omit trailing newline")
var sep = flag.String("s", " ", "separator")

func main() {
	flag.Parse()
	fmt.Print(strings.Join(flag.Args(), *sep))
	if !*n {
		fmt.Println()
	}
}

/*
usage log:
eg1
$ ./echo4.exe a bc def
a bc def

eg2
$ ./echo4.exe -s / a bc def
aC:/Program Files/Git/bcC:/Program Files/Git/def

eg3
$ ./echo4.exe -n a bc def
a bc def

eg4
$ ./echo4.exe -help
Usage of ./echo4.exe:
  -n    omit trailing newline
  -s string
        separator (default " ")
 */
