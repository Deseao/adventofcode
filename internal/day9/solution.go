package day9

import (
	"errors"
	"fmt"
	"log"

	"advent/internal/input"
)

var ErrNoInvalid = errors.New("No number found that wasn't a valid sum")
var ErrInputTooShort = errors.New("The provided numbers slice was not longer than the given preamble size")

const preambleLength = 5

func Solve() {
	data := input.NewInput("./input/day9example.txt").AsInts()
	firstInvalid, err := findFirstInvalid(preambleLength, data) //Change 5 to 25
	if err != nil {
		log.Fatalf("Error finding first invalid entry: %s", err)
	}
	fmt.Println("Part 1 Solution was", data[firstInvalid])
	weakness, err := findWeakness(firstInvalid, data)
	if err != nil {
		log.Fatalf("Error finding encryption weakness: %s", err)
	}
	fmt.Println("Part 2 Solution was", weakness)
}

//findFirstInvalid returns the index of the first value in numbers which is not a result of summing any two discrete values of the previous preambleLength values of numbers.
func findFirstInvalid(preambleLength int, numbers []int) (index int, err error) {
	if !(len(numbers) > preambleLength) {
		return 0, ErrInputTooShort
	}
	for i := preambleLength; i < len(numbers); i++ {
		if !isSumOfDiscreteNPrevious(numbers[i], numbers[i-preambleLength:i]) {
			return i, nil
		}
	}
	return 0, ErrNoInvalid
}

func isSumOfDiscreteNPrevious(target int, operands []int) bool {
	//This must check each combination without allowing them to match themselves, so the last index has to skip right?
	//compare operands[0] + operands[1] == target, compare operands[0] + operands[len(operands)-1] == target, compare operands[1] + operands[2] == target, compare operands[len(operands)-2] + operands(len(operands)-1] == target
	for i := 0; i < len(operands)-1; i++ { //i < len(operands) - 1 because I want to run the last loop with one unchecked pair.
		for j := i + 1; j < len(operands); j++ {
			if operands[i]+operands[j] == target {
				return true
			}
		}
	}
	return false
}

func findWeakness(target int, numbers []int) (weakness int, err error) {
	contiguousSet, err := findContiguousSet(target, numbers)
	if err != nil {
		return err
	}
	sort.Ints(contiguousSet)
	return contiguousSet[0] + contiguousSet[len(contiguousSet)-1]
}

func findContiguousSet(target int, numbers []int) (set []int, err error) {

	//compare numbers[0] + numbers[1] to target, then 0 + 1 + 2 to target, then 0 + 1 + 3. Then compare 1 + 2 to target, 1 + 2 + 3 to target... sum from i to end and track index.
	for i, first := range numbers {
		count := 0
		sum := first
		for j := i + 1; j < len(numbers); j++ {
			sum = sum + numbers[j]
		}
	}
	return numbers, nil
}
