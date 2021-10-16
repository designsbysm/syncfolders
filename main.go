package main

import (
	"time"

	"github.com/designsbysm/syncfolders/sync"
	"github.com/designsbysm/timber/v2"
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
	timber.Info(duration.Round(time.Millisecond))
}
