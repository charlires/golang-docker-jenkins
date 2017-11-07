package sample

import (
	"fmt"
)

func Hi(name string) string {
	return fmt.Sprintf("Hi, %s", name)
}

func Sum(a, b int) int {
	return a + b
}
