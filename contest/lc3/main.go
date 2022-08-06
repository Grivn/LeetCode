package main

import "fmt"

func main() {
	tasks := []int{1, 2, 1, 2, 3, 1}
	space := 3
	date := worker(tasks, space)
	fmt.Println(date)
}
func worker(tasks []int, space int) int64 {
	date := int64(0)
	pending := make(map[int]int64) // task type, prev finished date
	for _, taskTyp := range tasks {
		prevDate, ok := pending[taskTyp]
		if !ok {
			date++
			pending[taskTyp] = date
			//fmt.Println("a", date, ":", i, ":", taskTyp, ":", pending)
			continue
		}
		interval := date + 1 - prevDate
		if interval > int64(space) {
			date++
			pending[taskTyp] = date
			//fmt.Println("b", date, ":", i, ":", taskTyp, ":", pending)
		} else {
			date = prevDate + int64(space) + 1
			pending[taskTyp] = date
			//fmt.Println("c", date, ":", i, ":", taskTyp, ":", pending)
		}
	}
	return date
}
func taskSchedulerII(tasks []int, space int) int64 {
	date := int64(0)
	taskId := 0
	pending := make(map[int]int64) // task type, latest finished date
	for {
		if len(tasks) == taskId {
			break
		}
		date++
		taskTyp := tasks[taskId]
		latest, ok := pending[taskTyp]
		if !ok {
			pending[taskTyp] = date
			taskId++
			continue
		}
		waited := date - latest
		if waited > int64(space) {
			pending[taskTyp] = date
			taskId++
		} else {
			date += waited - 1
			pending[taskTyp] = date
			taskId++
		}
	}
	return date
}
