package util

import (
	"repeater/typies"
)

func GenerateId(data []typies.RepeatData) (id int) {
	high := 0

	for _, val := range data {
		if high <= val.Id {
			high = val.Id + 1
		}
	}

	return high
}
