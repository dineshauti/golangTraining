package main

import (
	"fmt"
)

type Node struct {
	data interface{}
	next *Node
}

type List struct {
	head *Node
}



func main() {
	l := List{}
	fmt.Println(l.head)

	n1 := Node{data:555,next:nil}
	fmt.Println(n1.data, n1.next)


}


