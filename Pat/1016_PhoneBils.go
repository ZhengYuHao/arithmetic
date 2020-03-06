package Pat

import "fmt"

var TimeValuue [24]int32 //每个小时阶段的长途电话耗费
type Customer struct {  //每个玩家消费的具体信息
	name string
	month int32
	day int32
	hour int32
	minute int32
	lineStatus bool  //1表示在线on_line，0表示离线off_line
	tipTime int32     //程序需要用到的tip，表示这个月0时0分开始的绝对时间
}

var NameCompilatons []string

//函数用于判断，原来的数组里面是否包含出传入的这个值
func IsRepetition(nameCompilation []string,name string)bool{
	for _,j:=range nameCompilation{
		if j==name{
			return true
		}
	}
	return false
}

//函数用于字母排序
func SortByAscll(nameCompilation []string)[]string{
	for i:=0;i<len(nameCompilation)-1;i++{
		for j:=0;j<len(nameCompilation)-1-i;j++{
			if CompileByAscll(nameCompilation[j],nameCompilation[j+1]){
				nameCompilation[j],nameCompilation[j+1]=nameCompilation[j+1],nameCompilation[j]
			}
		}
	}
	return nameCompilation
}

//函数用于判断term1是否大于term2,<=返回false
func CompileByAscll(term1 string,term2 string)bool{
	if len(term1)>len(term2){
		for i,_:=range term2{
			if term1[i]>term2[i]{
				return true
			}else if term1[i]<term2[i]{
				return false
			}else if term1[i]==term2[i]{
				continue
			}
		}
		return true
	}else {
		for i,_:=range term1{
			if term1[i]>term2[i]{
				return true
			}else if term1[i]<term2[i]{
				return false
			}else if term1[i]==term2[i]{
				continue
			}
		}
		return false
	}
	return true
}

func SortByTipTime(userData []Customer)[]Customer{
	for i:=0;i<len(userData)-1;i++{//冒泡最多冒n-1次
		for j:=0;j<len(userData)-1-i;j++{
			if userData[j].tipTime>userData[j+1].tipTime{
				userData[j],userData[j+1]=userData[j+1],userData[j]
			}
		}
	}
	return userData
}
func PhoneBils()  {
	//初始化24个小时，每个小时长途电话的单价，Start
	dayValue:=int32(0)
	for i:=0;i<24;i++{
		term:=int32(0)
		fmt.Scanf("%d",&term)
		TimeValuue[i]=term
		dayValue+=term*60
	}
	//End
	//数据条数,Start
	dataNum:=int32(0)
	fmt.Scanf("%d",&dataNum)
	//End
	//录入每个玩家的信息Start
	customerList:=make(map[string][]Customer,0)
	for i:=int32(0);i<dataNum;i++{
		term:=Customer{}
		termFlag:=string("")
		fmt.Scanf("%s%s",&term.name)
		fmt.Scanf("%d:%d:%d:%d",&term.month,&term.day,&term.hour,&term.minute)
		fmt.Scanf("%s",&termFlag)
		if termFlag=="on-line"{
			term.lineStatus=true
		}else if termFlag=="off-line"{
			term.lineStatus=false
		}
		term.tipTime=term.day*1440+term.hour*60+term.minute//一个月中的绝对时间
		customerList[term.name]=append(customerList[term.name],term)
		if !IsRepetition(NameCompilatons,term.name){
			NameCompilatons=append(NameCompilatons,term.name)
		}
	}
	//End
	//对玩家的name,按照字母排序Start
	NameCompilatons=SortByAscll(NameCompilatons)
	//End
	for _,userName:=range NameCompilatons{//遍历玩家，输出改玩家的账单
		userData:=customerList[userName]
		userData=SortByTipTime(userData)
		count:=int32(0)
		valueCountPlus:=float64(0)
		//fmt.Println(userData)
		for index,termData:=range userData{
			if index==len(userData)-1{
				break
			}
			nextData:=userData[index+1]
			valueCount:=int32(0)
			if termData.lineStatus==true&&nextData.lineStatus==false{
				if nextData.tipTime-termData.tipTime>0{//之间存在时间差，才进行账单输出
					if nextData.day-termData.day==0{//时间是同一天
						if nextData.hour-termData.hour==0{//如果是同一个小时
							valueCount+=TimeValuue[nextData.hour]*(nextData.minute-termData.minute)
						}else{//如果不是同一个小时
							for i:=termData.hour;i<=nextData.hour;i++{
								if i==termData.hour{
									valueCount+=TimeValuue[i]*(60-termData.minute)
								}else  if i==nextData.hour{
									valueCount+=TimeValuue[i]*nextData.minute
								}else{
									valueCount+=TimeValuue[i]*60
								}
							}
						}
					}else {//如果不是同一天
						for i:=termData.day;i<=nextData.day;i++{
							if i==termData.day{//on-line的那一天
								for j:=termData.hour;j<24;j++{
									if j==termData.hour{
										valueCount+=TimeValuue[termData.hour]*(60-termData.minute)
									}else{
										valueCount+=TimeValuue[j]*60
									}
								}
							}else  if i==nextData.day{//off-line的那一天
								for j:=int32(0);j<=nextData.hour;j++{
									if j==nextData.hour{
										valueCount+=TimeValuue[nextData.hour]*(nextData.minute)
									}else{
										valueCount+=TimeValuue[j]*60
									}
								}
							}else{
								valueCount+=dayValue//加上一整天的价值
							}
						}
					}
					if count==0{
						fmt.Printf("%s %02d\n",userName,termData.month)
						count++
					}
					fmt.Printf("%02d:%02d:%02d ",termData.day,termData.hour,termData.minute)
					fmt.Printf("%02d:%02d:%02d ",nextData.day,nextData.hour,nextData.minute)
					fmt.Printf("%d ",nextData.tipTime-termData.tipTime)
					fmt.Printf("$%.2f\n",(1.0*float64(valueCount))/100)
					valueCountPlus+=(1.0*float64(valueCount))/100
				}
			}
		}
		if count!=0{
			fmt.Printf("Total amount: $%.2f\n",valueCountPlus)
		}
	}
}