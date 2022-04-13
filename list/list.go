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

func New()*List{ return (&List{}).Init()}

// ==================== API =================

func (l *List) Push(content interface{}) *Node {
	l.LazyInit()
	newNode := &Node{Value: content, list: l}
	lastNo := l.sentinel.prev
	return l.add(lastNo,newNode)
}

// 删除节点

// ==================== Action ==============

func (l *List)Init ()*List{
	l.sentinel.next = &l.sentinel
	l.sentinel.prev = &l.sentinel
	l.sentinel.list = l
	l.len =0
	return l
}

func (l *List)LazyInit (){
	if l.sentinel.next == nil {
		 l.Init()
	}
}


func (l *List) add(before, node *Node) *Node {
	next := before.Next()
	before.next = node
	node.prev = before
	node.next = next
	next.prev = node
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
