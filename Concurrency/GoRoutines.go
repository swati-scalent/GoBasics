package main

import "fmt"
import "time"

func main(){
	for i:= 1;i<11;i++{
	go f()
	time.Sleep(2)
	}
	fmt.Scanln()
}

func f(){
	for i :=1;i<11;i++{
		fmt.Println(i)
	
	}
}
