package main

import (
	"fmt"
	"log"
	"time"
)

func Fibbonacci(n int) int {
	if n <= 1 {
		return n
	}
	return Fibbonacci(n-1) + Fibbonacci(n-2)
}

type Memory struct {
	cache map[int]Results
}

//Cache de operaciones
type Results struct {
	value int
	err   error
}

//Constructor
func NewCache() Memory {
	return Memory{
		cache: make(map[int]Results),
	}
}

//Logica cache
//Comparacion del cache
func (m *Memory) Get(key int) (int, error) {
	result, exist := m.cache[key]
	if !exist {
		result.value = Fibbonacci(key)
		m.cache[key] = result
	}
	return result.value, result.err
}

func main() {
	cache := NewCache()
	fibo := []int{42, 40, 41, 42, 38}
	for _, n := range fibo {
		start := time.Now()
		value, err := cache.Get(n)
		if err != nil {
			log.Println(err)
		}
		fmt.Printf("%d, %s, %d\n", n, time.Since(start), value)
	}
}
