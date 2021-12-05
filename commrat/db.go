package commrat

import (
	"context"
	"time"

	"github.com/go-kit/log"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type database struct {
	db     *mongo.Database
	logger log.Logger
}

func NewDatabase(db *mongo.Database, logger log.Logger) CommRatDB {
	return &database{
		db:     db,
		logger: log.With(logger, "database", "mongoDB"),
	}
}

func (dat *database) CreateComment(ctx context.Context, text string, userID string, chargerID string) error {
	userIDmongo, err := primitive.ObjectIDFromHex(userID)
	if err != nil {
		dat.logger.Log("Error creating comment: ", err.Error())
		return err
	}
	chargerIDmongo, err := primitive.ObjectIDFromHex(chargerID)
	if err != nil {
		dat.logger.Log("Error creating comment: ", err.Error())
		return err
	}
	comment := Comment{
		ChargerID: chargerIDmongo,
		UserID:    userIDmongo,
		Text:      text,
		Created:   time.Now().Format(time.RFC3339),
		Modified:  time.Now().Format(time.RFC3339),
	}
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	_, err = dat.db.Collection("comments").InsertOne(ctx, comment)
	if err != nil {
		dat.logger.Log("Error inserting comment into DB: ", err.Error())
		return err
	}
	return nil
}
func (dat *database) GetComment(ctx context.Context, id string) (Comment, error) {
	tempComment := Comment{}
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		dat.logger.Log("Error getting comment from DB: ", err.Error())
		return tempComment, err
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	err = dat.db.Collection("comments").FindOne(ctx, bson.M{"_id": objectID}).Decode(&tempComment)
	if err != nil {
		dat.logger.Log("Error getting comment from DB: ", err.Error())
		return tempComment, err
	}
	return tempComment, nil
}
func (dat *database) DeleteComment(ctx context.Context, id string) error {
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		dat.logger.Log("Error deleting comment from DB: ", err.Error())
		return err
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	filter := bson.M{"_id": objectID}
	res := dat.db.Collection("comments").FindOneAndDelete(ctx, filter)
	if res.Err() == mongo.ErrNoDocuments {
		dat.logger.Log("Error deleting comment from DB: ", err.Error())
		return err
	}
	return nil
}
func (dat *database) GetComments(ctx context.Context) ([]Comment, error) {
	tempComment := Comment{}
	tempComments := []Comment{}
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	cursor, err := dat.db.Collection("comments").Find(ctx, bson.D{})
	if err != nil {
		dat.logger.Log("Error getting comments from DB: ", err.Error())
		return tempComments, err
	}
	defer cursor.Close(ctx)
	for cursor.Next(ctx) {
		err := cursor.Decode(&tempComment)
		if err != nil {
			dat.logger.Log("Error getting comments from DB: ", err.Error())
			return tempComments, err
		}
		tempComments = append(tempComments, tempComment)
	}
	return tempComments, nil
}
