package repositories

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Filter interface {
	ById(id string) *QueryBuilder
	ByTags(tags []string) *QueryBuilder
	ByCityId(cityId string) *QueryBuilder
	ByServiceType(serviceType string) *QueryBuilder
	ByType(placeType string) *QueryBuilder
	ByReview(reviewAvg float64) *QueryBuilder
	ByReviewCount(reviewCount int) *QueryBuilder
	BySchedule(dayOfWeek time.Weekday, open time.Time, close time.Time) *QueryBuilder
	ByLocation(Latitude1 float64, Latitude2 float64, Longitude1 float64, Longitude2 float64) *QueryBuilder

	Execute(ctx context.Context, result interface{}) error
}

type QueryBuilder struct {
	Collection *mongo.Collection
	Filters    bson.M
	Sort       bson.D
	Limit      int64
	Skip       int64
}

// NewQueryBuilder возвращает новый экземпляр PlaceQueryBuilder
func NewQueryBuilder(collection *mongo.Collection) *QueryBuilder {
	return &QueryBuilder{
		Collection: collection,
		Filters:    bson.M{},
		Sort:       bson.D{},
	}
}

func (qb *QueryBuilder) ById(id string) *QueryBuilder {
	qb.Filters["_id"] = id
	return qb
}

func (qb *QueryBuilder) ByTags(tags []string) *QueryBuilder {
	qb.Filters["tags"] = bson.M{"$all": tags}
	return qb
}

func (qb *QueryBuilder) ByCityId(cityId string) *QueryBuilder {
	qb.Filters["city_id"] = cityId
	return qb
}

func (qb *QueryBuilder) ByServiceType(serviceType string) *QueryBuilder {
	qb.Filters["service_type"] = serviceType
	return qb
}

func (qb *QueryBuilder) ByType(placeType string) *QueryBuilder {
	qb.Filters["type"] = placeType
	return qb
}

func (qb *QueryBuilder) ByReview(reviewAvg float64) *QueryBuilder {
	qb.Filters["review_avg"] = bson.M{"$gte": reviewAvg} // Найти места с рейтингом >= reviewAvg
	return qb
}

func (qb *QueryBuilder) ByReviewCount(reviewCount int) *QueryBuilder {
	qb.Filters["review_count"] = bson.M{"$gte": reviewCount} // Найти места с количеством отзывов >= reviewCount
	return qb
}

func (qb *QueryBuilder) BySchedule(dayOfWeek time.Weekday, open, close time.Time) *QueryBuilder {
	qb.Filters["schedule"] = bson.M{
		"$elemMatch": bson.M{
			"day_of_week": dayOfWeek.String(),
			"open":        bson.M{"$lte": open.Format("15:04")},  // Открыто не позже `open`
			"close":       bson.M{"$gte": close.Format("15:04")}, // Закрыто не раньше `close`
		},
	}
	return qb
}

func (qb *QueryBuilder) ByLocation(Latitude1, Latitude2, Longitude1, Longitude2 float64) *QueryBuilder {
	qb.Filters["location"] = bson.M{
		"$geoWithin": bson.M{
			"$geometry": bson.M{
				"type": "Polygon",
				"coordinates": [][]float64{
					{Longitude1, Latitude1},
					{Longitude2, Latitude1},
					{Longitude2, Latitude2},
					{Longitude1, Latitude2},
					{Longitude1, Latitude1},
				},
			},
		},
	}
	return qb
}

func (qb *QueryBuilder) Execute(ctx context.Context, result interface{}) error {
	op := options.FindOptions{
		Sort:  qb.Sort,
		Limit: &qb.Limit,
		Skip:  &qb.Skip,
	}

	cursor, err := qb.Collection.Find(ctx, qb.Filters, &op)
	if err != nil {
		return err
	}
	defer cursor.Close(ctx)

	return cursor.All(ctx, result)
}
