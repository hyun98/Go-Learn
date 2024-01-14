package main

import (
	"CloudNative/Core"
	"context"
	"time"
)

func main() {
	serviceCircuit := func(ctx context.Context) (string, error) {
		return "success", nil
	}

	wrapped := Core.Breaker(Core.DebounceLast(serviceCircuit, time.Millisecond*100), 10)

}
