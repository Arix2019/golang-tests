package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	ctx := context.Background()
	sleepAndTalk(ctx, 5*time.Second, "Olar!")
}

func sleepAndTalk(ctx context.Context, d time.Duration, msg string) {
	time.Sleep(d)
	fmt.Println(msg)
}
