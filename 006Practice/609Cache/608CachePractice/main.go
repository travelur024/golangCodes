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

//Tipo de dato necesario en la asignacion
type Function func(key int) (interface{}, error)

type FunctionResult struct {
	value interface{}
	err   error
}

type Memory struct {
	f     Function
	cache map[int]FunctionResult
}

//Crear constructor... del cache
func NewCache(f Function) Memory {
	return Memory{
		f:     f,
		cache: make(map[int]FunctionResult),
	}
}

//Comparacion del cache
func (m *Memory) Get(key int) (interface{}, error) {
	result, exist := m.cache[key]
	if !exist {
		result.value, result.err = m.f(key)
		m.cache[key] = result
	}
	return result.value, result.err
}

func GetFibonacci(n int) (interface{}, error) {
	return Fibbonacci(n), nil
}

func main() {
	cache := NewCache(GetFibonacci)
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
