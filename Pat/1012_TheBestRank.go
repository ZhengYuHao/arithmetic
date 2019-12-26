package Pat

import (
	"fmt"
	"sort"
)

type  Student struct{
	c int  //c语言的成绩
	m int  //数学的成绩
	e int  //英语的成绩
	a int  //平均分的成绩
	sel [4]int//每一项的排名
}

func TheBestRank(){
	students:=make(map[string]Student,0)
	//searchStudents:=make([]string,0)
	n,m:=0,0
	dataSort:=[4][]int{}
	fmt.Scanf("%d%d",&n,&m)
	for i:=0;i<n;i++{
		d1,d2,d3,d4:="",0,0,0
		fmt.Scanf("%s%d%d%d",&d1,&d2,&d3,&d4)
		term:=new(Student)
		term.c=d2
		term.m=d3
		term.e=d4
		//平均分的四舍五入处理
		t1:=(d2+d3+d4)/3
		t2:=float64(d2+d3+d4)/3.0
		if t2-0.5>float64(t1){
			term.a=t1+1
		}else if t2-0.5<float64(t1){
			term.a=t1
		}
		students[d1]=*term
		dataSort[1]=append(dataSort[1],d2)//C语言
		dataSort[2]=append(dataSort[2],d3)//数学
		dataSort[3]=append(dataSort[3],d4)//英语
		dataSort[0]=append(dataSort[0],term.a)//平均分
	}
	//fmt.Println(dataSort)
	//然后对每一个类进行去重排序，然后下标就是排名
	sort.Ints(dataSort[0])
	sort.Ints(dataSort[1])
	sort.Ints(dataSort[2])
	sort.Ints(dataSort[3])
	//排完序之后，去除相同元素
	for i:=0;i<4;i++{
		tt:=make(map[int]int,0)
		for _,x:=range dataSort[i]{
			tt[x]++
		}
		dataSort[i]=[]int{}
		for k,_:=range tt{
			dataSort[i]=append(dataSort[i],k)
		}
	}
	//由于map的无序性，之后还是要重新排序
	sort.Sort(sort.Reverse(sort.IntSlice(dataSort[0])))
	sort.Sort(sort.Reverse(sort.IntSlice(dataSort[1])))
	sort.Sort(sort.Reverse(sort.IntSlice(dataSort[2])))
	sort.Sort(sort.Reverse(sort.IntSlice(dataSort[3])))

	//sort.Ints(dataSort[0])
	//sort.Ints(dataSort[1])
	//sort.Ints(dataSort[2])
	//sort.Ints(dataSort[3])
	//fmt.Println(dataSort)
	//fmt.Println(students)
	dataMap1:=make(map[int]int,0)
	dataMap2:=make(map[int]int,0)
	dataMap3:=make(map[int]int,0)
	dataMap4:=make(map[int]int,0)
	for i,j:=range dataSort{
		for x,y:=range j{
			switch i {
			case 0:
				dataMap1[y]=x+1
			case 1:
				dataMap2[y]=x+1
			case 2:
				dataMap3[y]=x+1
			case 3:
				dataMap4[y]=x+1
			}

		}
	}
	//fmt.Println(dataMap1)
	//fmt.Println(dataMap2)
	//fmt.Println(dataMap3)
	//fmt.Println(dataMap4)
	for i,j:=range students{
		//fmt.Println(j)
		j.sel[0]=dataMap1[j.a]
		j.sel[1]=dataMap2[j.c]
		j.sel[2]=dataMap3[j.m]
		j.sel[3]=dataMap4[j.e]
		//fmt.Println(j)
		students[i]=j
	}
	//fmt.Println(students)
	studentSearchList:=make([]string,0)
	for i:=0;i<m;i++{
		term:="n"
		fmt.Scanf("%s",&term)
		studentSearchList=append(studentSearchList,term)
	}
	for _,j:=range studentSearchList{//遍历每一个想要知道自己成绩的人的总表
		_,ok:= students[j]
		if ok{
			//fmt.Println(students[j].sel)
			min,index:=GetMinAndIndex(students[j].sel)
			switch index {
			case 0:
				fmt.Printf("%d %s\n",min,"A")
			case 1:
				fmt.Printf("%d %s\n",min,"C")
			case 2:
				fmt.Printf("%d %s\n",min,"M")
			case 3:
				fmt.Printf("%d %s\n",min,"E")
			}
		}else{
			fmt.Printf("N/A\n")
		}

	}

}

func GetMinAndIndex(list [4]int)(int,int){
	min,index:=list[0],0
	for i,j:=range list{
		if j<min{
			min=j
			index=i
		}
	}
	return min,index
}
//5 6
//310101 98 85 88
//310102 70 95 88
//310103 82 87 94
//310104 91 91 91
//310105 85 90 90