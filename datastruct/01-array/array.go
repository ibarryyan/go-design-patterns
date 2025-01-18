package array

import "fmt"

type Array struct {
	Items  []int32
	Length int32
}

func NewArray(length int32, items ...int32) *Array {
	if int32(len(items)) > length {
		return nil
	}

	newItems := make([]int32, length)
	for i := range len(items) {
		newItems[i] = items[i]
	}
	return &Array{Items: newItems, Length: length}
}

func (arr *Array) Print() {
	for _, item := range arr.Items {
		fmt.Printf("%d,", item)
	}
}
