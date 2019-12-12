package Pat

import "fmt"
import "strconv"
//这题主要考察的是字符串和整数之间的转换
//Given a non-negative integer N, your task is to compute the sum of all the digits of N, and output every digit of the sum in English.
//
//Input Specification:
//
//Each input file contains one test case. Each case occupies one line which contains an N (≤10
//​100
//​​ ).
//
//Output Specification:
//
//For each test case, output in one line the digits of the sum in English words. There must be one space between two consecutive words, but no extra space at the end of a line.
//
//Sample Input:
//
//12345
//Sample Output:
//
//one five
var noNegativeN string
func SpellItRight(){
	fmt.Scanf("%s",&noNegativeN)
	//fmt.Println(noNegativeN)
	sum:=0
	mapSum:=make(map[rune]string,0)
	mapSum['1']="one"
	mapSum['2']="two"
	mapSum['3']="three"
	mapSum['4']="four"
	mapSum['5']="five"
	mapSum['6']="six"
	mapSum['7']="seven"
	mapSum['8']="eight"
	mapSum['9']="nine"
	mapSum['0']="zero"
	for _,j:=range noNegativeN{
		sum+=int(j-48)//根据ascll属性
	}
	sumString:=strconv.Itoa(sum)
	for i,j:=range sumString{
		if i!=len(sumString)-1{
			fmt.Printf("%s ",mapSum[j])
		}else{
			fmt.Printf("%s",mapSum[j])
		}
	}
}

