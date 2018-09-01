package main

import (
	"fmt"
	"runtime"
	"time"
)

var queen int

func main() {
	cpus := runtime.NumCPU()
	runtime.GOMAXPROCS(cpus)

	fmt.Scan(&queen)

	bf_t := time.Now()
	fmt.Println(bf_t)

	var sum int

	receiver := worker(queen)
	for i := 0; i < queen; i++ {
		sum += <-receiver
	}
	fmt.Println(sum)

	af_t := time.Now()
	fmt.Println(af_t.Sub(bf_t))
}

func worker(queen int) <-chan int {
	receiver := make(chan int)
	for i := 0; i < queen; i++ {
		go func(i int) {
			org := make([]int, queen, queen)
			org[0] = i
			result := putQueen(1, org)
			receiver <- result
		}(i)
	}
	return receiver
}

func putQueen(pos int, history []int) int {
	if pos == queen {
		return 1
	}

	success := 0
	for i := 0; i < queen; i++ {
		history[pos] = i
		if validate(pos, history) {
			history_copy := make([]int, len(history))
			copy(history_copy, history)
			success += putQueen(pos+1, history_copy)
		}
	}

	return success
}

func validate(pos int, history []int) bool {
	candidate := history[pos]

	for i := 1; i <= pos; i++ {
		check_pos := pos - i
		if history[check_pos] == candidate {
			return false
		}
		if history[check_pos]+i == candidate {
			return false
		}
		if history[check_pos]-i == candidate {
			return false
		}
	}

	// fmt.Println(history)
	return true
}
