package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	n := 3        // макс число работающих горутин
	workCnt := 10 //число задач
	semafor := make(chan struct{}, n)
	wg := sync.WaitGroup{}
	wg.Add(workCnt)
	for i := 0; i < workCnt; i++ {
		semafor <- struct{}{} //захват слота
		fmt.Println("add semafor")
		go func(i int) {
			defer wg.Done()
			defer func() {
				<-semafor
				fmt.Println("free semafor")
			}() //освобождение слота
			fmt.Println("Work start", i)
			time.Sleep(time.Second)
			fmt.Println("Work finish", i)
		}(i)
	}
	wg.Wait()
}
