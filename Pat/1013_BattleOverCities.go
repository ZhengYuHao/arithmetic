package Pat

import (
	"fmt"
)

//1013 Battle Over Cities，实际上就是求一个连通图，去掉某一个节点之后的，联通模块
var Graph [1020][1020]int32//图的邻接矩阵
var TermGraph [1020][1020]int32
var GraphVisit [1020]bool //节点是否被遍历
var TermGraphVisit [1020]bool
var n,m,k int
var termll int

func BattleOverCities()  {

	fmt.Scanf("%d%d%d",&n,&m,&k)
	var kBuffer []int
	for i:=0;i<m;i++{//构建图的邻接矩阵
		t1,t2:=0,0
		fmt.Scanf("%d%d",&t1,&t2)
		Graph[t1][t2]=1
		Graph[t2][t1]=1
	}
	for i:=1;i<=n;i++{
		Graph[i][i]=1
	}
	for i:=0;i<k;i++{//需要处理的k值
		t:=0
		fmt.Scanf("%d",&t)
		kBuffer=append(kBuffer,t)
	}
	//fmt.Println(kBuffer)
	for _,j:=range kBuffer{
		termll=j
		//fmt.Println(Graph)
		TermGraph=Graph
		GraphVisit=TermGraphVisit
		term:=0
		for x:=1;x<=n;x++{//以每一个节点为起点，进行遍历
			if x==termll{
				continue
			}
			if !GraphVisit[x]{
				//fmt.Printf("SS:%v",x)
				DFS(x)
				term++
			}
			//这个循环体里面进行遍历
		}
		fmt.Println(term-1)
	}

}

func DFS(nowVisit int){
	if TermGraphVisit[nowVisit]==true{
		return
	}else{//如果这个节点没有被遍历
		for i:=1;i<=n;i++{
			if i==termll{
				continue
			}
			if  TermGraph[nowVisit][i]==1&&GraphVisit[i]==false{
				//fmt.Printf("t:%v\n",i)
				GraphVisit[i]=true
				DFS(i)
			}
		}
	}

}