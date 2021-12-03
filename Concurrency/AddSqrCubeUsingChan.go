package main

import "fmt"

func main(){
	num := 123
	SqrChan := make(chan int)
	CubeChan := make(chan int)
	
	go CalSqr(num, SqrChan)
	go CalCube(num, CubeChan)
	
	res := <-SqrChan + <-CubeChan
	
	fmt.Println(res)

}

func CalSqr(num int, c chan int){
	sum := 0
	for num>0{
		digit := num%10
		sum += digit*digit
		num = num/10	
	}
	c <- sum
}

func CalCube(num int, c chan int){
	sum := 0
	for num>0{
		digit := num%10
		sum += digit*digit*digit
		num = num/10	
	}
	c <- sum
}
