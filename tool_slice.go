package wfgotool

import (
	"errors"
)

type s1 struct {
	g1 int
}
type i1 interface {
	Up()
}

func (s s1) Up() {
	s.g1++
}

func FindInts(slicei []int, target int) (tgt int, err error, count int) {
	count = 0
	for _, num := range slicei {
		if num == target {
			count++
		}
	}
	if count > 0 {
		tgt = target
		err = nil
	} else {
		err = errors.New("目标值在切片中未找到")
	}
	return tgt, err, count
}
