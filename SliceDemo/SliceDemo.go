package main

import "fmt"

func main(){
	arr := make([]int,10,10)
	arr = append(arr,2,3,5)
	
	for _,val := range arr{
	
		fmt.Println(val)
	
	}


}
