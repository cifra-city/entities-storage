package cli

import (
	"context"
	"sync"
)

func runServices(ctx context.Context, wg *sync.WaitGroup) {
	run := func(f func()) {
		wg.Add(1)
		go func() {
			f()
			wg.Done()
		}()
	}

	run(func() { service.Run(ctx) })
}