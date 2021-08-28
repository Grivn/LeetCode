package main

import (
	"fmt"
	"io"
)

type stream struct {
	taskMap map[int]*task
}

type task struct {
	id int

	long int

	priori map[int]bool

	finished bool

	total int
}

func input() ([]*stream, error) {
	var T int

	_, err := fmt.Scan(&T)

	if err != nil {
		return nil, err
	}

	res := make([]*stream, T)

	for i:=0; i<T; i++ {
		res[i] = &stream{
			taskMap: make(map[int]*task),
		}

		var N int

		_, _ = fmt.Scan(&N)

		for j:=0; j<N; j++ {
			t := &task{id: j+1, priori: make(map[int]bool), finished: false}
			_, _ = fmt.Scan(&t.long)

			var prioriCount int

			_, _ = fmt.Scan(&prioriCount)

			for k:=0; k<prioriCount; k++ {
				var prioriID int

				_, _ = fmt.Scan(&prioriID)

				t.priori[prioriID] = true
			}
			res[i].taskMap[j+1] = t
		}

	}

	return res, nil
}

func main() {
	for {
		inputs, err := input()

		if err == io.EOF {
			break
		}

		for _, s := range inputs {
			long := calculate(s)
			fmt.Println(long)
		}
	}
}

func calculate(s *stream) int {
	long := 0

	for _, t := range s.taskMap {
		long = maxInt(long, t.totalTime(s))
	}

	return long
}

func maxInt(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func (t *task) totalTime(s *stream) int {
	if t.finished {
		return t.total
	}

	total := 0

	maxPriori := 0
	for id := range t.priori {
		prioriT := s.taskMap[id]

		maxPriori = maxInt(maxPriori, prioriT.totalTime(s))
	}

	total += maxPriori

	total += t.long

	t.finished = true

	t.total = total

	return total
}

func output(inputs []*stream) {
	for _, s := range inputs {
		for _, t := range s.taskMap {
			fmt.Printf("id %d, long %d, priori %v\n", t.id, t.long, t.priori)
		}
	}
}
