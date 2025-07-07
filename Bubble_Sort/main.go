package main

import "fmt"

func bubbleSort(nums []int) []int {

	temp := 0

	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[i] > nums[j] {

				temp = nums[i]
				nums[i] = nums[j]
				nums[j] = temp

			}

		}
	}

	return nums

}

func main() {

	nums := []int{5, 8, 2, 4, 0, 1, 3, 7, 9, 6}

	fmt.Println(bubbleSort(nums))
}
