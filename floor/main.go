package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()
	var workersAmount, leaveTime, workerThatLeaves int

	fmt.Fscan(in, &workersAmount, &leaveTime)
	floor := make([]int, workersAmount)
	for j := 0; j < workersAmount; j++ {
		fmt.Fscan(in, &(floor[j]))
	}

	fmt.Fscan(in, &workerThatLeaves)
	switch {
	case floor[workerThatLeaves-1]-floor[0] <= leaveTime:
		fmt.Fprint(out, floor[workersAmount-1]-floor[0])
	case floor[workersAmount-1]-floor[workerThatLeaves-1] <= leaveTime:
		fmt.Fprint(out, floor[workersAmount-1]-floor[0])
	case floor[workerThatLeaves-1]-floor[0] > floor[workersAmount-1]-floor[workerThatLeaves-1]:
		fmt.Fprint(out, floor[workersAmount-1]-floor[0]+floor[workersAmount-1]-floor[workerThatLeaves-1])
	case floor[workerThatLeaves-1]-floor[0] <= floor[workersAmount-1]-floor[workerThatLeaves-1]:
		fmt.Fprint(out, floor[workersAmount-1]-(2*floor[0])+floor[workerThatLeaves-1])
	}

}
