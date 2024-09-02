package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()
	var numberAmount, rotateAmount int

	fmt.Fscan(in, &numberAmount, &rotateAmount)
	numbers := make([]int, numberAmount)
	for j := 0; j < numberAmount; j++ {
		fmt.Fscan(in, &(numbers[j]))
	}
	var maxDiff int64
	maxDiff, _ = findMaxDiffNumbers(numbers, rotateAmount)
	//fmt.Printf("Максимальная разница: %d, число: %d\n", maxDiff, maxDiffNumber)
	fmt.Println(maxDiff)
}

func findMaxDiffNumbers(numbers []int, maxReplacements int) (int64, []int) {
	maxDiff := int64(0)
	maxDiffNumbers := []int{}

	if len(numbers) == 1 {
		// Если срез содержит одно число, то обрабатываем его отдельно
		maxDiff, maxDiffNumbers = findMaxDiffForSingleNumber(int64(numbers[0]), maxReplacements)
	} else {
		//  Если срез содержит более одного числа, то  выполняем обычный алгоритм
		for i := 0; i < len(numbers); i++ {
			for j := i + 1; j < len(numbers); j++ {
				// Рекурсивно находим максимальную разницу для двух чисел
				diff, modifiedNumbers := findMaxDiffForTwoNumbers(
					int64(numbers[i]),
					int64(numbers[j]),
					maxReplacements,
					0,
					[]int{},
				)

				// Обновляем максимум
				if diff > maxDiff {
					maxDiff = diff
					maxDiffNumbers = modifiedNumbers
				}
			}
		}
	}

	return maxDiff, maxDiffNumbers
}

func findMaxDiffForSingleNumber(number int64, maxReplacements int) (int64, []int) {
	maxDiff := int64(0)
	maxDiffNumber := int(number)
	for i := 0; i <= maxReplacements; i++ {
		diff, modifiedNumber := findMaxDiffForReplacements(number, i, 0)
		if diff > maxDiff {
			maxDiff = diff
			maxDiffNumber = modifiedNumber
		}
	}
	return maxDiff, []int{maxDiffNumber}
}

func findMaxDiffForTwoNumbers(
	number1 int64,
	number2 int64,
	replacementsLeft int,
	currentDiff int64,
	modifiedNumbers []int,
) (int64, []int) {
	if replacementsLeft == 0 {
		return currentDiff, modifiedNumbers
	}

	maxDiff := currentDiff
	maxDiffNumbers := modifiedNumbers

	// Проверяем число 1
	for i := 0; i < len(strconv.Itoa(int(number1))); i++ {
		modifiedNumber1 := replaceDigit(number1, i, 9)
		diff, newNumbers := findMaxDiffForTwoNumbers(
			modifiedNumber1,
			number2,
			replacementsLeft-1,
			currentDiff+modifiedNumber1-number1,
			append(modifiedNumbers, int(modifiedNumber1)),
		)

		if diff > maxDiff {
			maxDiff = diff
			maxDiffNumbers = newNumbers
		}
	}

	// Проверяем число 2
	for i := 0; i < len(strconv.Itoa(int(number2))); i++ {
		modifiedNumber2 := replaceDigit(number2, i, 9)
		diff, newNumbers := findMaxDiffForTwoNumbers(
			number1,
			modifiedNumber2,
			replacementsLeft-1,
			currentDiff+modifiedNumber2-number2,
			append(modifiedNumbers, int(modifiedNumber2)),
		)

		if diff > maxDiff {
			maxDiff = diff
			maxDiffNumbers = newNumbers
		}
	}

	return maxDiff, maxDiffNumbers
}

func findMaxDiffForReplacements(number int64, replacementsLeft int, currentDiff int64) (int64, int) {
	if replacementsLeft == 0 {
		return currentDiff, int(number)
	}

	numberStr := strconv.Itoa(int(number))
	maxDiff := currentDiff
	maxDiffNumber := int(number)

	for i := 0; i < len(numberStr); i++ {
		// Заменяем цифру на 9
		modifiedNumber := replaceDigit(number, i, 9)

		// Рекурсивно вызываем функцию для оставшихся замен
		diff, newNumber := findMaxDiffForReplacements(modifiedNumber, replacementsLeft-1, currentDiff+modifiedNumber-number)

		// Обновляем максимум
		if diff > maxDiff {
			maxDiff = diff
			maxDiffNumber = newNumber
		}
	}

	return maxDiff, maxDiffNumber
}

func replaceDigit(number int64, digitIndex int, newDigit int) int64 {
	numberStr := strconv.Itoa(int(number))
	numberStr = numberStr[:digitIndex] + strconv.Itoa(newDigit) + numberStr[digitIndex+1:]
	newNumber, _ := strconv.ParseInt(numberStr, 10, 64)
	return newNumber
}
