package main

import (
	"CloudNative/Core"
	"context"
	"errors"
	"fmt"
	"time"
)

var count int

func EmulateTransientError(ctx context.Context) (string, error) {
	count++

	if count <= 3 {
		return "intentional fail", errors.New("error")
	} else {
		return "success", nil
	}
}

func main() {
	//serviceCircuit := func(ctx context.Context) (string, error) {
	//	return "success", nil
	//}

	//wrapped := Core.Breaker(Core.DebounceLast(serviceCircuit, time.Millisecond*100), 10)

	r := Core.Retry(EmulateTransientError, 5, 2*time.Second)

	res, err := r(context.Background())
	fmt.Println(res, err)
}
