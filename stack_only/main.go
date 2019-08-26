package stack_only

func main()  {

	x := 1
	y := 2

	sum := add(x,y)

	println(sum)
}

func add(x int, y int) int {
	return x + y
}
