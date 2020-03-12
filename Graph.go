package arithmetic

import "fmt"

type  VertexType string   //顶点类型
type EdgeType int		  //边上的权值
const Max=^uint16(0)
var boolBuff []bool
var quene []int           //宽度优先遍历需要用到的队列
type  Graph struct {
	vexs []VertexType 		//节点集合
	Edges [999][999]EdgeType		//边的邻接矩阵
	numVexs,numEdgs int		//顶点的个数和边的数量
}

func creatGraph(g *Graph){
	fmt.Println("输入你想要建立的图的节点的数量和边的数量")
	fmt.Scanf("%d%d",&g.numVexs,&g.numEdgs)
	fmt.Println("请输入每一个节点的数据")
	for i:=0;i<g.numVexs;i++{
		term:=""
		fmt.Scanf("%s",&term)
		g.vexs=append(g.vexs,VertexType(term))
	}
	for i:=0;i<g.numEdgs;i++{
		for j:=0;j<g.numEdgs;j++{
			if i==j{
				g.Edges[i][j]=EdgeType(0)
			}else{
				g.Edges[i][j]=EdgeType(Max)
			}
		}
	}
	fmt.Println(g.numEdgs)
	fmt.Println("请输入每一条边的数据,分别是起点，终点，长度")
	for i:=0;i<g.numEdgs;i++{
		n,m,k:=0,0,0
		fmt.Scanf("%d%d%d",&n,&m,&k)
		g.Edges[n][m]=EdgeType(k)
		//g.Edges[m][n]=EdgeType(k)、、
	}
	//DeBug代码，用于检验输入的数据是否正常
	fmt.Println("图的邻接矩阵")
	for i:=0;i<g.numVexs;i++{
		for j:=0;j<g.numVexs;j++{
			t:=g.Edges[i][j]
			if t==EdgeType(Max){
				t=-1
			}
			fmt.Printf("%d	",t)
		}
		fmt.Println()
	}
}
//深度优先遍历
func DFS(g *Graph,j int){//g是图，i是遍历的下标
	if !boolBuff[j]{
		boolBuff[j]=true
		fmt.Printf("%v->",g.vexs[j])
		for i:=0;i<g.numVexs;i++{
			if g.Edges[j][i]!=EdgeType(Max)&&!boolBuff[i]{
				DFS(g,i)
			}
		}
	}
}
//广度优先遍历
func BFS(g *Graph){//g是图，i是遍历的下标
	//入队
	for i:=0;i<g.numVexs;i++{//遍历每一个节点@@这一步是是为了兼容非连通图
		if !boolBuff[i]{
			boolBuff[i]=true
			fmt.Printf("->%v",g.vexs[i])
			quene=append(quene,i)
			for len(quene)!=0{
				x:=quene[0]
				quene=quene[1:len(quene)]//出队列
				for j:=0;j<g.numVexs;j++{
					if g.Edges[x][j]!=EdgeType(Max)&&!boolBuff[j]{
						boolBuff[j]=true
						fmt.Printf("->%v",g.vexs[j])
						quene=append(quene,j)
						//fmt.Println(quene)
					}
				}
			}
		}
	}
}
func RunGraph(){
	g:=new(Graph)
	creatGraph(g)
	for i:=0;i<g.numVexs;i++{
	boolBuff=append(boolBuff,false)
	}

	//图的深度优先遍历
	fmt.Println("深度优先遍历")
	for i:=0;i<g.numVexs;i++{
		DFS(g,i)
	}
	//图的广度优先遍历
	for i:=0;i<g.numVexs;i++{
		boolBuff[i]=false
	}
	fmt.Println()
	fmt.Println("广度优先遍历")
	BFS(g)
}
/*
测试用例的网站
https://blog.csdn.net/nrsc272420199/article/details/82810444
 */
/*
0 1 4
0 4 2
0 3 7
1 5 3
2 1 9
2 4 6
2 3 5
5 4 1
 */
 /*
 0
1
2
3
4
5
 */