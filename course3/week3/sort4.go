package main

import (
	"fmt"
	"sort"
	"sync"
)

func sortSlicePiece(wg *sync.WaitGroup, pieceOfSlice []int) {
	fmt.Println("Sub-array to sort", pieceOfSlice)
	sort.Ints(pieceOfSlice)
	wg.Done()
}

func main() {
	var wg sync.WaitGroup
	wg.Add(4)
	numbers := []int{}
	var n int

	fmt.Println("Enter the number of integers : ")
	fmt.Scanf("%d", &n)

	for i := 0; i < n; i++ {
		var number int
		fmt.Scanf("%d", &number)
		numbers = append(numbers, number)
	}
	division := n / 4.0
	fmt.Println("Array to sort :", numbers)
	fmt.Println()

	slice1 := numbers[0:division]
	slice2 := numbers[division : 2*division]
	slice3 := numbers[2*division : 3*division]
	slice4 := numbers[3*division : n]

	go sortSlicePiece(&wg, slice1)
	go sortSlicePiece(&wg, slice2)
	go sortSlicePiece(&wg, slice3)
	go sortSlicePiece(&wg, slice4)

	wg.Wait()
	sort.Ints(numbers)
	fmt.Println(numbers)

}
