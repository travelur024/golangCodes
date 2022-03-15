package main

import "fmt"

func main() {

	tasks := []int{2, 3, 4, 5, 7, 10, 12, 40}
	nWorkers := 3
	jobs := make(chan int, len(tasks))
	result := make(chan int, len(tasks))

	for i := 0; i < nWorkers; i++ {
		go Worker(i, jobs, result)
		fmt.Println("Creada", i)
	}
	//Pero no mandamos data entonces se termina

	//
	//Mandamos la data pero no la extraemos se termina
	for _, value := range tasks {
		jobs <- value
	}
	close(jobs)

	//Bloquea el programa hasta que se termine de ejecutar los resultados
	for r := 0; r < len(tasks); r++ {
		<-result
	}

}

//Primero definir la funcnion que va a realizar el trabajo
//Worker contine id, canal que lee data jobs y canal que escribe el nial del job result
func Worker(id int, jobs <-chan int, result chan<- int) {
	for job := range jobs {
		fmt.Printf("Worker with id %d started fib with %d\n", id, job)
		fib := Fibonacci(job)
		fmt.Printf("Worked with id %d, job %d and fib %d\n", id, job, fib)
		result <- fib
	}
}

func Fibonacci(n int) int {
	if n <= 1 {
		return n
	}
	return Fibonacci(n-1) + Fibonacci(n-2)
}
