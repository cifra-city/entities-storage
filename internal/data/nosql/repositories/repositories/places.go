package repositories

import (
	"context"
	"fmt"

	"github.com/cifra-city/entities-storage/internal/data/nosql/repositories"
	"github.com/cifra-city/entities-storage/internal/data/nosql/repositories/models"
	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Places interface {
	Add(ctx context.Context, place models.Place) (models.Place, error)
	AddReview(ctx context.Context, id uuid.UUID, review int) (models.Place, error)

	UpdateName(ctx context.Context, id uuid.UUID, newName string) (models.Place, error)
	UpdateDescription(ctx context.Context, id uuid.UUID, description string) (models.Place, error)
	UpdateType(ctx context.Context, id uuid.UUID, newType string) (models.Place, error)
	UpdateTags(ctx context.Context, id uuid.UUID, newTags []string) (models.Place, error)
	UpdateSchedule(ctx context.Context, id uuid.UUID, newSchedule []models.Schedule) (models.Place, error)
	UpdateLocation(ctx context.Context, id uuid.UUID, newLocation models.GeoPoint) (models.Place, error)
	UpdateCityID(ctx context.Context, id uuid.UUID, cityID uuid.UUID) (models.Place, error)

	Delete(ctx context.Context, id uuid.UUID) (models.Place, error)

	Filter() *repositories.QueryBuilder
}

type place struct {
	client     *mongo.Client
	database   *mongo.Database
	collection *mongo.Collection
}

// NewPlaces создает репозиторий для коллекции Places
func NewPlaces(uri, dbName, collectionName string) (Places, error) {
	clientOptions := options.Client().ApplyURI(uri)
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to MongoDB: %w", err)
	}

	database := client.Database(dbName)
	coll := database.Collection(collectionName)

	return &place{
		client:     client,
		database:   database,
		collection: coll,
	}, nil
}

// Add добавляет новый документ в коллекцию Places
func (p *place) Add(ctx context.Context, place models.Place) (models.Place, error) {
	_, err := p.collection.InsertOne(ctx, place)
	if err != nil {
		return models.Place{}, fmt.Errorf("failed to add place: %w", err)
	}
	return place, nil
}

// AddReview добавляет отзыв и обновляет счетчики
func (p *place) AddReview(ctx context.Context, id uuid.UUID, review int) (models.Place, error) {
	if review < 1 || review > 5 {
		return models.Place{}, fmt.Errorf("review must be between 1 and 5")
	}

	filter := bson.M{"_id": id}
	update := bson.M{
		"$inc": bson.M{
			"review_sum":   review,
			"review_count": 1,
		},
	}
	opts := options.FindOneAndUpdate().SetReturnDocument(options.After)

	var updatedPlace models.Place
	err := p.collection.FindOneAndUpdate(ctx, filter, update, opts).Decode(&updatedPlace)
	if err != nil {
		return models.Place{}, fmt.Errorf("failed to add review: %w", err)
	}
	return updatedPlace, nil
}

// UpdateName обновляет имя места
func (p *place) UpdateName(ctx context.Context, id uuid.UUID, newName string) (models.Place, error) {
	filter := bson.M{"_id": id}
	update := bson.M{"$set": bson.M{"name": newName}}
	return p.updatePlace(ctx, filter, update)
}

// UpdateDescription обновляет описание места
func (p *place) UpdateDescription(ctx context.Context, id uuid.UUID, description string) (models.Place, error) {
	filter := bson.M{"_id": id}
	update := bson.M{"$set": bson.M{"description": description}}
	return p.updatePlace(ctx, filter, update)
}

// UpdateType обновляет тип места
func (p *place) UpdateType(ctx context.Context, id uuid.UUID, newType string) (models.Place, error) {
	filter := bson.M{"_id": id}
	update := bson.M{"$set": bson.M{"type": newType}}
	return p.updatePlace(ctx, filter, update)
}

// UpdateTags обновляет теги места
func (p *place) UpdateTags(ctx context.Context, id uuid.UUID, newTags []string) (models.Place, error) {
	filter := bson.M{"_id": id}
	update := bson.M{"$set": bson.M{"tags": newTags}}
	return p.updatePlace(ctx, filter, update)
}

// UpdateSchedule обновляет расписание места
func (p *place) UpdateSchedule(ctx context.Context, id uuid.UUID, newSchedule []models.Schedule) (models.Place, error) {
	filter := bson.M{"_id": id}
	update := bson.M{"$set": bson.M{"schedule": newSchedule}}
	return p.updatePlace(ctx, filter, update)
}

// UpdateLocation обновляет местоположение
func (p *place) UpdateLocation(ctx context.Context, id uuid.UUID, newLocation models.GeoPoint) (models.Place, error) {
	filter := bson.M{"_id": id}
	update := bson.M{"$set": bson.M{"location": newLocation}}
	return p.updatePlace(ctx, filter, update)
}

// UpdateCityID обновляет ID города места
func (p *place) UpdateCityID(ctx context.Context, id uuid.UUID, cityID uuid.UUID) (models.Place, error) {
	filter := bson.M{"_id": id}
	update := bson.M{"$set": bson.M{"city_id": cityID}}
	return p.updatePlace(ctx, filter, update)
}

// Delete удаляет место из коллекции
func (p *place) Delete(ctx context.Context, id uuid.UUID) (models.Place, error) {
	filter := bson.M{"_id": id}
	var deletedPlace models.Place
	err := p.collection.FindOneAndDelete(ctx, filter).Decode(&deletedPlace)
	if err != nil {
		return models.Place{}, fmt.Errorf("failed to delete place: %w", err)
	}
	return deletedPlace, nil
}

// Filter возвращает новый экземпляр QueryBuilder
func (p *place) Filter() *repositories.QueryBuilder {
	return &repositories.QueryBuilder{
		Collection: p.collection,
		Filters:    bson.M{},
		Sort:       bson.D{},
	}
}

func (p *place) updatePlace(ctx context.Context, filter bson.M, update bson.M) (models.Place, error) {
	opts := options.FindOneAndUpdate().SetReturnDocument(options.After)
	var updatedPlace models.Place
	err := p.collection.FindOneAndUpdate(ctx, filter, update, opts).Decode(&updatedPlace)
	if err != nil {
		return models.Place{}, fmt.Errorf("failed to update place: %w", err)
	}
	return updatedPlace, nil
}
