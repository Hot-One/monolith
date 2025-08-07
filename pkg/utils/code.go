package utils

import (
	mr "math/rand"
)

func GenerateOrderNumber() int {
	return 10000000 + mr.Intn(90000000)
}
