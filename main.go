package main

import (
	"fmt"
	"time"

	"github.com/designsbysm/syncfolders/sync"
)

func main() {
	start := time.Now()

	if err := config(); err != nil {
		fmt.Println(err)
		return
	}

	if err := sync.Go(); err != nil {
		fmt.Println(err)
		return
	}

	duration := time.Since(start)
	fmt.Println("Duration:", duration.Round(time.Millisecond))
}
