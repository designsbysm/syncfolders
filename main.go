package main

import (
	"fmt"
	"time"

	"github.com/designsbysm/syncfolders/sync"
)

func main() {
	start := time.Now()

	if err := config(); err != nil {
		panic(err)
	}

	if err := sync.Go(); err != nil {
		panic(err)
	}

	duration := time.Since(start)
	fmt.Println("Duration:", duration.Round(time.Millisecond))
}
