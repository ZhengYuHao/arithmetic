package Pat

import "fmt"

const Inf =int(^uint32(0))
const MaxCity  =501//城市最大的节点
var Side=[MaxCity][MaxCity]int{}//边权信息
var Weight=[MaxCity]int{}//点权信息
var Router=[MaxCity]int{}//起点到每一个节点的最短路径
var NumRouter=[MaxCity]int{}//起点到每一个节点最短路径的条数
var AllWeight=[MaxCity]int{}//起点到每一个节点的点权总和
var Visit=[MaxCity]bool{}//节点是否遍历
func Emergency(){
	n,m,c1,c2:=0,0,0,0
	//先对数据结构进行初始化
	fmt.Scanf("%d%d%d%d",&n,&m,&c1,&c2)
	//点权的初始化
	for i:=0;i<n;i++{
		fmt.Scanf("%d",&Weight[i])
	}
	//边权的初始花
	for i:=0;i<n;i++{
		for j:=0;j<n;j++{
			if i==j{
				Side[i][j]=0
			}else{
				Side[j][i]=Inf
				Side[i][j]=Inf
			}
		}
	}
	for i:=0;i<m;i++{
		x,y,z:=0,0,0
		fmt.Scanf("%d%d%d",&x,&y,&z)
		Side[x][y]=z
		Side[y][x]=z
	}
	//最短路径初始化
	for i:=0;i<n;i++{
		Router[i]=Side[c1][i]
	}
	//最短路径条数初始化
	NumRouter[c1]=1
	AllWeight[c1]=Weight[c1]
	//开始使用dijikstra算法
	for i:=0;i<n;i++{//对每一个节点进行检测
		minL:=Inf//当最短路径的长度
		term:=-1//当前最短路径的最后一个节点
		for j:=0;j<n;j++{
			if Router[j]<minL&&!Visit[j]{
				minL=Router[j]
				term=j
			}
		}
		if term==-1{
			break
		}
		Visit[term]=true
		for k:=0;k<n;k++{
			if !Visit[k]&&Side[term][k]!=Inf{
				if Router[term]+Side[term][k]<Router[k]{
					Router[k]=Router[term]+Side[term][k]
					NumRouter[k]=NumRouter[term]
					AllWeight[k]=AllWeight[term]+Weight[k]
				}else if Router[term]+Side[term][k]==Router[k]{
					NumRouter[k]+=NumRouter[term]
					if AllWeight[term]+Weight[k]>AllWeight[k]{
						AllWeight[k]=AllWeight[term]+Weight[k]
					}
				}
			}

		}
	}
	fmt.Printf("%d %d",NumRouter[c2],AllWeight[c2])
}