package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var reader *bufio.Reader

func main(){
	
	name:=readString("What is your name? ")
	
	fmt.Println("Hello",name)
}

func prompt(){
	fmt.Print("->")
}

func readString(s string )	string{
	fmt.Println(s)
	prompt()
	reader= bufio.NewReader(os.Stdin)
	userInput,_:=reader.ReadString('\n')
	userInput=strings.Replace(userInput,"\n","",-1)
	userInput=strings.Replace(userInput,"\r\n","",-1)
	return userInput

}