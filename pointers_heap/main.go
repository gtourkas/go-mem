package pointers_heap

func main()  {

	x := 1
	y := 2
	var sum *int

	sum = add(x,y) // sharing up: typically to the heap

	println(*sum)
}

func add(x int, y int) *int {
	sum := x + y
	return &sum
}
