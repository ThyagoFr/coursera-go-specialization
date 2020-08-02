package main

import (
	"fmt"
)

func BubbleSort(sliceNumbers []int) {

	n := len(sliceNumbers)
	for i := 0; i < n; i++ {
		for index := 0; index < n-i-1; index++ {
			if sliceNumbers[index] > sliceNumbers[index+1] {
				Swap(sliceNumbers, index)
			}
		}
	}

}

func Swap(slice []int, index int) {

	slice[index], slice[index+1] = slice[index+1], slice[index]

}

func main() {
	slice := []int{}
	n := 0
	var number int
	fmt.Print("Enter the numbers: ")
	for n != 10 {
		fmt.Scanf("%d", &number)
		slice = append(slice, number)
		n++
	}
	BubbleSort(slice)
	fmt.Println("Slice sorted:", slice)
}
