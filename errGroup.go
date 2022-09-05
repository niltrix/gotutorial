package main

import (
	"context"
	"errors"
	"fmt"
	"time"

	"golang.org/x/sync/errgroup"
)

func counterWork5(ctx context.Context) error {
	timer := time.NewTimer(time.Second * 5)
	select {
	case <-timer.C:
		fmt.Println("5sec timer")
		return nil
	case <-ctx.Done():
		fmt.Println("ctx.Done in 5sec timer")
		return ctx.Err()
	}
}

func counterWork2(ctx context.Context) error {
	timer := time.NewTimer(time.Second * 2)
	select {
	case <-timer.C:
		fmt.Println("2sec timer")
		fmt.Println("Occur error 2sec timer")
		return errors.New("error 2sec timer")
	case <-ctx.Done():
		fmt.Println("ctx.Done in 2sec timer ")
		return ctx.Err()
	}
}

func main() {

	ctx, _ := context.WithCancel(context.Background())
	errGrp, errCtx := errgroup.WithContext(ctx)

	errGrp.Go(func() error { return counterWork5(errCtx) })
	errGrp.Go(func() error { return counterWork2(errCtx) })

	//cancel()

	if err := errGrp.Wait(); err != nil {
		fmt.Println(err.Error())
	}
}
