package main

import "fmt"

//冒泡
func bubbleSort(sli []int) []int {
	sLen := len(sli)
	for i := 1; i < sLen; i++ {
		for j := 0; j < (sLen - 1); j++ {
			if sli[i] < sli[j] {
				sli[j], sli[i] = sli[i], sli[j]
			}
		}
	}
	return sli
}

func main() {
	sli := []int{49, 43, 54, 62, 21, 66, 32, 78, 36, 76, 39}
	sli = bubbleSort(sli)
	fmt.Println(sli)
}
