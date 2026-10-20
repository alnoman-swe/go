package main

import "fmt"


func main(){
	sum :=1
	// for i:=0;i<10;i++{
	// 	sum +=i
	// }
	//while loop
	// for sum<9{
	// 	sum+=sum
	// }
	//condition to break the loop infinite loop
	for{
		if(sum>100){
			break
		}
		sum+=sum
	}
	fmt.Println(sum)
}
