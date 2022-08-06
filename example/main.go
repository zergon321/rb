package main

import (
	"fmt"
	"strconv"

	"github.com/zergon321/rb"
)

const (
	DataLength = 1000
)

func main() {
	tree := rb.NewTree[int, string]()

	for i := 0; i < DataLength; i++ {
		tree.Insert(i, strconv.Itoa(i))
	}

	tree.Traverse(func(currentKey int, currentValue string) error {
		fmt.Println(currentKey, currentValue)
		return nil
	})
}
