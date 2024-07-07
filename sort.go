package main

import "fmt"

//冒泡排序 O(n^2)
func BubbleSort(arr []int) []int {
	//控制轮次
	for i := 0; i < len(arr)-1; i++ {
		flag := false
		//控制每一轮排序
		for j := 0; j < len(arr)-i-1; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
				flag = true
			}
		}
		if !flag {
			break
		}
	}
	return arr
}

//归并排序 O(nlogn)
func MergeSort(arr []int) []int {
	if len(arr) <= 1 {
		return arr
	}
	middle := len(arr) / 2
	left := arr[:middle]
	right := arr[middle:]
	left = MergeSort(left)
	right = MergeSort(right)
	return Merge(left, right)
}
func Merge(left, right []int) []int {
	leftIndex := 0
	rightIndex := 0
	arr := make([]int, 0, len(left)+len(right))
	for leftIndex < len(left) && rightIndex < len(right) {
		if left[leftIndex] < right[rightIndex] {
			arr = append(arr, left[leftIndex])
			leftIndex++
		} else {
			arr = append(arr, right[rightIndex])
			rightIndex++
		}
	}
	if leftIndex < len(left) {
		arr = append(arr, left[leftIndex:]...)
	} else {
		arr = append(arr, right[rightIndex:]...)
	}
	return arr
}

//快速排序 O(nlogn)
func QuickSort(arr []int) []int {
	if len(arr) <= 1 {
		return arr
	}
	middle := arr[0]
	left := make([]int, 0, len(arr)/2)
	right := make([]int, 0, len(arr)/2)
	for i := 1; i < len(arr); i++ {
		if arr[i] <= middle {
			left = append(left, arr[i])
		} else {
			right = append(right, arr[i])
		}
	}
	left = QuickSort(left)
	right = QuickSort(right)
	return append(append(left, middle), right...)
}

func main() {
	fmt.Println(BubbleSort([]int{42, 32, 12, 88, 64, 55}))
	fmt.Println(MergeSort([]int{42, 32, 12, 88, 64, 55}))
	fmt.Println(QuickSort([]int{42, 32, 12, 88, 64, 55}))
}
