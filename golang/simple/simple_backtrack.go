package main

import (
	"fmt"
	"time"
)

var queen int

func main() {
	fmt.Scan(&queen)

	bf_t := time.Now()
	fmt.Println(bf_t)

	qhistory := make([]int, queen, queen)

	fmt.Println(putQueen(0, qhistory))

	af_t := time.Now()
	fmt.Println(af_t.Sub(bf_t))
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
