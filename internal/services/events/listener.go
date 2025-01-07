package events

import (
	"context"

	"github.com/cifra-city/comtools/cifractx"
	"github.com/cifra-city/entities-storage/internal/config"
	"github.com/cifra-city/entities-storage/internal/services/events/callbacks"
	"github.com/sirupsen/logrus"
)

const (
	accountCreateQ = "account.create"
)

func Listener(ctx context.Context) {
	server, err := cifractx.GetValue[*config.Service](ctx, config.SERVER)
	if err != nil {
		logrus.Fatalf("failed to get server from context: %v", err)
	}

	err = server.Broker.Listen(ctx, server.Logger, accountCreateQ, "entities-storage.events", callbacks.UpdateReview)
	if err != nil {
		logrus.Fatalf("Listener encountered an error: %v", err)
	}
}
