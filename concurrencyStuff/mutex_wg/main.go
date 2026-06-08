//Race conditions occur when 2 processes try to access same data, which may result in both processes getting stale/old data 
// and wwork on that rather what shouldhave happened was worker a, access data, operates on it, then worker b pickup the data, 
// operate on the data modified by A and not the stale one. The program we wrote here, basically increaments a counter inside a 
// loop. and with each iteration, a new go routine is formed of fucn() that increaments counter. 
// but as those go routines are working concurrently on a single variable, i.e counter which may have any value written to it.
//  so there is a chance that 2 workers acces a int value of 8 and both update it to 9 instead of 10.


// func main() {
// 	var wg sync.WaitGroup

// 	counter := 0

// 	for i := 0; i < 1000; i++ {
// 		wg.Add(1)

// 		go func() {
// 			defer wg.Done()

// 			counter++
// 		}()
// 	}

// 	wg.Wait()

// 	fmt.Println(counter)
// }

//FIX, add mutex
package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	var mu sync.Mutex
	counter := 0

	for i := 0; i < 1000; i++ {
		wg.Add(1)

		go func() {
			defer wg.Done()
			mu.Lock()
			counter++
			//why defer here? because suppose you have a function between lock and unlock that may or may not return something.
			//(the part between lock and unlock is called the critical section) so if it returns an err or smthn, then the lock is held 4ever(deadlock)
			//so to prevent that, a defer is used to gurantee a unlock at just before end of function
			defer mu.Unlock()
		}()
	}

	wg.Wait()

	fmt.Println(counter)
}

//alright if 10k goroutines are increamenint/operating on a same counter then i think adding a exclusive lock (coming from dbms) would work. say on the first iteration, func is ran, it adds wg.add(1) then implements a lock on counter, an x lock to be specific which prevents other workers to access the counter. they are allowed to access only after worker 1 releases its lock. that is when after the counter is increamented. now worker 1 did its job, next iteration say a random worker 69 gets chance and locks it. updates it to int value 2, then releases the lock. this way it ensures that all workers access the same state. but one question, wont this kill the whole purpose of concurrecny?