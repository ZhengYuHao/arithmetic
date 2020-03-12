package Pat


import (
"fmt"
)

var cMax,n,sp,m int32
var e   [510][510]int32//临接矩阵存储图的边信息
var dis [510]int32  //节点信息
var weight [510]int32 //每一个节点还需要的数据量，正数表示多余的，负数表示缺少的
var visti [510]bool
var pre [510][]int32   //存储先驱节点
var path,temppath []int32
const Inf  =99999999
var minNeed,minBack=int32(Inf),int32(Inf)
func main(){//主函数
	//初始化图结构，数据Inf表示点与点之间不可达
	for i:=0;i<510;i++{
		for j:=0;j<510;j++{
			e[i][j]=Inf
		}
	}
	//初始化节点信息
	for i:=0;i<510;i++{
		dis[i]=Inf
	}
	//获取玩家输入的第一行数据，分别是每个节点最多自行车数量，节点个数，问题节点，边的条数
	fmt.Scanf("%d%d%d%d",&cMax,&n,&sp,&m)
	//输入第二行数据，分别是每个节点目前自行车的数量
	for i:=int32(1);i<=n;i++{//遍历每一个节点，输入数据
		fmt.Scanf("%d",&weight[i])
		weight[i]=weight[i]-cMax/2
	}
	//输入第三部分数据，边信息。分别是边起点，边终点，边长度
	for i:=int32(0);i<m;i++{
		x1,x2,x3:=int32(0),int32(0),int32(0)
		fmt.Scanf("%d%d%d",&x1,&x2,&x3)
		e[x1][x2]=x3
		e[x2][x1]=x3
	}
	//开始使用dijkstra算法求最短路径Start
	dis[0]=0

	for k:=int32(0);k<=n;k++{//进行n次匹配
		termIndex:=int32(-1)
		termMin:=int32(Inf)
		for j:=int32(0);j<=n;j++{//每一次遍历所有节点，找出和源点最小距离的点
			if dis[j]<termMin&&visti[j]==false{
				termIndex=j
				termMin=dis[j]
			}
		}
		visti[termIndex]=true
		if termIndex==-1{
			break
		}
		//下面就是松弛代码
		for i:=int32(0);i<=n;i++{
			if visti[i]==false&&e[termIndex][i]!=Inf{
				if dis[termIndex]+e[termIndex][i]<dis[i]{
					dis[i]=dis[termIndex]+e[termIndex][i]
					pre[i]=[]int32{}
					pre[i]=append(pre[i],termIndex)
				}else if dis[termIndex]+e[termIndex][i]==dis[i]{
					pre[i]=append(pre[i],termIndex)
				}
			}
		}
	}
	//End
	//通过深度优先遍历遍历有最短路径构建的结构
	dfs(sp)
	fmt.Printf("%d 0",minNeed)
	for i:=len(path)-2;i>=0;i--{
		fmt.Printf("->%d",path[i])
	}
	fmt.Printf(" %d",minBack)
}

func dfs(v int32){
	temppath=append(temppath,v)//堆栈，有堆栈就有出栈
	if v==0{//如果是源点,这也是递归终止条件
		need,back:=int32(0),int32(0)//当前最短路径需要的自行车数量和自行车返回数量
		for i:=len(temppath)-1;i>=0;i--{//反向遍历，从0源点往下遍历
			id:=temppath[i]//当前节点
			if weight[id]>0{
				back+=weight[id]
			}else{
				if back>(0-weight[id]){//上面节点搜集到的数量大于当前节点需要的数量
					back+=weight[id]
				}else{
					need+=((0-weight[id])-back)
					back=0
				}
			}

		}
		if need<minNeed{
			minNeed=need
			minBack=back
			path=temppath
		}else if (need==minNeed&&back<minBack){
			minBack=back
			path=temppath
		}
		temppath=temppath[0:len(temppath)-1]
		return
	}
	for i:=0;i<len(pre[v]);i++{//遍历所有的前缀节点，这里有多个数据其实就是有多条最短路径的意思
		dfs(int32(pre[v][i]))
	}
	temppath=temppath[0:len(temppath)-1]
}