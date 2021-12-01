package main

import "fmt"

func main() {
	l := &LinkedList{}

	l.append("c")
	l.append("h")
	l.append("i")
	l.append("b")
	l.append("u")
	l.append("z")
	l.append("o")
	l.append("r")

	fmt.Println("checking palindrome")
	result := PanlindromeLinkedList(l.Head)
	fmt.Println(result)
}

func moveThroughList(l *LinkedList) bool {
	currentNode := l.Head
	for {
		if currentNode == nil {
			return false
		}
		fmt.Printf("Node: { \nvalue: %s, \nnext: %+v, \nprevious: %+v }\n\n", currentNode.Value, currentNode.Next, currentNode.Previous)
		currentNode = currentNode.Next
		if currentNode == nil {
			fmt.Println("Done")
			return false
		}
	}
}
