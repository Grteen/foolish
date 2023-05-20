package uuid

import (
	"fmt"
	"math/rand"
	"time"
)

func GetVerify() string {
	rand.Seed(time.Now().UnixNano())
	code := rand.Intn(10000)
	return fmt.Sprintf("%04d", code)
}
