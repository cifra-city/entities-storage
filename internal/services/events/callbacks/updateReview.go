package callbacks

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/cifra-city/comtools/cifractx"
	"github.com/cifra-city/entities-storage/internal/config"
	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
)

type ReviewSend struct {
	PlaceID       uuid.UUID `json:"place_id"`
	UserID        uuid.UUID `json:"user_id"`
	ReviewMessage string    `json:"review_message"`
	ReviewGrade   int       `json:"review_grade"`
}

func UpdateReview(ctx context.Context, body []byte) error {
	var event ReviewSend
	err := json.Unmarshal(body, &event)
	if err != nil {
		return fmt.Errorf("failed to unmarshal event body: %w", err)
	}

	server, err := cifractx.GetValue[*config.Service](ctx, config.SERVER)
	if err != nil {
		logrus.Fatalf("failed to get server from context: %v", err)
		return err
	}

	log := server.Logger

	place, err := server.MongoDB.Places.AddReview(ctx, event.PlaceID, event.ReviewGrade)
	if err != nil {
		log.Errorf("failed to add review: %v", err)
		return err
	}

	log.Infof("Review added to place %s", place.ID)

	return nil
}
