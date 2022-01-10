package piscine

type oNodeL struct {
	Data interface{}
	Next *NodeL
}

type oList struct {
	Head *NodeL
	Tail *NodeL
}

func ListSize(l *List) int {
	n := l.Head
	size := 0
	for n != nil {
		size++
		n = n.Next
	}
	return size
}
