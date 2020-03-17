package main

import (
	"fmt"
	"sync"
	)

func increment(val *int, l sync.Mutex){
	l.Lock()
	//val = val + 1
	(*val)++
	l.Unlock()
}
func decrement(val *int, l sync.Mutex){
	l.Lock()
	//val = val - 1
	(*val)--
	l.Unlock()
}

func main() {
	var wg sync.Mutex
	var lol sync.Mutex
	j := 0
	for i := 0; i < 6; i++{
		go increment(&j,wg)
		lol.Lock()
		fmt.Println("Inc ", j)
		lol.Unlock()
		go decrement(&j,wg)
		lol.Lock()
		fmt.Println("Dec ", j)
		lol.Unlock()
	}
}