package arithmetic

import "fmt"

//二叉排序树的定义以及一些实现算法

type bst struct {//二叉排序树的结构定义
	data int
	leftNode *bst
	rightNode *bst
}

func createBinarySortTree(t **bst){//创建二叉排序树
	fmt.Println("输入二叉排序树的数据,不要求顺序")
	var element int
	term:=0
	for {
		fmt.Scanf("%d",&element)
		if element==-1{
			break
		}
			//fmt.Printf("1:%v\n",(*t).data)
		if term==0{
			insertNode(t,element,true)
			term++
		}else{
			insertNode(t,element,false)
		}
			//fmt.Printf("2:%v\n",(*t).data)
	}
}

func insertNode(t **bst,element int,f bool)bool{//往二叉排序树里面插入
	if (*t)==nil||(f&&(*t).leftNode==nil&&(*t).rightNode==nil){
		tt:=new(bst)
		tt.data=element
		tt.rightNode=nil
		tt.leftNode=nil
		(*t)=tt
		return true
	}
	if (*t).data==element{
		return false
	}

	if (*t).data>element{
		//fmt.Println("ds")
		return insertNode((&((*t).leftNode)),element,false)
	}
	//fmt.Println("sd")
	return insertNode((&((*t).rightNode)),element,false)
}

func firstErgodicForBst(t *bst){//
	if t==nil{
		return
	}
	fmt.Printf("%d->",t.data)
	firstErgodicForBst(t.leftNode)
	firstErgodicForBst(t.rightNode)
}

func middleErgodicForBst(t *bst){
	if t==nil{
		return
	}
	middleErgodicForBst(t.leftNode)
	fmt.Printf("%d->",t.data)
	middleErgodicForBst(t.rightNode)
}

func laterErgodicForBst(t *bst){
	if t==nil{
		return
	}
	laterErgodicForBst(t.leftNode)
	laterErgodicForBst(t.rightNode)
	fmt.Printf("%d->",t.data)
}

func RunBinarySortTree(){
	t:=new(bst)
	createBinarySortTree(&t)
	firstErgodicForBst(t)
	fmt.Println("end")
	middleErgodicForBst(t)
	fmt.Println("end")
	laterErgodicForBst(t)
	fmt.Println("end")
}