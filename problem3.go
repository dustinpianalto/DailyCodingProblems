package main

import (
	"log"
	"strings"
)

//Given the root to a binary tree, implement serialize(root), which serializes the tree into a string, and deserialize(s), which deserializes the string back into the tree.
//
//For example, given the following Node class
//
//class Node:
//    def __init__(self, val, left=None, right=None):
//        self.val = val
//        self.left = left
//        self.right = right
//The following test should pass:
//
//node = Node('root', Node('left', Node('left.left')), Node('right'))
//assert deserialize(serialize(node)).left.left.val == 'left.left'

type Node struct {
	Val   string
	Left  *Node
	Right *Node
}

func (n *Node) Serialize() string {
	if *n == (Node{}) {
		return ""
	}
	return n.Val + "\u200b" + n.Left.Serialize() + "\u200b" + n.Right.Serialize() + "\u200b"
}

func Deserialize(s string) *Node {
	sc := make(chan string)
	go func() {
		for _, val := range strings.Split(s, "\u200b") {
			sc <- val
		}
		close(sc)
	}()
	return processNode(sc)
}

func processNode(sc <-chan string) *Node {
	root := Node{}
	for val := range sc {
		if val == "" {
			return &root
		}
		root.Val = val
		root.Left = processNode(sc)
		root.Right = processNode(sc)
	}
	return &root
}

func main() {
	n := &Node{
		Val: "root",
		Left: &Node{
			Val: "left",
			Left: &Node{
				Val:   "left.left",
				Left:  &Node{},
				Right: &Node{},
			},
			Right: &Node{},
		},
		Right: &Node{
			Val:   "right",
			Left:  &Node{},
			Right: &Node{},
		},
	}
	s := n.Serialize()
	log.Println(s)
	log.Println(Deserialize(s).Left.Left.Val)
}
