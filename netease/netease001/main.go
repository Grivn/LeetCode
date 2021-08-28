package main

import (
	"fmt"
	"io"
)

func main() {
	for {
		inputs, err := input()
		if err == io.EOF {
			break
		}

		for _, s := range inputs {
			hp, success := findGetThrough(s)

			if success {
				fmt.Println(hp)
			} else {
				fmt.Println(-1)
			}
		}
	}
}

func findGetThrough(s *stream) (int, bool) {
	me := s.me

	me.hp = 0

	lowestHP := 0

	for _, monster := range s.monsters {
		if me.attack <= monster.defence {
			return -1, false
		}

		myActualAttack := me.attack - monster.defence

		monsterActualAttack := monster.attack - me.defence

		if monsterActualAttack <= 0 {
			monsterActualAttack = 0
		}

		for {
			monster.hp -= myActualAttack

			if monster.hp <= 0 {
				me.hp += 0-monster.hp
				break
			}

			me.hp -= monsterActualAttack

			if lowestHP > me.hp {
				lowestHP = me.hp
			}
		}
	}

	return 0-lowestHP+1, true
}

func preCheck(s *stream) bool {
	me := s.me

	for _, monster := range s.monsters {
		if me.attack <= monster.defence {
			return false
		}
	}

	return true
}

type stream struct {
	me *object

	monsters []*object
}

type object struct {
	hp      int
	attack  int
	defence int
}

func input() ([]*stream, error) {
	var T int

	_, err := fmt.Scan(&T)

	if err != nil {
		return nil, err
	}

	res := make([]*stream, T)

	for i:=0; i<T; i++ {
		part := &stream{me: &object{}}

		var N int

		_, _ = fmt.Scan(&N)

		_, _ = fmt.Scan(&part.me.attack, &part.me.defence)

		monsters := make([]*object, N)
		for j:=0; j<N; j++ {
			monsters[j] = &object{}
			_, _ = fmt.Scan(&monsters[j].attack, &monsters[j].defence, &monsters[j].hp)
		}

		part.monsters = monsters

		res[i] = part
	}

	return res, nil
}
