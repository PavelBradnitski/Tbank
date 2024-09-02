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
	var l, r, count int64

	fmt.Fscan(in, &l, &r)
	minDigits := int(math.Log10(float64(l))) // Минимальное количество цифр в l
	maxDigits := int(math.Log10(float64(r))) // Максимальное количество цифр в r
	if minDigits == 0 {
		for digit := 1; digit <= 9; digit++ {
			if int64(digit) >= l && int64(digit) <= r {
				count++
			}
		}
		minDigits = 1
	}
	for digit := 0; digit <= 9; digit++ {
		for numDigits := minDigits; numDigits <= maxDigits; numDigits++ {
			num := int64(digit)
			// Изменяем условие цикла, чтобы сгенерировать числа с одинаковыми цифрами
			//  с  количеством  цифр  numDigits
			for i := 0; i < numDigits; i++ {
				num = num*10 + int64(digit)
			}
			if num >= l && num <= r {
				count++
			}
		}
	}

	fmt.Fprint(out, count)
}
