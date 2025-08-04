package main

import (
	"fmt"
)

// Merge принимает два read-only канала и возвращает выходной канал,
// в который последовательно (в любом порядке) будут отправлены все значения
// из обоих входных каналов.
//
// Выходной канал должен быть закрыт после того, как оба входных канала закроются.
// Merge не должен закрывать входные каналы
//
// Для проверки решения запустите тесты: go test -v
func Merge(ch1, ch2 <-chan int) <-chan int {
	resChan := make(chan int)
	go func() {
		defer close(resChan)
		for ch1 != nil || ch2 != nil { // Типичный паттерн fan-in через for select и проверки на nil
			select {
			case number, ok := <-ch1:
				if !ok {
					ch1 = nil
					continue
				}
				resChan <- number
			case number, ok := <-ch2:
				if !ok {
					ch2 = nil
					continue
				}
				resChan <- number
			}
		}
	}()
	return resChan
}

func main() {
	a := make(chan int)
	b := make(chan int)

	go func() {
		defer close(a)
		a <- 4
		a <- 1
	}()

	go func() {
		defer close(b)
		b <- 2
		b <- 4
	}()

	for v := range Merge(a, b) {
		fmt.Println(v)
	}
}
