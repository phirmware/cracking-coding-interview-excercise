package main

import "fmt"

func main() {
	l := &LinkedList{}

	l.Prepend("a")
	l.append("b")
	l.append("b")
	l.append("a")
	l.append("c")

	fmt.Println("Start removing duplicates")
	RemoveDups(l)
	fmt.Println("Done removing duplicates")

	moveThroughList(l)

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
