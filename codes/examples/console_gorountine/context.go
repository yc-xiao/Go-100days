package main

import (
	"context"
	"fmt"
	"time"
)

func reqTask(ctx context.Context, name string) {
	for {
		select {
		case <-ctx.Done():
			fmt.Println("stop", name)
			return
		}
	}
}

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	go reqTask(ctx, "worker1")
	time.Sleep(time.Second)
	cancel()
	time.Sleep(time.Second)
}
