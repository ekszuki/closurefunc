package main

import (
	"fmt"
	"math/rand"
)

func main() {
	target := 50
	quitChan := make(chan bool)

	closureFunc := createClosureFunc()
	for {
		go closureFunc(target, quitChan)

		select {
		case v := <-quitChan:
			if v {
				fmt.Println("Target found...: ", target)
				return
			}
		}
	}
}

func createClosureFunc() func(target int, quitChan chan bool) {
	count := 0
	return func(target int, quitChan chan bool) {
		count++
		r := rand.Intn(target * 2)

		fmt.Printf("%d (executions) ->  Target: %d - rand value: %d \n", count, target, r)

		if r == target {
			quitChan <- true
		}

		quitChan <- false
	}
}
