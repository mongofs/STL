package list

// 这里是一个单向链表
type Node struct {
	next *Node
	list *List
	Value interface{}
}


func (n*Node)Next()*Node{
	if p := n.next; n.list != nil && p != &n.list.head {
		return p
	}
	return nil
}

type List struct {
	head Node 	// head.next ->第一个节点
	tail Node	// tail.next ->最后一个节点
	len  int
}


// 创建一个带头指针的队列
func (l *List) Init() *List {
	l.head.next = &l.head
	l.tail.next = &l.head
	l.len = 0
	return l
}

// 创建一个新的list
func New() *List { return new(List).Init() }

// 返回长度
func (l *List) Len() int { return l.len }

// 返回头结点
func (l *List) Front() *Node {
	if l.len == 0 {
		return nil
	}
	return l.head.next
}

// 返回尾结点
func (l *List) Back() *Node {
	if l.len == 0 {
		return nil
	}
	return l.tail.next
}

// 插入一个节点，在at节点之后
func (l *List) insert(n, at *Node) *Node {

	// 如果at节点是最后的尾结点
	if at ==l.tail.next{
		l.tail.next = n
	}
	// 将元素插入到 at 节点之后
	temp :=at.next
	at.next=n
	n.next=temp


	n.list=l
	l.len ++
	return n
}

// 在某某节点之后 插入一个值
func (l *List) insertValue(v interface{}, at *Node) *Node {
	return l.insert(&Node{Value: v}, at)
}


// 删除一个节点,单链表不太好处理：head->1->2->3->4<-tail  ，如果要删除2号节点，
// 传入也是2号节点，此时2号节点找不到前置节点，那么不通用 n.next ->n.next.next 公式
// 此时传入的节点只能是要删除节点的前置节点
func (l *List) remove(pre *Node) *Node {

	// 删除节点注意点
	// 传入节点是需要删除节点的父节点
	// 1.如果当前节点是tail 节点
	// 2.如果被删除节点是tail 节点
	// 3.删除节点的方式：node.next =node.next.next
	// 4.当需要删除的结点是头结点的话：head.next =head.next.next

	// 如果当前节点是尾结点
	if pre.list != l && pre!= l.tail.next{
		return nil
	}
	tem :=pre.next //被删除的节点
	pre.next= pre.next.next

	if tem == l.tail.next{ //如果被删除的节点是tail节点
		l.tail.next=pre
	}
	// 将尾结点指向最后一个
	l.len--
	return tem
}


// 在list 上面进行append操作
func (l*List)Append(node *Node)*Node{
	return l.insert(node,l.tail.next)
}

// 在list 进行append操作
func (l*List)AppendValue(v interface{})*Node{
	return l.insertValue(v,l.tail.next)
}

// 从头部弹出数据
func (l*List)Pop()*Node{
	return l.remove(&l.head)
}

// 查找父节点
func (l*List)findParent(node *Node)*Node{
	var n *Node
	n =l.head.next
	for{
		if n ==nil{
			return nil
		}
		if n.next!=nil && n.next ==node{
			return n
		}
		n= n.next
	}
}

// 删除节点 O(n)复杂度
func (l*List)RemoveNode (node *Node)*Node{
	p :=l.findParent(node)
	return l.remove(p)
}

// 节点反转,将链表进行反转 -就地反转
func (l *List)ReverseByStanding()*Node{
	// node

	var cur,pre *Node // pre 是第一位，cur 是当前的位置

	cur =l.head.next // cur =1

	for cur !=nil {
		pre =cur		// 1 pre =1 ,cur =1
		cur =cur.next 	// cur =2 ,
		cur.next=pre	// cur.next =1
	}

	return pre

}