package main

import "fmt"
func main(){
	// names := [4]string{//fixed type array
	// 	"Tanvir",
	// 	"Rifat",
	// 	"Shuvo",
	// 	"Rasel",
	// }
	// fmt.Println(names)
	// a:= names[0:2]
	// b:=names[1:3] //here length is 2 and capacity is 3 and pointer is pointing to the second element of the array
	// fmt.Println(a,b)
	
	// b[1]="Noman"//slice is a reference type so if we change the value of slice b then it will also change the value of array names
	// fmt.Println(a,b)
	// fmt.Println(names)

 //literal syntax of slice
	slice1 := []int{10,20,30,40,50}

	//slice struct
	slice2:= []struct{
		X string
		Y int
	}{
		{"Tanvir",20},
		{"Tahsin",40},

	}

	fmt.Println(slice1)
	fmt.Println(slice2)
//zero value of slice is nil
	slice3:= []int{}
	fmt.Println(slice3)
	var slice4 []int
	if slice4==nil{
	fmt.Println("slice4 is nil")
	}


	}



