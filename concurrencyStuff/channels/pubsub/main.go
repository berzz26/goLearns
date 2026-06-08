package main

import (
	"fmt"
	"sync"
)

func worker(id int, ch chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	for job := range ch {
		fmt.Printf("worker %d got %d\n", id, job)
	}
}

func main() {
	ch := make(chan int)
	//adding worker grp coz channels do not solve execution of go routines. they solve data transfer between go routines.
	//so without this, i was having the main execute and one job staying stale. so added a wg to ensure all workers execture before main
	var wg sync.WaitGroup

	wg.Add(2)

	go worker(1, ch, &wg)
	go worker(2, ch, &wg)

	ch <- 1
	ch <- 2
	ch <- 3
	ch <- 4

	close(ch)
	wg.Wait()

}
