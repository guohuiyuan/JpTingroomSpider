package main

import "math"

func Partition[T any](list []T, size int) (plist [][]T) {
	plen := int(math.Ceil(float64(len(list)) / float64(size)))
	plist = make([][]T, plen)
	for i := 0; i < plen; i++ {
		if i == plen-1 {
			plist[i] = list[i*size:]
			break
		}
		plist[i] = list[i*size : (i+1)*size]
	}
	return
}
