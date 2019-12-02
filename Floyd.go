package arithmetic

import "fmt"

//弗洛伊德算法，运用动态规划的思想解决多元最短路径的问题

var G=[999][999]int32{}//图邻接矩阵
var path=[999][999]int32{}
var Floyd_n,Floyd_m int32
func initFloyd(){
	for i:=int32(0);i<Floyd_n;i++{
		for j:=int32(0);j<Floyd_n;j++{
			G[i][j]=Inf
			path[i][j]=j
		}
}
}
func floyd(){
	for i:=int32(0);i<int32(Floyd_n);i++{//每一个节点都作为中转节点
		for j:=int32(0);j<int32(Floyd_n);j++{//每一个节点都作为起始节点
			for k:=int32(0);k<int32(Floyd_n);k++{//每一个节点都作为最终节点
				if G[j][i]+G[i][k]<G[j][k]{
					G[j][k]=G[j][i]+G[i][k]
					path[j][k]=i
				}
			}
		}
	}

}
func showRecursion(path [999][999]int32,srcNode int32,disNode int32){//递归打印输出
	if path[srcNode][disNode]!=disNode{
		fmt.Printf("%d->",path[srcNode][disNode])
		showRecursion(path,path[srcNode][disNode],disNode)
	}else{
		fmt.Printf("%d\n",disNode)
	}
}
func Floyd(){
	fmt.Println("请输入图的节点数和边数，用于初始化图的信息")
	fmt.Scanf("%d%d",&Floyd_n,&Floyd_m)
	fmt.Println("请输入图的每一天边的信息")
	initFloyd()
	for i:=int32(0);i<Floyd_m;i++{
		n,m,w:=0,0,0
		fmt.Scanf("%d%d%d",&n,&m,&w)//这里输入每一个边的信息
		if int32(w)<G[n][m]&&w>=0{
			G[n][m]=int32(w)
		}
	}
	//调用floyd算法
	floyd()
	//输出经过floyd算法处理过的最短路径图
	fmt.Println("floyd算法已经进行了处理，请输入起始节点和最终节点:")
	var s,d int32
	fmt.Scanf("%d%d",&s,&d)
	fmt.Printf("%d->",s)
	showRecursion(path,s,d)
}