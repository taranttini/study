package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	ctx := context.Background()
	//ctx, cancel := context.WithCancel(ctx) // done
	//ctx, cancel := context.WithTimeout(ctx, time.Second*3) // cancel
	ctx, cancel := context.WithTimeout(ctx, time.Second*7) // done
	defer cancel()
	bookHotel(ctx)
}

func bookHotel(ctx context.Context) {
	select {
	case <-ctx.Done():
		fmt.Println("Hotel booking cancelled. Timeout reached.")
		return
	case <-time.After(time.Second * 5):
		fmt.Println("Hotel booked.")
		return
	}
}
