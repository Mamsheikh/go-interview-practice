package main

import (
	"fmt"
)

func main() {
	// Example sorted array for testing
	arr := []int{1, 3, 5, 7, 9, 11, 13, 15, 17, 19}

	// Test binary search
	target := 7
	index := BinarySearch(arr, target)
	fmt.Printf("BinarySearch: %d found at index %d\n", target, index)

	// Test recursive binary search
	recursiveIndex := BinarySearchRecursive(arr, target, 0, len(arr)-1)
	fmt.Printf("BinarySearchRecursive: %d found at index %d\n", target, recursiveIndex)

	// Test find insert position
	insertTarget := 8
	insertPos := FindInsertPosition(arr, insertTarget)
	fmt.Printf("FindInsertPosition: %d should be inserted at index %d\n", insertTarget, insertPos)
}

// BinarySearch performs a standard binary search to find the target in the sorted array.
// Returns the index of the target if found, or -1 if not found.
func BinarySearch(arr []int, target int) int {
	// TODO: Implement this function
	if len(arr) == 0 {
		return -1
	}

	if len(arr) == 1 {

		if arr[0] == target {
			return 0
		} else if arr[0] != target {
			return -1
		}

	}
	left, right := 0, len(arr)-1

	for left <= right {
		mid := left + (right-left)/2
		if arr[mid] == target {
			return mid
		} else if arr[mid] < target {
			left = mid + 1
		} else {
			right = mid - 1
		}

	}
	return -1
}

// BinarySearchRecursive performs binary search using recursion.
// Returns the index of the target if found, or -1 if not found.
func BinarySearchRecursive(arr []int, target int, left int, right int) int {
	// TODO: Implement this function
	if left > right {
		return -1
	}
	if len(arr) == 0 {
		return -1
	}

	if len(arr) == 1 {

		if arr[0] == target {
			return 0
		} else if arr[0] != target {
			return -1
		}

	}

	mid := left + (right-left)/2
	if arr[mid] == target {
		return mid
	}
	if arr[mid] < target {
		return BinarySearchRecursive(arr, target, mid+1, right)
	}
	if arr[mid] > target {
		return BinarySearchRecursive(arr, target, left, mid-1)
	}
	return -1
}

// FindInsertPosition returns the index where the target should be inserted
// to maintain the sorted order of the array.
func FindInsertPosition(arr []int, target int) int {
	// TODO: Implement this function

	if len(arr) == 0 {
		return 0
	}

	if target <= arr[0] {
		return 0
	}

	if target > arr[len(arr)-1] {
		return len(arr)
	}

	left, right := 0, len(arr)-1

	for left <= right {
		mid := left + (right-left)/2
		if arr[mid] == target {
			return mid
		} else if arr[mid] < target {
			left = mid + 1
		} else {
			right = mid - 1
		}

	}
	return left
}
