package main

import (
	"fmt"
)

type Node struct {
	le   *Node
	data interface{}
	ri   *Node
}

func NewNode(left, right *Node) *Node {
	return &Node{left, nil, right}
}

func (n *Node) SetDaata(data interface{}) {
	n.data = data
}

func main() {
	root := NewNode(nil, nil)
	root.SetDaata("root node")
	a := NewNode(nil, nil)
	a.SetDaata("left node")
	b := NewNode(nil, nil)
	b.SetDaata("right node")
	root.le = a
	root.ri = b
	fmt.Printf("%v\n", root)
}
