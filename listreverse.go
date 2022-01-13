package piscine

type yNodeL struct {
	Data interface{}
	Next *NodeL
}

type yList struct {
	Head *NodeL
	Tail *NodeL
}

func ListReverse(l *List) {
	current := l.Head
	prev := l.Head
	prev = nil
	for current != nil {
		next := current.Next
		current.Next = prev
		prev = current
		current = next
	}
	temp := l.Head
	l.Head = l.Tail
	l.Tail = temp
}
