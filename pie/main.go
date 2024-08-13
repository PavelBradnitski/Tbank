package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()
	var a int

	fmt.Fscan(in, &a)
	if a == 1 {
		fmt.Println(0)
	} else {
		for i := 0; i < a; i++ {
			if int(math.Pow(2, float64(i))) >= a {
				fmt.Fprint(out, i)
				break
			}
		}
	}

}
