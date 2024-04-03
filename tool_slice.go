package wfslice

import (
	"errors"
)

// FindInts 查找切片里相同的int
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
