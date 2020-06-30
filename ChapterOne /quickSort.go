package main

import "fmt"

//快排
func quickSort(sli []int) []int {
	sLen := len(sli)
	if sLen <= 1 {
		return sli
	}
	baseNum := sli[0]
	var leftSli []int
	var rightSli []int
	for i := 1; i < sLen; i++ {
		if baseNum > sli[i] {
			leftSli = append(leftSli, sli[i])
		} else {
			rightSli = append(rightSli, sli[i])
		}
	}
	leftSli = quickSort(leftSli)
	rightSli = quickSort(rightSli)
	leftSli = append(leftSli, baseNum)
	return append(leftSli, rightSli...)
}
func main() {
	sli := []int{49, 43, 54, 62, 21, 66, 32, 78, 36, 76, 39}
	sli = quickSort(sli)
	fmt.Println(sli)

}
