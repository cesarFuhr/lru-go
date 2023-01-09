package lrugo

type Node[T any] struct {
	next, prev *Node[T]
	key        string
	value      *T
}

type LRU[T any] struct {
	size int
	head *Node[T]
	tail *Node[T]
	m    map[string]*Node[T]
}

func NewLRU[T any](size int) *LRU[T] {
	m := make(map[string]*Node[T])
	return &LRU[T]{
		size: size,
		m:    m,
	}
}

func (c *LRU[T]) Insert(key string, value T) {
	n := Node[T]{
		key:   key,
		value: &value,
	}

	c.insert(&n)

	c.m[key] = &n

	if len(c.m) >= c.size {
		c.delete(c.tail)
	}
}

func (c *LRU[T]) Get(key string) (T, bool) {
	n, ok := c.m[key]
	if !ok {
		var noop T
		return noop, false
	}

	return *n.value, true
}

func (c *LRU[T]) insert(n *Node[T]) {
	if n == nil {
		return
	}

	if c.head == nil {
		c.head = n
		c.tail = n
		return
	}

	n.next = c.head
	c.head.prev = n
	c.head = n
}

func (c *LRU[T]) delete(n *Node[T]) {
	if n == nil {
		return
	}

	if n == c.tail {
		c.tail = n.prev
		delete(c.m, n.key)
	}
}
