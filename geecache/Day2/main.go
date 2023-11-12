package main

import (
	"fmt"
	"sync"
	"time"
)

var m sync.Mutex
var set = make(map[int]bool, 0)

func printOnce(num int) {
	if _, ok := set[num]; !ok {
		fmt.Printf("%d", num)
	}
	set[num] = true
}
func printOnce1(num int) {
	m.Lock()
	if _, ok := set[num]; !ok {
		fmt.Printf("%d", num)
	}
	set[num] = true
	m.Unlock()
}
func printOnce2(num int) {
	m.Lock()
	defer m.Unlock()
	if _, ok := set[num]; !ok {
		fmt.Printf("%d", num)
	}
	set[num] = true
}
func main() {
	for i := 0; i < 10; i++ {
		go printOnce(100)
	}
	time.Sleep(time.Second)
}
