package lib

import (
	"fmt"
	"math/rand"
)

func RandomNumberStr(length int) string {
	var result string
	for i := 0; i < length; i++ {
		result += fmt.Sprintf(`%d`, rand.Intn(9))
	}
	return result
}