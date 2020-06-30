package main

import "fmt"

func binarySearch(Nums []int, Target int) int {
	start := 0
	end := len(Nums) - 1
	for start + 1 < end {
		mid := start + (end-start)/2
		if Nums[mid] == Target {
			return mid
		}else if Nums[mid] > Target {
			end = mid
		}else if Nums[mid] < Target {
			start = mid
		}
	}

	if Nums[start] == Target {
		return start
	}

	if Nums[end] == Target {
		return end
	}

	return -1
}

func main() {
	point := binarySearch([]int{1, 2, 3, 4, 5, 6, 7, 8}, 7)
	fmt.Println(point)

}