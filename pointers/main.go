package main

import "fmt"




func main(){

	i,j:=43,222

	p :=&i
	fmt.Println(*p)
	*p=21
	fmt.Println(i)//see the new value as using pointer we changed the value of i because p is pointing to the address of i
	
	//now point p to j

	p=&j
	*p=*p/22

	fmt.Println(i,j)




}
