package answer

import "fmt"

func mergeSortedArrays(arr1, arr2 []int) []int {
	mergedArr := []int{}
	i, j := 0, 0

	for i < len(arr1) && j < len(arr2) {
		if arr1[i] <= arr2[j] {
			mergedArr = append(mergedArr, arr1[i])
			i++
		} else {
			mergedArr = append(mergedArr, arr2[j])
			j++
		}
	}

	for i < len(arr1) {
		mergedArr = append(mergedArr, arr1[i])
		i++
	}

	for j < len(arr2) {
		mergedArr = append(mergedArr, arr2[j])
		j++
	}

	return mergedArr
}

func Q4() {
	arr1 := []int{1, 3, 6, 10, 14}
	arr2 := []int{2, 4, 7, 8, 13, 20}

	mergedArr := mergeSortedArrays(arr1, arr2)

	fmt.Println("Input 1:", arr1)
	fmt.Println("Input 2:", arr2)
	fmt.Println("Output:", mergedArr)
}