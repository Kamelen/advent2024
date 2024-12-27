package utils

import (
	"fmt"
	"strconv"
)

func StringToIntList(stringList []string) ([]int, error) {
	intList := make([]int, len(stringList))
	for i, str := range stringList {
		num, err := strconv.Atoi(str)
		if err != nil {
			return nil, fmt.Errorf("error converting string to int at index %d: %v", i, err)
		}
		intList[i] = num
	}
	return intList, nil
}

func Assert(condition bool, message string) {
	if !condition {
		panic(message)
	}
}

func Abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func SumInts(numbers []int) int {
	sum := 0
	for _, num := range numbers {
		sum += num
	}
	return sum
}
