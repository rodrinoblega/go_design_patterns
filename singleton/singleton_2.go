package main

import (
	"fmt"
	"sync"
)

var once sync.Once

type single struct{}

var singleInstance *single

func getInstance() *single {
	if singleInstance == nil {
		once.Do(
			func() {
				fmt.Println("Creating single instance now.")
				singleInstance = &single{}
			})
	} else {
		fmt.Println("Single instance already created.")
	}

	return singleInstance
}

func main() {
	var wg sync.WaitGroup

	for i := 0; i < 30; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			getInstance()
		}()
	}

	wg.Wait()
}
