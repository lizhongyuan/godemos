package sorts

import "fmt"

//var arr = [6]int{0,1,2,3,4,5}

func Public_fuction() {
	fmt.Println("used in anywhere!")
}

func test(arr []int) {
	var i int
	for i = 0; i < len(arr) ; i++ {
		fmt.Print(arr[i])
	}
}

func test2(arr []int) []int {
	var i int
	for i = 0; i < len(arr); i++ {
		arr[i] = 0
	}
	return arr
}

func swap(a *int, b *int) {
	var tmp = *a
	*a = *b
	*b = tmp
}

func Pop_sort(arr *[]int) {
	var i int
	var j int
	var curArr = *arr
	for i = 0; i < len(curArr) - 1; i++ {
		for j = 0; j < len(curArr) - i - 1; j ++ {
			if (curArr)[j] > (curArr)[j + 1] {
				swap(&curArr[j], &curArr[j + 1])
			}
		}
	}
}

/*
func main() {
	arr := []int{0,1,2,3,4,5,10,9,8,7,6}
	// arr := []int{1,0}

	// test(arr)
	pop_sort(&arr)
	fmt.Print(arr)
}
*/
