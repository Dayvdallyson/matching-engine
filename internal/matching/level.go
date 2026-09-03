package matching

import (
	"sync"

	"github.com/dayvd/matching-engine/internal/domain"
)

type node struct {
	order domain.Order
	prev  *node
	next  *node
}

var nodePool = sync.Pool{
	New: func() any { return new(node) },
}

func acquireNode(o domain.Order) *node {
	n := nodePool.Get().(*node)
	n.order = o
	n.prev = nil
	n.next = nil
	return n
}

func releaseNode(n *node) {
	n.order = domain.Order{}
	n.prev = nil
	n.next = nil
	nodePool.Put(n)
}

type Level struct {
	price    domain.Price
	head     *node
	tail     *node
	totalQty domain.Qty
	count    int
}

func newLevel(p domain.Price) *Level {
	return &Level{price: p}
}

func (l *Level) Price() domain.Price { return l.price }

func (l *Level) TotalQty() domain.Qty { return l.totalQty }

func (l *Level) Count() int { return l.count }

func (l *Level) IsEmpty() bool { return l.head == nil }

func (l *Level) Front() *node { return l.head }

func (l *Level) pushBack(n *node) {
	n.prev = l.tail
	n.next = nil
	if l.tail != nil {
		l.tail.next = n
	} else {
		l.head = n
	}
	l.tail = n
	l.totalQty += n.order.Remaining
	l.count++
}

func (l *Level) unlink(n *node) {
	if n.prev != nil {
		n.prev.next = n.next
	} else {
		l.head = n.next
	}
	if n.next != nil {
		n.next.prev = n.prev
	} else {
		l.tail = n.prev
	}
	n.prev = nil
	n.next = nil
	l.totalQty -= n.order.Remaining
	l.count--
}

func (l *Level) reduceQty(delta domain.Qty) {
	l.totalQty -= delta
}
