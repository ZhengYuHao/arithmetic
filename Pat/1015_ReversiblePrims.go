package Pat

import (
	"fmt"
	"math"
)
//这道题目考察的其实就是进制转换（10进制转换成其他进制，其他进制取反之后，转换成10进制）
//10进制转换成其他进制使用短除法
//其他进制转换成10进制使用用权相加法
func ReversiblePrims(){
	//数据结构的设计
	var datainput [][2]int
	temr1,temr2:=0,0
	//数据录入
	fmt.Scanf("%d",&temr1)
	fmt.Scanf("%d",&temr2)
	for temr1>=0{
		datainput=append(datainput,[2]int{temr1,temr2})
		fmt.Scanf("%d",&temr1)
		if temr1<0{
			break
		}
		fmt.Scanf("%d",&temr2)
	}
	//数据的处理
	for _,j:=range datainput{
		term:=[]int{}
		n,radix:=j[0],j[1]
		if !IsPrime(n){//before不是素数的话
			fmt.Println("No")
			continue
		}else{
			for{
				term=append(term,n%radix)
				n=n/radix
				if n==0{
					break
				}
			}
			reverse,num:=0,0
			for i:=len(term)-1;i>=0;i--{
				reverse+=term[i]*int(math.Pow(float64(radix),float64(num)))
				num++
			}
			if IsPrime(reverse){
				fmt.Println("Yes")
			}else{
				fmt.Println("No")
			}
		}
	}
}
//函数用于判断参数是否是素数
func IsPrime(n int)bool{
	if n<=1{//小于等于1的数不是素数
		return false
	}
	for i:=2;i<=int(math.Sqrt(1.0*float64(n)));i++{
		if n%i==0{
			return false
		}
	}
	return true
}