package main

func MergeSequential(items []int) {
	defer track(executionTime())
	sortSequential(items)
}

func sortSequential(items []int) {
	length := len(items)

	if length == 1 {
		return
	}

	middle := length / 2

	left := make([]int, middle)
	right := make([]int, length-middle)
	copy(left, items[:middle])
	copy(right, items[middle:])

	sortSequential(left)
	sortSequential(right)

	merge(items, left, right)
}
