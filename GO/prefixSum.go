package main

import (
    "fmt"
)

func PrefixSum(data []int, parent chan int) {
	if len(data) > 1 {
		mid := len(data)/2
		left := make(chan int)
		right := make(chan int)
		go PrefixSum(data[:mid], left)
		go PrefixSum(data[mid:], right)
		leftSum := <-left
		parent<- <-right + leftSum
		fromLeft := <-parent
		left<- fromLeft
		right<- fromLeft + leftSum
		parent<- <-left + <-right
	} else {
		parent<- data[0]
		data[0] += <-parent
		parent<- 0
	}
}

func main() {
	data:= []int{2, 6, 1, 7, 3, 0, 12, 8}
	ch := make(chan int)
	fmt.Println(data)
	go PrefixSum(data, ch)
	<-ch
	ch<-0
	<-ch
	fmt.Println(data)
}