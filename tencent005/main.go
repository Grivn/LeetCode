package main

import (
	"fmt"
	"io"
)

func main() {
	for {
		myLand, err := input()
		if err == io.EOF {
			break
		}

		if len(myLand) == 1 {
			if myLand[0][0] == 1 {
				myLand[0][0] = 2
			}
			fmt.Println(myLand[0][0])
			continue
		}

		for x, raw := range myLand {
			for y, val := range raw {
				if val == 1 || val == 2 {
					continue
				}

				p := pos{x, y}

				var tmp land
				for _, myRaw := range myLand {
					var tmpRaw []int

					for _, myVal := range myRaw {
						tmpRaw = append(tmpRaw, myVal)
					}

					tmp = append(tmp, tmpRaw)
				}

				if !findExit(p, &tmp) {
					myLand[x][y] = 2
					for tmpX, tmpCheckRaw := range tmp {
						for tmpY, tmpCheckVal := range tmpCheckRaw {
							if tmpCheckVal == 3 {
								myLand[tmpX][tmpY] = 2
							}
						}
					}
				}
			}
		}

		for _, raw := range myLand {
			for _, val := range raw {
				fmt.Printf("%d ", val)
			}
			fmt.Printf("\n")
		}
	}
}

type pos struct {
	x int
	y int
}

type land [][]int

func findExit(p pos, myLand *land) bool {

	if p.x < 0 || p.y < 0 || p.x >= len(*myLand) || p.y >= len((*myLand)[0]) {
		return true
	}

	if (*myLand)[p.x][p.y] == 1 || (*myLand)[p.x][p.y] == 2 {
		return false
	}

	if p.x == 0 || p.y == 0 || p.x == len(*myLand)-1 || p.y == len((*myLand)[0])-1 {
		return true
	}
	if (*myLand)[p.x][p.y] == 3 {
		return false
	}
	(*myLand)[p.x][p.y] = 3

	rPos := pos{x: p.x, y: p.y+1}
	lPos := pos{x: p.x, y: p.y-1}
	uPos := pos{x: p.x-1, y: p.y}
	dPos := pos{x: p.x+1, y: p.y}

	return findExit(dPos, myLand) || findExit(rPos, myLand) || findExit(lPos, myLand) || findExit(uPos, myLand)
}

func input() (land, error) {
	var n int
	_, err := fmt.Scan(&n)
	if err != nil {
		return nil, err
	}

	var myLand land

	myLand = make([][]int, n)

	for i:=0; i<n; i++ {
		raw := make([]int, n)
		for j:=0; j<n; j++ {
			var val int
			_, err = fmt.Scan(&val)
			if err != nil {
				return nil, err
			}

			raw[j] = val
		}
		myLand[i] = raw
	}

	return myLand, nil
}