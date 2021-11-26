package main

// remove duplicates from an unsorted linked list
func RemoveDups(l *LinkedList) bool {
	unicodeMax := 144697
	unicodeTracker := []int{}
	for i := 0; i < unicodeMax; i++ {
		unicodeTracker = append(unicodeTracker, 0)
	}

	currentNode := l.Head
	for {
		if currentNode == nil {
			return false
		}

		unicodeChar := []rune(currentNode.Value)[0] 
		unicodeTracker[unicodeChar]++
		if unicodeTracker[unicodeChar] > 1 {
			l.RemoveNode(currentNode)
		}
		currentNode = currentNode.Next
	}
}

// kth to last of a single linked list
func ReturnkthToLast(l *LinkedList, k int) *Node {
	var positions []*Node

	currentPosition := 0
	currentNode := l.Head
	for {
		if currentNode == nil {
			break
		}
		positions = append(positions, currentNode)
	  currentPosition++
		currentNode = currentNode.Next
	}

	numberOfNodes := len(positions)
	kthNode := positions[numberOfNodes - k]

	return kthNode
}
