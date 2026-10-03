package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var reader *bufio.Reader

func main(){
	reader= bufio.NewReader(os.Stdin)
	name:=readString("What is your name? ")
	
	fmt.Println("Hello",name)
}

func prompt(){
	fmt.Print("->")
}

func readString(s string )	string{
	fmt.Println(s)
	prompt()
	
	userInput,_:=reader.ReadString('\n')
	userInput=strings.Replace(userInput,"\n","",-1)
	userInput=strings.Replace(userInput,"\r\n","",-1)
	return userInput

}

func readInt(s string) int{
	fmt.Println(s)
	prompt()
	
	userInput,_:=reader.ReadString('\n')
	userInput=strings.Replace(userInput,"\n","",-1)
	userInput=strings.Replace(userInput,"\r\n","",-1)
	num,err:=strconv.Atoi(userInput)

	if err!=nil{
		fmt.Println("Invalid input. Please enter a valid integer.")
	}
	return num
}