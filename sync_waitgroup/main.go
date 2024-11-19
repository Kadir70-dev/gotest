package main

import (
	"fmt"
	"sync"
)


func worker (i int,wg *sync.WaitGroup){
	defer wg.Done()
	fmt.Printf("worker %d ,started\n",i)
	// some task is happening
	fmt.Printf("worker %d ,end\n",i)

}
func main() {
	fmt.Println("Explore goroutines")

	var wg sync.WaitGroup
	//start 3 worker goroutines
	for i := 0; i <= 3; i++ {
		wg.Add(1) //INCREMENT THE WAITGROUP COUNTER
		go worker(i,&wg);
	}
	wg.Wait()
	fmt.Println("Workers task completed")

}
