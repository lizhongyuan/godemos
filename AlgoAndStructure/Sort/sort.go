package SortFunc

func swap(a *int, b *int) {
	var tmp = *a
	*a = *b
	*b = tmp
}

var arr []int = []int{7,2,1,4,8,5,3,10,6,9}

func partition(arr *[]int, left int, right int) int {
	var curArr []int = *arr

	pivot := curArr[left]
	for ;left < right; {
		for ;left < right && curArr[right] >= pivot; {
			right--
		}
		curArr[left] = curArr[right]
		for ;left < right && curArr[left] <= pivot; {
			left++
		}
		curArr[right] = curArr[left]
	}

	// left == right
	curArr[left] = pivot
	return left
}

func PopSort(arr *[]int) {
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

func QuickSortRecursion(arr *[]int, left int, right int) {
	if left < right {
		mid := partition(arr, left, right)
		QuickSortRecursion(arr, left, mid - 1)
		QuickSortRecursion(arr, mid + 1, right)
	}
}

/*
func main() {
	QuickSort1(&arr, 0, 9)
	fmt.Print(arr)
}
*/