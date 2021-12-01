package main

import (
	"fmt"
	"strconv"
)

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

		unicodeChar := []rune(currentNode.Value.(string))[0]
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
	kthNode := positions[numberOfNodes-k]

	return kthNode
}

// sum list
func SumListReverse(l1 *LinkedList, l2 *LinkedList) *LinkedList {
	currentNode := l1.Head
	l1Total := getNodeTotal(currentNode)

	currentNode = l2.Head
	l2Total := getNodeTotal(currentNode)

	sum := l2Total + l1Total

	newList := &LinkedList{}
	str := strconv.Itoa(sum)
	for _, char := range str {
		fmt.Println(char, "char")
		newList.Prepend(int(char))
	}

	return newList
}

func getNodeTotal(currentNode *Node) int {
	currentUnit := 1
	Total := 0
	for {
		if currentNode == nil {
			break
		}

		value := currentNode.Value
		valueint := value.(int)
		valueint = valueint * currentUnit
		currentUnit = currentUnit * 10
		Total = Total + valueint
		currentNode = currentNode.Next
	}

	return Total
}

// check if a linked list is a palindrome
func PanlindromeLinkedList(head *Node) bool {
	unicodeMax := 144697
	unicodeTracker := []int{}
	unevenTracker := []int{}
	unevenCharCount := 0

	for i := 0; i < unicodeMax; i++ {
		unicodeTracker = append(unicodeTracker, 0)
		unevenTracker = append(unevenTracker, 0)
	}

	currentNode := head
	lengthOfList := 0
	for {
		if currentNode == nil {
			break
		}

		fmt.Print(currentNode.Value)
		value := currentNode.Value
		unicodeChar := []rune(value.(string))[0]
		lengthOfList++
		unicodeTracker[unicodeChar]++

		currentNode = currentNode.Next
	}

	isEven := lengthOfList%2 == 0
	currentNode = head
	if isEven == true {
		for {
			if currentNode == nil {
				break
			}
	
			value := currentNode.Value
			unicodeChar := []rune(value.(string))[0]
			sumValue := unicodeTracker[unicodeChar]
			currentNode = currentNode.Next
			if sumValue%2 != 0 {
				fmt.Println("\nnot a palindrome: length is even but had a non repeat character")
				return false
			}
		}
	} else {
		for {
			if currentNode.Next == nil {
				break
			}

			value := currentNode.Value
			unicodeChar := []rune(value.(string))[0]
			sumValue := unicodeTracker[unicodeChar]
			currentNode = currentNode.Next

			if sumValue%2 != 0 {
				unevencount := unevenTracker[unicodeChar]
				if unevencount > 0 {
					break
				} else {
					unevenTracker[unicodeChar]++
					unevenCharCount++
					if unevenCharCount > 1 {
						fmt.Println("\nnot a palindrome: length is odd but had more than 1 non repeat character")
						return false
					}
				}
			}
		}
	}

	fmt.Println("\nperfect palindrome")
	return true
}
