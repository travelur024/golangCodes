package main

import (
	"fmt"
	"sync"
	"time"
)

func ExtensiveFinobacci(n int) int {
	fmt.Printf("Calcule Expensive Fibonacci for %d \n", n)
	time.Sleep(5 * time.Second)
	return n
}

//Crear service

type Service struct {
	InProgress map[int]bool       //Para saber si el numero esta en proceso
	IsPending  map[int][]chan int //slice de canales/Representado como los workes que esperan resultado
	Lock       sync.RWMutex
}

//Metodo
func (s *Service) Work(job int) {
	s.Lock.RLock()
	exists := s.InProgress[job]
	if exists {
		s.Lock.RUnlock()
		response := make(chan int)
		defer close(response)

		s.Lock.Lock()
		s.IsPending[job] = append(s.IsPending[job], response)
		s.Lock.Unlock()
		fmt.Printf("Waiting for Response job: %d\n", job)
		resp := <-response
		fmt.Printf("Response Done, received %d\n", resp)
		return
	}
	//Work - exist false
	s.Lock.RUnlock()

	s.Lock.Lock()
	s.InProgress[job] = true
	s.Lock.Unlock()
	fmt.Printf("Calculate Fibonacci for %d\n", job)
	//Traer resultado
	result := ExtensiveFinobacci(job)
	/////////////////////////////////////////////////////////////////////////////////////////
	s.Lock.Lock()
	//Pendingo workes que esperan el reusltado///
	pendingWorkes, exist := s.IsPending[job]
	s.Lock.Unlock()

	if exist {
		for i, pendingWorker := range pendingWorkes {
			fmt.Println("Pendind...", i, job)
			pendingWorker <- result
		}
		fmt.Printf("Result sent - all pending workers ready job:%d\n", job)
	}
	s.Lock.Lock()
	s.InProgress[job] = false
	s.IsPending[job] = make([]chan int, 0)
	s.Lock.Unlock()
}

func NewService() *Service {
	return &Service{
		InProgress: make(map[int]bool),
		IsPending:  make(map[int][]chan int),
	}
}

func main() {
	service := NewService()
	jobs := []int{5, 6, 7, 5, 6, 7}
	var wg sync.WaitGroup
	wg.Add(len(jobs))
	for _, n := range jobs {

		go func(job int) {
			defer wg.Done()
			service.Work(job)
		}(n)
	}
	wg.Wait()
}
