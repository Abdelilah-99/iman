package piscine

func ListAt(l *NodeL, pos int) *NodeL {
	if pos < 0 {
		return nil
	}
	if pos == 0 {
		return l
	}
	for i := 0; i < pos; i++ {
		if l != nil {
			l = l.Next
		}
	}
	return l
}
