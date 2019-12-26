package Pat

import "fmt"

func WorldCupBetting(){
	summit:=float64(1)
	indexMap:=make(map[int]rune,0)
	for i:=0;i<3;i++{
		t1,t2,t3:=0,0,0
		fmt.Scanf("%f%f%f",&t1,&t2,&t3)
		term:=t1
		indexMap[i]='W'
		if t2>term{
			indexMap[i]='T'
			term=t2
		}
		if t3>term{
			term=t3
			indexMap[i]='L'
		}
		summit*=float64(term)
	}
	summit=(summit*0.65-1)*2
	fmt.Printf("%v %v %v %.2f",string(indexMap[0]),string(indexMap[1]),string(indexMap[2]),summit)
}