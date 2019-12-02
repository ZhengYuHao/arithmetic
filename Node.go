package arithmetic

import "fmt"

type node struct {//定义树结构
	data int32
	next *node
}

func NodeTest(){
	l:=int32(0)
	fmt.Println("输入一个数字用于构建链表的长度" )
	fmt.Scanf("%d",&l)
	head:=&node{}//保存头指针
	term:=head
	//1:头插法建立
	for i:=int32(0);i<l;i++{//尾插法建立一个链表(带头结点的链表)
		t:=new(node)
		t.data=i
		term.next=t
		term=t
	}

	term.next=nil
	term=head
	for{
		if term.next!=nil{
			fmt.Printf("%d->",term.next.data)
			term=term.next
		}else{
			fmt.Printf("end")
			break
		}
	}
	//头插法建立链表
	//l:=int32(0)
	//fmt.Println("输入一个数字用于构建链表的长度" )
	//fmt.Scanf("%d",&l)
	//head:=&node{}//保存头指针
	////1:头插法建立
	//for i:=int32(0);i<l;i++{//头插法建立一个链表(带头结点的链表)
	//	t:=new(node)
	//	t.data=i
	//	if i==0{
	//		t.next=nil
	//	}else{
	//		t.next=head.next
	//	}
	//	head.next=t
	//}
	//term:=head
	//for{
	//	if term.next!=nil{
	//		fmt.Printf("%d->",term.next.data)
	//		term=term.next
	//	}else{
	//		fmt.Printf("end")
	//		break
	//	}
	//}

}