package util

import (
	"fmt"
	"time"
)

func Timed(f func()) {
	t0 := time.Now()
	f()
	fmt.Println("Operation took: ", time.Since(t0))
}
