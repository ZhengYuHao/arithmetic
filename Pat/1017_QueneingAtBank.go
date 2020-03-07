package Pat
import (
	"fmt"
)
type userData struct {
	hour int32   //到达的时
	minute int32 //分
	second int32 //秒
	defineTime int32 //将时分秒统一换算成秒的绝对值
	processTime int32 //服务需要消耗的时间
	inTime      int32

}
//对玩家信息根据绝对值时间进行排序(冒泡排序)
func SortByDefineTime(termUserDataInfor []userData)[]userData{
	for i:=0;i<len(termUserDataInfor)-1;i++{
		for j:=0;j<len(termUserDataInfor)-1-i;j++{
			if termUserDataInfor[j].defineTime>termUserDataInfor[j+1].defineTime{
				termUserDataInfor[j],termUserDataInfor[j+1]=termUserDataInfor[j+1],termUserDataInfor[j]
			}
		}
	}
	return termUserDataInfor
}
func QueueningAtBank(){
	N,K:=int(0),int(0)
	fmt.Scanf("%d %d",&N,&K)
	var UserDataInfor []userData
	var TermWindows []userData
	for i:=int(0);i<N;i++{
		termUserInfor:=userData{}
		termHour,termMinute,termSecond,termProcessTime:=int32(0),int32(0),int32(0),int32(0)
		fmt.Scanf("%d:%d:%d %d",&termHour,&termMinute,&termSecond,&termProcessTime)
		termDefineTime:=termHour*3600+termMinute*60+termSecond
		termUserInfor.hour=termHour
		termUserInfor.minute=termMinute
		termUserInfor.second=termSecond
		termUserInfor.processTime=termProcessTime*60
		termUserInfor.defineTime=termDefineTime
		if termDefineTime>=(17*3600+1){//如果来的时间是大于等于17时多1秒，不服务
			continue
		}
		UserDataInfor=append(UserDataInfor,termUserInfor)
	}
	UserDataInfor=SortByDefineTime(UserDataInfor)
	userIndex:=0
	sumTime:=int32(0)
	for i:=int32(8*3600);len(UserDataInfor)>0;i++{
		//先出队
		for x,y:=range TermWindows{
			if i-y.inTime==y.processTime{
				TermWindows=append(TermWindows[0:x],TermWindows[x+1:]...)
			}
		}
		for len(TermWindows)<K&&i>=UserDataInfor[userIndex].defineTime&&userIndex<len(UserDataInfor){//这表示窗口还有位置
			sumTime+=i-UserDataInfor[userIndex].defineTime//等待时间累加
			UserDataInfor[userIndex].inTime=i
			TermWindows=append(TermWindows,UserDataInfor[userIndex])
			userIndex++//入完列之后，用户信息index往后面+
			if userIndex==len(UserDataInfor){
				goto f
			}
		}
		if userIndex==len(UserDataInfor){//所有符合条件的玩家都服务完了
			break
		}
	}
f:
	if len(UserDataInfor)==0{
		fmt.Println("0.0")
	}else{
		fmt.Printf("%.1f\n",float64(sumTime)/float64(60*len(UserDataInfor)))
	}
}