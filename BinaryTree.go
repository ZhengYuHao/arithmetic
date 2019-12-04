package arithmetic

import (
	"fmt"
)

type binarytree struct {//二叉树结构定义
	data int
	leftNode *binarytree
	rightNode *binarytree
}

func createBinaryTree()*binarytree{//二叉树的创建
	data:=0
	fmt.Scanf("%d",&data)
	if data==-1{
		return nil
	}else{
		var term=new(binarytree)
		term.data=data
		term.leftNode=createBinaryTree()
		term.rightNode=createBinaryTree()
		return term
	}
}

func firstErgodic(t *binarytree){//
	if t==nil{
		return
	}
	fmt.Printf("->%d",t.data)
	firstErgodic(t.leftNode)
	firstErgodic(t.rightNode)
}

func middleErgodic(t *binarytree){
	if t==nil{
		return
	}
	middleErgodic(t.leftNode)
	fmt.Printf("->%d",t.data)
	middleErgodic(t.rightNode)
}

func laterErgodic(t *binarytree){
	if t==nil{
		return
	}
	laterErgodic(t.leftNode)
	laterErgodic(t.rightNode)
	fmt.Printf("->%d",t.data)
}

func RunBinaryTree(){
	fmt.Println("先序创建二叉树,#表示叶子节点")
	head:=createBinaryTree()
	fmt.Println("该二叉树的先序遍历为:")
	firstErgodic(head)
	fmt.Println("end")
	fmt.Println("该二叉树的中序遍历为:")
	middleErgodic(head)
	fmt.Println("end")
	fmt.Println("该二叉树的后序遍历为:")
	laterErgodic(head)
	fmt.Println("end")
}