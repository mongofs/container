package list

type Node struct {
	next, prev *Node
	list       *List
	Value      interface{}
}

func (e *Node) Next() *Node {
	if p := e.next; e.list != nil && p != &e.list.sentinel {
		return p
	}
	return nil
}

func (e *Node) Prev() *Node {
	if p := e.prev; e.list != nil && p != &e.list.sentinel {
		return p
	}
	return nil
}

type List struct {
	sentinel Node
	len      int
}

func New() *List { return (&List{}).Init() }

// ==================== API =================

func (l *List) Head() *Node {
	return l.sentinel.Next()
}

func (l *List) Tail() *Node {
	return l.sentinel.Prev()
}

// Pop 弹出啦了
func (l *List) Pop() *Node {

}

func (l *List) Push(content interface{}) *Node {
	l.LazyInit()
	newNode := &Node{Value: content, list: l}
	lastNo := l.sentinel.prev
	return l.add(lastNo, newNode)
}

// ==================== Action ==============

func (l *List) Init() *List {
	l.sentinel.next = &l.sentinel
	l.sentinel.prev = &l.sentinel
	l.sentinel.list = l
	l.len = 0
	return l
}

func (l *List) LazyInit() {
	if l.sentinel.next == nil {
		l.Init()
	}
}

func (l *List) add(target, node *Node) *Node {
	node.prev = target
	node.next = target.next
	target.next = node
	node.next.prev = node
	l.len++
	return node
}

func (l *List) del(node *Node) *Node {
	prev := node.Prev()
	node.prev = nil
	prev.next = node.next
	if node.next != nil {
		node.next.prev = prev
		node.next = nil
	}
	l.len--
	return node
}
