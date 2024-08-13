package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var in *bufio.Reader
	var out *bufio.Writer
	in = bufio.NewReader(os.Stdin)
	out = bufio.NewWriter(os.Stdout)
	defer out.Flush()
	var a, b, c, d int

	fmt.Fscan(in, &a, &b, &c, &d)
	if d <= b {
		fmt.Println(a)
	} else {
		fmt.Println(a + (d-b)*c)
	}

}
