package main

import "sync"

func MergeParallel(items []int) {
	defer track(executionTime())
	sortParallel(items)
}

func sortParallel(items []int) {
	length := len(items)

	if length > 1 {
		if length <= 2048 {
			// Sequential
			sortSequential(items)
		} else {
			// Parallel

			middle := length / 2
			left := make([]int, middle)
			right := make([]int, length-middle)
			copy(left, items[:middle])
			copy(right, items[middle:])

			var wg sync.WaitGroup
			wg.Add(2)

			go func() {
				defer wg.Done()
				sortParallel(left)
			}()

			go func() {
				defer wg.Done()
				sortParallel(right)
			}()

			wg.Wait()
			merge(items, left, right)
		}
	}
}