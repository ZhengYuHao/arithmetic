package Pat


import (
"fmt"
)
//本题主要考察二叉树的结构已经数的遍历方法之间的联系技巧
type bTree struct {
	data int32
	lNode *bTree
	rNode *bTree
}
var inOrder []int32//中序遍历的数组
var postOrder []int32//逆序遍历数组
var nodeNum int32   //节点个数
func dfs(lInoder int32,rInorder int32,lPostOrder int32,rPostOrder int32)*bTree{
	if lInoder>rInorder{
		return nil
	}
	k:=int32(0)
	tempNode:=&bTree{}
	tempNode.data=postOrder[rPostOrder]//当前节点是逆序的最后一个节点
	for k=lInoder;k<=rInorder;k++{//遍历中序节点,取得index的偏移量
		if inOrder[k]==tempNode.data{
			break
		}
	}
	numLeft:=k-lInoder
	tempNode.lNode=dfs(lInoder,k-1,lPostOrder,lPostOrder+numLeft-1)
	tempNode.rNode=dfs(k+1,rInorder,lPostOrder+numLeft,rPostOrder-1)
	return tempNode
}
func bfs(treeNode *bTree){
	indexBtree:=treeNode
	termBtree:=make([]*bTree,0)
	termBtree=append(termBtree,indexBtree)//入列
	num:=int32(0)
	for len(termBtree)!=0{
		tt:=termBtree[0]
		fmt.Printf("%d",tt.data)
		termBtree=append(termBtree[0:0],termBtree[1:]...)
		num++
		if num<nodeNum{
			fmt.Printf(" ")
		}
		if tt.lNode!=nil{
			termBtree=append(termBtree,tt.lNode)
		}
		if tt.rNode!=nil{
			termBtree=append(termBtree,tt.rNode)
		}
	}
}
func main(){
	//初始化Start
	fmt.Scanf("%d",&nodeNum)
	for i:=int32(0);i<nodeNum;i++{
		temp:=int32(0)
		fmt.Scanf("%d",&temp)
		postOrder=append(postOrder,temp)
	}
	for i:=int32(0);i<nodeNum;i++{
		temp:=int32(0)
		fmt.Scanf("%d",&temp)
		inOrder=append(inOrder,temp)
	}
	//End
	//根据数遍历规则，构建树Start
	root:=dfs(0,nodeNum-1,0,nodeNum-1)
	//End
	//层次遍历
	bfs(root)
}