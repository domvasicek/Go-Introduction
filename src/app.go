package main

import (
	"fmt"
	"log"
	"math/rand"
	"runtime"
	"time"
)

func swap(first string, second string) (string, string) {
	return second, first
}

func executionTime() time.Time {
	return time.Now()
}

func track(startTime time.Time) {
	endTime := time.Now()
	log.Println("Merge done in :", endTime.Sub(startTime))
}

func main() {

	type Person struct {Name, Surname string}
	person := Person {
		Name:    "John",
		Surname: "Doe",
	}

	fmt.Println(person)
	fmt.Println(runtime.NumCPU())

	return1, return2 := swap("First", "Second")
	fmt.Println(return1, return2)

	rand.Seed(987654321)
	toSortSequential := rand.Perm(50 * 1000 * 1000)
	toSortParallel := make([]int, len(toSortSequential))
	copy(toSortParallel, toSortSequential)

	fmt.Println(toSortSequential[:100])

	MergeSequential(toSortSequential)
	fmt.Println(toSortSequential[:100])

	MergeParallel(toSortParallel)
	fmt.Println(toSortParallel[:100])
}

func merge(result, left, right []int) {
	var leftCounter, rightCounter, i int

	for leftCounter < len(left) || rightCounter < len(right) {
		if leftCounter < len(left) && rightCounter < len(right) {
			if left[leftCounter] <= right[rightCounter] {
				result[i] = left[leftCounter]
				leftCounter++
			} else {
				result[i] = right[rightCounter]
				rightCounter++
			}
		} else if leftCounter < len(left) {
			result[i] = left[leftCounter]
			leftCounter++
		} else if rightCounter < len(right) {
			result[i] = right[rightCounter]
			rightCounter++
		}
		i++
	}
}