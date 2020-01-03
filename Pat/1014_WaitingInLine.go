package Pat
import "fmt"
type Customer struct{
	inTime int32    //进入队列时间
	outTime int32	//出队列时间
	costTime int32	//花费时间
}
var quere [21][]int32//最多10个窗口，每个窗口最多20个人在黄线内
func WaitingInLine()  {
	n,m,k,q:=0,0,0,0
	var customers []Customer
	customers=append(customers,*new(Customer))
	customerNow:=int32(1)//第n*m+1个玩家
	fmt.Scanf("%d%d%d%d",&n,&m,&k,&q)
	for i:=1;i<=k;i++{//该循环输入每一个玩家的信息
		term:=new(Customer)
		t:=int32(0)
		fmt.Scanf("%d",&t)
		term.inTime=-1//-1表示还没进入计算
		term.costTime=t
		customers=append(customers,*term)
		if i%m==0 {//将数据入列
			if len(quere[m])<m{
				quere[m]=append(quere[m],int32(i))
				customerNow++
			}
		} else {
			if len(quere[i%m])<m{
				quere[i%m]=append(quere[i%m],int32(i))
				customerNow++
			}
		}
	}
	query:=[]int32{}
	for i:=0;i<q;i++{
		term:=int32(0)
		fmt.Scanf("%d",&term)
		query=append(query,term)
	}
	for nowTime:=int32(0);nowTime<=540;nowTime++{//从8点到晚上17点，一共540秒
		for quereIndex:=1;quereIndex<=n;quereIndex++{//从左到右遍历队列
			if len(quere[quereIndex])>0{//队列有东西的话
				term:=quere[quereIndex][0]//队列最前面的那个人
				if customers[term].inTime==-1&&nowTime<540{//该玩家还没进行处理
					customers[term].inTime=nowTime
				}else if customers[term].inTime!=-1&&nowTime-customers[term].inTime==customers[term].costTime{//满足出列,出列一个人，如果后面还有人的话，直接入列一个人
					customers[term].outTime=nowTime
					quere[quereIndex]=quere[quereIndex][1:len(quere[quereIndex])]
					if customerNow<int32(k){
						quere[quereIndex]=append(quere[quereIndex],customerNow)
						customerNow++
					}
					if len(quere[quereIndex])>0&&nowTime<540{
						first:=quere[quereIndex][0]
						customers[first].inTime=nowTime
					}
				}else if customers[term].inTime!=-1&&nowTime-customers[term].inTime<customers[term].costTime{//不满足出列
					continue
				}
			}
		}
	}
	for i:=0;i<q;i++{
		if customers[query[i]].inTime!=-1{
			if customers[query[i]].outTime==0{//到了5点还没被服务完的
				customers[query[i]].outTime=customers[query[i]].inTime+customers[query[i]].costTime
			}
			t1,t2:=int32(0),int32(0)
			t1=customers[query[i]].outTime/60
			t2=customers[query[i]].outTime%60
			fmt.Printf("%02d:%02d\n",8+t1,t2)
		}else{
			fmt.Println("Sorry")
		}
	}
}
//2 2 7 5
//1 2 6 4 3 534 2
//3 4 5 6 7

