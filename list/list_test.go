package list

import (
	"fmt"
	"testing"
)

func TestNew(t *testing.T) {
	l :=New()
	fmt.Printf("这里是第一个节点地址：%p\n",l.AppendValue(1))
	fmt.Printf("这里是第二个节点地址：%p\n",l.AppendValue(2))
	fmt.Printf("这里是第二个节点地址：%p\n",l.AppendValue(3))

	fmt.Printf("1 :%+v ，地址：%p\n",l.head.next,l.head.next)// 这里是第一个数
	fmt.Printf("2 :%+v ，地址：%p\n",l.head.next.next,l.head.next.next)
	fmt.Printf("2 :%+v ，地址：%p\n",l.head.next.next.next,l.head.next.next.next)
	fmt.Printf("头部 :%+v 地址：%p\n",l.head,&l.head)
	fmt.Printf("尾部 :%+v 地址：%p\n",l.tail,&l.tail)

	node :=l.ReverseByStanding()
	fmt.Printf("1 :%+v ，地址：%p\n",node,node)// 这里是第一个数

}