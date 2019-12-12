package Pat

import (
	"fmt"
)

//要求统计树每一层叶子节点的个数
//1.树存储的数据结构(这个题目可以使用孩子表示法)
const MaxLimit=int(102)//节点最大数
type treeNode struct {
	data int
	next *treeNode
}
var tree [MaxLimit]treeNode
var layerNum [MaxLimit]int
var layerMax int
func CountingLeaves(){//统计每一层叶子节点的个数
	n,m:=0,0//节点总数和非叶子节点数
	id,k,idChild:=0,0,0
	fmt.Scanf("%d%d",&n,&m)
	for i:=1;i<=n;i++{
		tree[i].data=i
		tree[i].next=nil
	}
	for i:=1;i<=m;i++{//输入m行非叶子节点和他的子节点信息

		fmt.Scanf("%d%d",&id,&k)
		termIdex:=&tree[id]
		for j:=1;j<=k;j++{
			fmt.Scanf("%d",&idChild)
			//fmt.Println(idChild)
			term:=new(treeNode)
			term.data=idChild
			term.next=nil
			termIdex.next=term
			termIdex=termIdex.next
		}
	}
	//到了这一步的话，树是按照输入构建完成了
	//for i:=1;i<=n;i++{
	//	term:=&tree[i]
	//	for term!=nil{
	//		fmt.Printf("-->:%v",term.data)
	//		term=term.next
	//	}
	//	fmt.Println()
	//}
	DFS(1,1)
	for i:=1;i<=layerMax;i++{
		if i!=layerMax{
			fmt.Printf("%d ",layerNum[i])
		}else{
			fmt.Printf("%d",layerNum[i])
		}

	}
	//for i:=1;i<=layerMax;i++{
	//	if i!=1 {
	//		fmt.Printf(" ")
	//	}
	//	fmt.Printf("%d",layerNum[i])
	//}
}
//树的深度优先遍历
func DFS(index int,layer int){
	term:=&tree[index]
	if layerMax<layer{
		layerMax=layer
	}
	if term.next==nil{//递归终止条件
		layerNum[layer]+=1
		return
	}
	for term.next!=nil{
		if term.next!=nil{
			DFS(term.next.data,layer+1)
		}
		term=term.next
	}
}
//10 4
//1 3 2 3 4
//2 2 5 6
//4 1 7
//7 3 8 9 10
//http://c.biancheng.net/view/3395.html