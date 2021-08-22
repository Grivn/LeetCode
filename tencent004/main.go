package main

import (
	"fmt"
	"io"
)

type layer struct {
	state map[int]*unit
}

type unit struct {
	id     int
	count  float64
	remain float64
}

func main() {
	for {
		n, t, err := input()

		if err == io.EOF {
			break
		}

		layers := generate(n)

		start(layers, t)

		total := cal(layers)

		//fmt.Printf("\n")
		fmt.Println(total)
	}
}

func input() (int, int, error) {
	var n, t int
	_, err := fmt.Scan(&n, &t)
	if err != nil {
		return 0, 0, err
	}

	return n, t, nil
}

func generate(n int) []*layer {
	var res []*layer
	for i:=0; i<n; i++ {
		currentLayer := &layer{state: make(map[int]*unit)}

		for j:=0; j<i+1; j++ {
			id := j+1
			currentLayer.state[id] = &unit{id: id, count: 0, remain: 0}
		}

		res = append(res, currentLayer)
	}
	return res
}

func start(layers []*layer, t int) {
	for {
		t--
		first := layers[0].state[1]
		first.remain += float64(1)

		for index, lay := range layers {
			for id, u := range lay.state {
				if u.remain == 0 {
					continue
				}

				if u.count < float64(1) {
					need := float64(1)-u.count

					if u.remain >= need {
						u.count = float64(1)
						u.remain = u.remain-need
					} else {
						u.count += u.remain
						u.remain = 0
					}
				}

				if u.count == float64(1) && u.remain > 0 {
					if index+1 == len(layers) {
						u.remain = 0
						continue
					}

					push := u.remain/2

					next1 := layers[index+1].state[id]
					next2 := layers[index+1].state[id+1]

					next1.remain += push
					next2.remain += push

					u.remain = 0
				}
			}
		}

		if t == 0 {
			break
		}
	}
}

func cal(layers []*layer) int {
	count := 0
	for _, lay := range layers {
		for _, u := range lay.state {
			//fmt.Printf("id %d count %v remain %v\n", u.id, u.count, u.remain)
			if u.count == float64(1) {
				count++
			}
		}
	}

	return count
}
