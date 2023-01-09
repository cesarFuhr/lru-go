package lrugo

import "fmt"

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

	if len(c.m) > c.size {
		delete(c.m, c.tail.key)
		c.delete(c.tail)
	}
}

func (c *LRU[T]) Get(key string) (T, bool) {
	n, ok := c.m[key]
	if !ok {
		var noop T
		return noop, false
	}

	if c.head != n {
		c.delete(n)
		c.insert(n)
	}

	return *n.value, true
}

func (c *LRU[T]) Delete(key string) bool {
	n, ok := c.m[key]
	if !ok {
		return false
	}

	delete(c.m, n.key)
	c.delete(n)
	return true
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

	if n == c.head {
		c.head = c.head.next
		c.head.prev = nil
		return
	}

	if n == c.tail {
		c.tail = c.tail.prev
		c.tail.next = nil
		return
	}

	n.prev.next, n.next.prev = n.next, n.prev
}

func (c *LRU[T]) String() string {
	s := fmt.Sprintln("\nLRU: ", &c)

	s += fmt.Sprintln("\nLinked List:")
	next := c.head
	i := 0
	for next != nil {
		s += fmt.Sprintf("%p | %d - key: %+v \t|\tvalue: %+v\n", next, i, next.key, *next.value)
		next = next.next
		i++
	}

	s += fmt.Sprintln("\nMap: ")
	for k, v := range c.m {
		s += fmt.Sprintf("key: %+v \t|\tvalue: %+v\n", k, v)
	}

	return s
}
