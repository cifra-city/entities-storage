package cli

import (
	"context"
	"sync"

	service "github.com/cifra-city/entities-storage/internal/services"
	"github.com/cifra-city/entities-storage/internal/services/events"
)

func runServices(ctx context.Context, wg *sync.WaitGroup) {
	run := func(f func()) {
		wg.Add(1)
		go func() {
			f()
			wg.Done()
		}()
	}

	run(func() { events.Listener(ctx) })

	run(func() { service.Run(ctx) })
}
