package answer

import (
	"fmt"
)

func removeDuplicates(arr []int) []int {
	uniqueArr := []int{}
	for _, element := range arr {
		if !contains(uniqueArr, element) {
			uniqueArr = append(uniqueArr, element)
		}
	}
	return uniqueArr
}

func contains(arr []int, elem int) bool {
	for _, value := range arr {
		if value == elem {
			return true
		}
	}
	return false
}

func Q3() {
	inputArr := []int{3, 5, 5, 7, 8, 8, 9, 19, 19}
	uniqueArr := removeDuplicates(inputArr)
	fmt.Println("Input:", inputArr)
	fmt.Println("Output:", uniqueArr)
}