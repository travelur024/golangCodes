//Aplication that resolve multiple actions, with an id Worker, showing the current Job,  using concurrency
package main

import (
	"fmt"
)

func main() {
	jobs := []int{1, 2, 3, 4, 5, 6, 7, 8}
	Workers := 4
	jo := make(chan int, len(jobs))
	re := make(chan int, len(jobs))
	//Creating the Workers (Ready for <-chan)
	for i := 0; i < Workers; i++ {
		go WorkerFunc(i, jo, re)
	}
	//Passing the data for operation (Send to chan <-)
	for i := 0; i < len(jobs); i++ {
		jo <- jobs[i]
	}
	//Waiting
	for i := 0; i < len(jobs); i++ {
		<-re
	}

}

func WorkerFunc(id int, jobs <-chan int, result chan<- int) {
	for job := range jobs {
		fmt.Printf("Worker: %v . Job: %v\n", id, job)
		r := DoubleNumber(&job)
		fmt.Printf("Worker: %v . Job: %v , The result of the operation is: %v\n", id, job, r)
		result <- r
	}
}

func DoubleNumber(a *int) int {
	return *a * 2
}
