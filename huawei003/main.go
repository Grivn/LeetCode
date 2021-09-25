package main

import (
	"fmt"
)

func main() {
	//passengers, _ := input()

	var name string

	var start int

	var stop int

	fmt.Scanf("%s,%d,%d",&name, &start, &stop)

	fmt.Println(name)
	fmt.Println(start)
	fmt.Println(stop)

	//fmt.Println(passengers)
	//var passengers []*passenger

	//r := bufio.NewReader(os.Stdin)
	//for {
	//	input, err := r.ReadString('\n')
	//
	//	if err == io.EOF {
	//		return
	//	}
	//
	//	//var name, start, stop string
	//
	//	res := strings.Split(input, ",")
	//
	//	fmt.Println(res)
	//
	//}

}

type passenger struct {
	name string

	start int

	stop int
}
