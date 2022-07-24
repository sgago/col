package skiplist

type node struct {
	next  []*node
	key   interface{}
	value interface{}
}

type skiplist struct {
	maxLevel int
	head     node
}

func New(maxLevel int) *skiplist {

	l := skiplist{
		maxLevel: maxLevel,
	}

	l.head = node{
		next:  make([]*node, 0, l.maxLevel),
		key:   nil,
		value: nil,
	}

	return &l
}
