package Demo

import (
	"fmt"
	"../AlgoAndStructure/Sort"
)


func ShowSortFuncs() {
	var arr []int = []int{3,2,1,4,6,5,7,10,8,9}
	SortFunc.PopSort(&arr)
	fmt.Println("PopSort:", arr)

	SortFunc.QuickSortRecursion(&arr, 0, len(arr) - 1)
	fmt.Println("QuickSortRecursion:", arr)
}
