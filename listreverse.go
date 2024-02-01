package piscine

func ListReverse(l *List) {
	c := l.Head
	var prev *NodeL
	for c != nil {
		n := c.Next
		c.Next = prev
		prev = c
		c = n
	}
	l.Head = prev
}
