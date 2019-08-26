## Golang Heap vs Stack

### Cases

#### Stack Only

No pointers or pointer-behaving variables used at all. Everything is allocated to the stack.

#### Pointers Stack

main's sum variable is "shared down" with add, since it is passed as a parameter. Does not escape to the heap.

#### Pointers Heap

add's sum pointer variable is "shared up" with main, since it is returned. Escapes to the heap.

#### Non-Pointers



### Compile and Run

Need to build with the -m and -l flags: 

* `m` is for printing optimization decisions while running.
* `l` is for disabling inlining. 

This is done as follows:

```
cd <case_dir> 

go build -gcflags "-m -l" main.go
```


### Resources

Understanding Allocations: the Stack and the Heap - GopherCon SG 2019
 
https://youtu.be/ZMZpH4yT7M0


