package arithmetic

import "fmt"

const Inf  =int32(^uint16(0))//无穷大数
var mapTwoDimension =[500][500]int32{}//用于保存边权的数组
var mapVist=[500]int32{}//用于保存该节点是否已经遍历
var mapDist=[500]int32{}//用于保存起点到该节点的最短路径
var n,m int32//n个点，m个边
func dijkstra()  {
	//for i,j:=range mapTwoDimension{
	//	fmt.Println(i,j)
	//}
	for i:=int32(0);i<n;i++{//初始化最短路径
		mapDist[i]=mapTwoDimension[0][i]
	}
	//开始进行最短路径查找
	for i:=int32(0);i<n;i++{//一共进行n次遍历

		minL:=Inf//当前最短路径
		temp:=int32(0)//temp节点

		for j:=int32(0);j<n;j++{//找到距离原点最近的点
			if mapVist[j]==0&&mapDist[j]<minL{
				minL=mapDist[j]
				temp=j
			}
		}
		mapVist[temp]=1
		for k:=int32(0);k<=n;k++{
			if mapTwoDimension[temp][k]+mapDist[temp]<mapDist[k]{
				mapDist[k]=mapTwoDimension[temp][k]+mapDist[temp]
			}
		}
	}

}

func InitMap(){//边图的初始化以及用户键入边信息
	//边信息的初始化
	for i:=int32(0);i<n;i++{
		mapTwoDimension[i][i]=0
		for x:=int32(0);x<n;x++{
			if i!=x{
				mapTwoDimension[i][x]=Inf
			}
		}
	}
	//键入边的信息
	var u,v,w int32
	for i:=int32(0);i<m;i++{
		fmt.Scanf("%d%d%d",&u,&v,&w)
		if mapTwoDimension[u][v]>w{
			mapTwoDimension[u][v]=w
			mapTwoDimension[v][u]=w
		}
	}

}
func Dijkstra()  {

	//提示输入n个节点，m条边。
	fmt.Println("输入节点数和边数")
	fmt.Scanf("%d%d",&n,&m)
	InitMap()
	dijkstra()//调用djkstra算法
	for i:=int32(0);i<n;i++{
		fmt.Printf("起点到%d节点的最短路径为:%v\n",i,mapDist[i])
	}
}


//测试用例
//6 8
//0 1 2
//0 2 3
//0 3 6
//1 4 4
//1 5 6
//2 3 2
//3 4 1
//3 5 3
// https://www.cnblogs.com/bigsai/p/11537975.html