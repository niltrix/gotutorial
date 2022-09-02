package main

import (
	"context"
	"errors"
	"fmt"
	"time"

	"golang.org/x/sync/errgroup"
)

func test1(ctx context.Context) error {
	timer := time.NewTimer(time.Second * 5)
	select {
	case <-timer.C:
		fmt.Println("test1 timer")
		return nil
	case <-ctx.Done():
		fmt.Println("test1 cancel")
		return ctx.Err()
	}
}

func test2(ctx context.Context) error {
	timer := time.NewTimer(time.Second * 2)
	select {
	case <-timer.C:
		fmt.Println("test2 timer")
		fmt.Println("Occur error test2")
		return errors.New("error test2")
	case <-ctx.Done():
		fmt.Println("test2 cancel")
		return ctx.Err()
	}
}

func main() {
	ctx, _ := context.WithCancel(context.Background())
	errGrp, errCtx := errgroup.WithContext(ctx)

	errGrp.Go(func() error { return test1(errCtx) })
	errGrp.Go(func() error { return test2(errCtx) })

	//cancel()

	if err := errGrp.Wait(); err != nil {
		fmt.Println(err.Error())
	}
}
