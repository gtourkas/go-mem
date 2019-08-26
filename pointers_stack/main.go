package pointers_stack

func main()  {

	x := 1
	y := 2
	sum := 0

	add(x,y,&sum) // sharing down: typically to the stack

	println(sum)
}

func add(x int, y int, sum *int) {
	*sum = x + y
}
