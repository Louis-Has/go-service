package sorts

import (
	"math/rand"
)

// SelectionSort 选择排序 (selection sort) O(n^2) O(1) 不稳定
func SelectionSort(arr []int) []int {

	for i := 0; i < len(arr); i++ {
		min := i

		for j := i + 1; j < len(arr); j++ {
			if arr[j] < arr[min] {
				min = j
			}
		}
		arr[i], arr[min] = arr[min], arr[i]
	}

	return arr
}

// BubbleSort 冒泡排序 (bubble sort) O(n^2)->O(n) O(1) 稳定
func BubbleSort(arr []int) []int {
	swapped := true

	for swapped {
		swapped = false
		for i := 0; i < len(arr)-1; i++ {
			if arr[i+1] < arr[i] {
				arr[i+1], arr[i] = arr[i], arr[i+1]
				swapped = true
			}
		}
	}

	return arr
}

// InsertionSort 插入排序 (insertion sort) O(n^2)->O(n) O(1) 稳定
func InsertionSort(arr []int) []int {
	for i := 1; i < len(arr); i++ {
		tmpArr := arr[i]
		j := i
		for ; j > 0 && arr[j-1] >= tmpArr; j-- {
			arr[j] = arr[j-1]
		}
		arr[j] = tmpArr
	}
	return arr
}

// ShellSort 希尔排序 (shell sort) O(n^2)->O(n) O(1) 不稳定
func ShellSort(arr []int) []int {
	for d := int(len(arr) / 2); d > 0; d /= 2 {
		for i := d; i < len(arr); i++ {
			for j := i; j >= d && arr[j-d] > arr[j]; j -= d {
				arr[j], arr[j-d] = arr[j-d], arr[j]
			}
		}
	}
	return arr
}

// 归并排序 (merge sort) O(logN) O(N) 稳定
func merge(a []int, b []int) []int {

	var r = make([]int, len(a)+len(b))
	var i = 0
	var j = 0

	for i < len(a) && j < len(b) {

		if a[i] <= b[j] {
			r[i+j] = a[i]
			i++
		} else {
			r[i+j] = b[j]
			j++
		}

	}

	for i < len(a) {
		r[i+j] = a[i]
		i++
	}
	for j < len(b) {
		r[i+j] = b[j]
		j++
	}

	return r

}

// Mergesort 合并两个数组 O(N/logN)
func Mergesort(items []int) []int {

	if len(items) < 2 {
		return items

	}

	var middle = len(items) / 2
	var a = Mergesort(items[:middle])
	var b = Mergesort(items[middle:])
	return merge(a, b)

}

// QuickSort 三向切分快速排序 (quick sort) O(logN)->O(N^2) O(log2N) 不稳定
func QuickSort(arr []int) []int {

	if len(arr) <= 1 {
		return arr
	}

	pivot := arr[rand.Intn(len(arr))]

	lowPart := make([]int, 0, len(arr))
	highPart := make([]int, 0, len(arr))
	middlePart := make([]int, 0, len(arr))

	for _, item := range arr {
		switch {
		case item < pivot:
			lowPart = append(lowPart, item)
		case item == pivot:
			middlePart = append(middlePart, item)
		case item > pivot:
			highPart = append(highPart, item)
		}
	}

	lowPart = QuickSort(lowPart)
	highPart = QuickSort(highPart)

	lowPart = append(lowPart, middlePart...)
	lowPart = append(lowPart, highPart...)

	return lowPart
}

// 堆排序 (heap sort)
type maxHeap struct {
	slice    []int
	heapSize int
}

func buildMaxHeap(slice []int) maxHeap {
	h := maxHeap{slice: slice, heapSize: len(slice)}
	for i := len(slice) / 2; i >= 0; i-- {
		h.MaxHeapify(i)
	}
	return h
}

func (h maxHeap) MaxHeapify(i int) {
	l, r := 2*i+1, 2*i+2
	max := i
	size := h.size()

	if l < size && h.slice[l] > h.slice[max] {
		max = l
	}
	if r < size && h.slice[r] > h.slice[max] {
		max = r
	}
	if max != i {
		h.slice[i], h.slice[max] = h.slice[max], h.slice[i]
		h.MaxHeapify(max) // important
	}
}

func (h maxHeap) size() int { return h.heapSize }

// HeapSort O(logN) O(1) 不稳定
func HeapSort(slice []int) []int {
	h := buildMaxHeap(slice)
	for i := len(h.slice) - 1; i >= 1; i-- {
		h.slice[0], h.slice[i] = h.slice[i], h.slice[0]
		h.heapSize--
		h.MaxHeapify(0)
	}
	return h.slice
}

// 计数排序（Counting Sort） O(n) O(n)
// 输入的数据必须是有确定范围的整数

// 桶排序（Bucket Sort）O(n) O(n)

// 基数排序（Radix Sort）O(n) O(n)
