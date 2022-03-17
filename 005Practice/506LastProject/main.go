package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"
)

// Job represents a job to be executed, with a name and a number and a delay
type Job struct {
	Name   string
	Delay  time.Duration
	Number int
}

// Worker will be our concurrency-friendly worker
type Worker struct {
	Id         int           //IN -Id of worker
	JobQueue   chan Job      //OUT-send itself and extract the data contain in workerpool dispat
	WorkerPool chan chan Job //IN-Worker pool recived from dispatcher
	QuitChan   chan bool     //ONLINE EXIST IN WORKER
}

//Dispatcher encargado de enviar los jobs a los Workers
type Dispatcher struct {
	WorkerPool chan chan Job //OUT- Worker pool send to NewWorker
	MaxWorkers int           //IN- limit the workers
	JobQueue   chan Job      //IN- contains the data to send worker
}

//Constructor para el Worker
func NewWorker(id int, workerPool chan chan Job) *Worker {
	return &Worker{
		Id:         id,
		JobQueue:   make(chan Job),
		WorkerPool: workerPool,
		QuitChan:   make(chan bool),
	}
}

//Metodo para worker - siempre pendiente a todos los trabajos...
func (w Worker) Start() {
	go func() {
		for {
			w.WorkerPool <- w.JobQueue
			select {
			case job := <-w.JobQueue:
				fmt.Printf("Worker with id %d Started\n", w.Id)
				fib := Fibonacci(job.Number)
				time.Sleep(job.Delay)
				fmt.Printf("Worker with id %d Finished with result %d\n", w.Id, fib)
			case <-w.QuitChan:
				fmt.Printf("Worker with id %d Stopped\n", w.Id)
			}
		}
	}()
}

func (w Worker) Stop() {
	go func() {
		w.QuitChan <- true
	}()
}

func Fibonacci(n int) int {
	if n <= 1 {
		return n
	}
	return Fibonacci(n-1) + Fibonacci(n-2)
}

//Constructor para dispatcher
func NewDispatcher(jobQueue chan Job, maxWorkers int) *Dispatcher {
	worker := make(chan chan Job, maxWorkers)
	return &Dispatcher{
		JobQueue:   jobQueue,
		MaxWorkers: maxWorkers,
		WorkerPool: worker,
	}
}

//Encolar jobs iterando de manera infinita/ y cuando se lea un job con select
//enviar los datos del JobQueue al chan chan compartido por medio de su chan
//el propio chan
//Para encolar los jobs

func (d *Dispatcher) Dispatch() {
	for {
		select {
		case job := <-d.JobQueue:
			go func() {
				workerJobQueue := <-d.WorkerPool
				workerJobQueue <- job
			}()
		}
	}
}

//Itera de la cantidad maxima de Workes aquie los creamos
//creandolo con el workerpool que contiene los datos...
//E iniciamos el worker...
func (d *Dispatcher) Run() {
	for i := 0; i < d.MaxWorkers; i++ {
		worker := NewWorker(i, d.WorkerPool)
		worker.Start()
	}
	go d.Dispatch()

}

/////////////////////////////////////////
func RequestHandler(w http.ResponseWriter, r http.Request, jobQueue chan Job) {
	if r.Method != "POST" { //GET, PUT, DELETE
		w.Header().Set("Allow", "POST")
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
	delay, err := time.ParseDuration(r.FormValue("delay"))
	if err != nil {
		http.Error(w, "Invalid Delay", http.StatusBadRequest)
		return //para no seguir prosesando
	}
	value, err := strconv.Atoi(r.FormValue("value"))
	if err != nil {
		http.Error(w, "Invalid Value", http.StatusBadRequest)
		return //para no seguir prosesando
	}
	name := r.FormValue("name")
	if name == "" {
		http.Error(w, "Invalid Name", http.StatusBadRequest)
		return //para no seguir prosesando
	}

	job := Job{
		Name:   name,
		Delay:  delay,
		Number: value,
	}
	jobQueue <- job
	w.WriteHeader(http.StatusCreated)
}

func main() {
	const (
		maxWorkers   = 4  //maximos trabajadores
		maxQueueSize = 20 //maxima tareas concurrentes
		port         = ":8081"
	)
	jobQueue := make(chan Job, maxQueueSize)
	dispatcher := NewDispatcher(jobQueue, maxWorkers)
	dispatcher.Run()

	//http://localhost:8081/fib
	http.HandleFunc("/fib", func(w http.ResponseWriter, r *http.Request) {
		RequestHandler(w, *r, jobQueue)
	})
	log.Fatal(http.ListenAndServe(port, nil))
}
