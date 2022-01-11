package piscine

type uNodeL struct {
	Data interface{}
	Next *NodeL
}

func ListAt(l *NodeL, pos int) *NodeL {
	it := l
	inc := 0
	for it != nil {
		if pos == inc {
			return it
		}
		inc++
		it = it.Next
	}
	return nil
}
