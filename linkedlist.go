package main

type Node struct {
	Next     *Node
	Previous *Node
	Value    string
}

type LinkedList struct {
	Head *Node
	Tail *Node
}

func (l *LinkedList) append(value string) *LinkedList {
	if l.Tail == nil {
		l.Tail = &Node{Value: value, Previous: nil, Next: nil}
		l.Head = l.Tail
		return l
	}

	currentTail := l.Tail
	t := &Node{Value: value, Next: nil, Previous: currentTail}
	l.Tail = t
	t.Previous.Next = t

	return l
}

func (l *LinkedList) Prepend(value string) *LinkedList {
	if l.Head == nil {
		l.Head = &Node{Previous: nil, Value: value, Next: nil}
		l.Tail = l.Head

		return l
	}

	currentHead := l.Head
	h := &Node{Previous: nil, Value: value, Next: currentHead}
	l.Head = h
	h.Next.Previous = h

	return l
}

func (l *LinkedList) RemoveTail() *LinkedList {
	if l.Tail == nil {
		return nil
	}

	if l.Tail == l.Head {
		l.Tail = nil
		l.Head = nil

		return l
	}

	nextTail := l.Tail.Previous
	l.Tail = nextTail
	l.Tail.Next = nil

	return l
}

func (l *LinkedList) RemoveNode(n *Node) bool {
	if n.Previous == nil {
		n.Next.Previous = nil
		return true
	}

	if n.Next == nil {
		n.Previous.Next = nil
		return true
	}
	
	n.Previous.Next = n.Next
	n.Next.Previous = n.Previous

	return true
}

func (l *LinkedList) RemoveHead() *LinkedList {
	if l.Head == nil {
		return nil
	}

	if l.Head == l.Tail {
		l.Head = nil
		l.Tail = nil

		return l
	}

	nexthead := l.Head.Next
	l.Head = nexthead
	l.Head.Previous = nil

	return l
}

func (l *LinkedList) Search(value string) *Node {
	currentNode := l.Head

	for {
		if currentNode.Value == value {
			return currentNode
		}
		currentNode = currentNode.Next
	}
}
