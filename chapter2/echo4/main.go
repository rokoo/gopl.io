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
	var x, y int
	fmt.Println(x==y, &x==&y)
	fmt.Println(&x)
	fmt.Println(*&x)
	fmt.Println(*sep)
	fmt.Println(*n)

	p := new(int)
	q := new(int)
	fmt.Println(q==p)
	fmt.Println(*p==*q)
}