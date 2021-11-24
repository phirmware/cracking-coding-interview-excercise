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
