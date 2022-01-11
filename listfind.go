package piscine

type eNodeL struct {
	Data interface{}
	Next *NodeL
}

type eList struct {
	Head *NodeL
	Tail *NodeL
}

func CompStr(a, b interface{}) bool {
	return a == b
}

func ListFind(l *List, ref interface{}, comp func(a, b interface{}) bool) *interface{} {
	ite := l.Head
	for ite != nil {
		if comp(ite.Data, ref) {
			return &ite.Data
		}
		ite = ite.Next
	}
	return nil
}
