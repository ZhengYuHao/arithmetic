package arithmetic

import "fmt"
//小顶堆和大顶堆的区别主要在关键判断领域的不同
var dataList=make([]int,0)//初始化一个空数组，用来存放堆数据

func getParentIndex(index int)int{//获取父节点的下标,返回-1表示出错
	if index==0||index>len(dataList)-1{
		return -1
	}
	return (index-1)/2
}

func swap(indexA,indexB int){//交换两个下标代表的值
	dataList[indexA],dataList[indexB]=dataList[indexB],dataList[indexA]
}

func insert(data int){//往堆里面插入数据
	//思路是往数组的末尾添加数据，然后依次往前堆化
	dataList=append(dataList,data)
	index:=len(dataList)-1//末尾字段的index
	parent:=getParentIndex(index)//父节点的index
	for parent!=-1&&dataList[parent]<dataList[index]{//大顶堆的处理，父节点一定要大于子节点
		swap(index,parent)
		index=parent
		parent=getParentIndex(parent)
	}
}

func heapify(index int){//函数用于下沉构建堆，这里的思想很重要，这里一定要理解到位
	totalIndex:=len(dataList)-1//最终节点的坐标
	for{
		maxValueIndex:=index
		if (index*2)+1<=totalIndex&&dataList[index*2+1]>dataList[maxValueIndex]{
			maxValueIndex=index*2+1
		}
		if (index*2)+2<=totalIndex&&dataList[index*2+2]>dataList[maxValueIndex]{
			maxValueIndex=index*2+2
		}
		if maxValueIndex==index{
			break
		}
		swap(maxValueIndex,index)
		index=maxValueIndex
	}
}
func removeMax()int{//删除堆顶元素，然后将最后一个元素放到堆顶，再从上到下依次堆化
	removeData:=dataList[0]
	dataList[0]=dataList[len(dataList)-1]
	dataList=dataList[0:len(dataList)-1]
	heapify(0)
	return removeData
}

func RunHeap(){
	n:=0
	fmt.Println("输入一个整数，系统将建立一个堆，堆元素包含1~N")
	fmt.Scanf("%d",&n)
	for i:=1;i<=n;i++{//建堆
		insert(i)
	}
	fmt.Println(dataList)
	//删除堆顶元素
	for i:=1;i<=n;i++{
		j:=removeMax()
		if i!=n{
			fmt.Printf("%d->",j)
		}else{
			fmt.Printf("%d--End.",j)
		}

	}
}