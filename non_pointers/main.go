package non_pointers

import "math/rand"

func main() {

	const n = 10 // try with n := 10

	nums_stack := make([]int, n)
	randomize_stack(nums_stack)

	println(nums_stack)

	nums_heap := randomize_heap(n)

	println(nums_heap)
}

func randomize_stack(nums []int) {

	for i:=0; i<len(nums); i++ {
		nums[i] = rand.Intn(100)
	}

}

func randomize_heap(n int) []int {
	nums_heap := make([]int, n)

	for i:=0; i<n; i++ {
		nums_heap[i] = rand.Intn(100)
	}

	return nums_heap
}