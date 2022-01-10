package piscine

type iNodeL struct {
	Data interface{}
	Next *NodeL
}

type iList struct {
	Head *NodeL
	Tail *NodeL
}

func ListLast(l *List) interface{} {
	if l.Head == nil {
		return nil
	} else {
		now := l.Head
		for now.Next != nil {
			now = now.Next
		}
		return now.Data
	}
}
