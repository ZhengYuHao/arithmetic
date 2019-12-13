package Pat

import (
	"fmt"
)

//The highest building in our city has only one elevator. A request list is made up with N positive numbers. The numbers denote at which floors the elevator will stop, in specified order. It costs 6 seconds to move the elevator up one floor, and 4 seconds to move down one floor. The elevator will stay for 5 seconds at each stop.
//
//For a given request list, you are to compute the total time spent to fulfill the requests on the list. The elevator is on the 0th floor at the beginning and does not have to return to the ground floor when the requests are fulfilled.
//
//Input Specification:
//
//Each input file contains one test case. Each case contains a positive integer N, followed by N positive numbers. All the numbers in the input are less than 100.
//
//Output Specification:
//
//For each test case, print the total time on a single line.
//
//Sample Input:
//
//3 2 3 1
//Sample Output:
//
//41

//这题主要考察写程序的灵活思路

func Elevator(){//电梯调度时间统计
	n:=0
	sum:=0//电梯运行一共花费的时间
	fmt.Scanf("%d",&n)
	v:=make([]int,n+1)
	v[0]=0
	for i:=1;i<=n;i++{
		fmt.Scanf("%d",&v[i])
	}
	for i:=0;i<n;i++{
		//fmt.Println(sum)
		term:=v[i]-v[i+1]
		if term<0{//电梯向上
			sum+=(6*(-term)+5)
		}else if term>0{//电梯向下
			sum+=(4*(term)+5)
		}else if term==0{//电梯不动
			sum+=(4*(term)+5)
		}
	}
	fmt.Printf("%d",sum)

}