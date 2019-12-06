package arithmetic

import "fmt"

//哈夫曼树的构造以及特性

type huffumanTree struct {
	weight int
	parent,left,right int//父节点，左右节点的下标
}

func getSelect(buff []huffumanTree,end int)(int,int){
	s1,s2:=0,0
	min1,min2:=0,0//最小两个数的下标
	i:=0
	for buff[i].parent!=-1&&i<=end{
		i++
	}
	min1=buff[i].weight
	s1=i
	i++
	for buff[i].parent!=-1&&i<=end{
		i++
	}
	if buff[i].weight<min1{//第二个节点的权值比第一个小
		min2=min1
		s2=s1
		s1=i
		min1=buff[i].weight
	}else{
		min2=buff[i].weight
		s2=i
	}
	for j:=i+1;j<=end;j++{//遍历之后的找出最小的两个没有父节点的两个节点
		//如果有父节点的话，直接跳过
		if buff[j].parent!=-1{
			continue
		}
		//如果比最小的还小
		if buff[j].weight<min1{
			min2=min1
			s2=s1
			s1=j
			min1=buff[j].weight
		}else if buff[j].weight>=min1&&buff[j].weight<min2{
			min2=buff[j].weight
			s2=j
		}

	}
	return s1,s2
}

func hffumanTreeShow(start int,buff []huffumanTree){//哈夫曼树的（先序遍历，中序遍历，后序遍历其实否是一样的）
	if start<0{
		return
	}
	hffumanTreeShow(buff[start].left,buff)
	hffumanTreeShow(buff[start].right,buff)
	if buff[start].left==-1&&buff[start].right==-1{
		fmt.Printf("->%d",buff[start].weight)
		return
	}
}
func RunHffumanTree(){
	fmt.Println("输入哈弗曼树的节点的个数，以及每一个节点的权值")
	n:=0
	fmt.Scanf("%d",&n)
	huffBuff:=[]huffumanTree{}
	//初始化哈夫曼树的数据
	for i:=0;i<(2*n-1);i++{//给每一个叶子节点赋初始值,不存在单叶子节点的二叉树，所以一共有2*n-1个节点
		m:=0
		if i<=(n-1){//叶子节点需要给予权重
			fmt.Scanf("%d",&m)
		}
		t:=new(huffumanTree)
		t.weight=m
		t.parent=-1
		t.left=-1
		t.right=-1
		huffBuff=append(huffBuff,*t)
	}
	//构造哈夫曼树
	for i:=n;i<(2*n-1);i++{//之后这里的每一个节点都是新构建的节点
		x,y:=getSelect(huffBuff,i-1)
		huffBuff[i].weight=huffBuff[x].weight+huffBuff[y].weight
		huffBuff[i].left=x
		huffBuff[i].right=y
		huffBuff[x].parent,huffBuff[y].parent=i,i
	}
	fmt.Println(huffBuff)
	hffumanTreeShow((2*n-2),huffBuff)
}
