package Pat
import "fmt"

//本题主要考察进制之间的转换
func run(){
	var a,b int32
	fmt.Scanf("%d%d",&a,&b)
	//筛数求余法
	var resultNum []int32
	for a!=0{
		resultNum=append(resultNum,a%b)
		a=a/b
	}
	//首为对应匹配
	flag:=false
	for i:=0;i<len(resultNum)/2;i++{
		if resultNum[0]!=resultNum[len(resultNum)-1-i]{
			flag=true
			fmt.Printf("No\n")
			break
		}
	}
	if !flag{//是回文的话
		fmt.Printf("Yes\n")
	}
	for i:=len(resultNum)-1;i>=0;i--{
		fmt.Printf("%d",resultNum[i])
		if i!=0{
			fmt.Printf(" ")
		}
	}
	if len(resultNum)==0{
		fmt.Printf("0")
	}
}
