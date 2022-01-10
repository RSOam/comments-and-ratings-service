package commrat

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http"
	"time"

	"github.com/go-kit/log"
	consulapi "github.com/hashicorp/consul/api"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type database struct {
	db     *mongo.Database
	logger log.Logger
	consul consulapi.Client
}

func NewDatabase(db *mongo.Database, logger log.Logger, consul consulapi.Client) CommRatDB {
	return &database{
		db:     db,
		logger: log.With(logger, "database", "mongoDB"),
		consul: consul,
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
	_, err = dat.db.Collection("Comments").InsertOne(ctx, comment)
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
	err = dat.db.Collection("Comments").FindOne(ctx, bson.M{"_id": objectID}).Decode(&tempComment)
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
	res := dat.db.Collection("Comments").FindOneAndDelete(ctx, filter)
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
	cursor, err := dat.db.Collection("Comments").Find(ctx, bson.D{})
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
func (dat *database) GetCommentsFilter(ctx context.Context, chargerID string, userID string) ([]Comment, error) {
	var mchargerID primitive.ObjectID
	var muserID primitive.ObjectID
	var err error
	filter := bson.M{}
	tempComment := Comment{}
	tempComments := []Comment{}
	if chargerID != "" {
		mchargerID, err = primitive.ObjectIDFromHex(chargerID)
		if err != nil {
			dat.logger.Log("Error getting comments from DB: ", err.Error())
			return tempComments, err
		}
	}
	if userID != "" {
		muserID, err = primitive.ObjectIDFromHex(userID)
		if err != nil {
			dat.logger.Log("Error getting comments from DB: ", err.Error())
			return tempComments, err
		}
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if userID != "" && chargerID == "" {
		filter = bson.M{"userid": muserID}
	} else if userID == "" && chargerID != "" {
		filter = bson.M{"chargerid": mchargerID}
	} else if userID != "" && chargerID != "" {
		filter = bson.M{"chargerid": mchargerID, "userid": muserID}
	}
	cursor, err := dat.db.Collection("Comments").Find(ctx, filter)
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
func (dat *database) UpdateComment(ctx context.Context, id string, text string) error {
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		dat.logger.Log("Error updating comment: ", err.Error())
		return err
	}
	update := bson.M{
		"$set": bson.M{
			"text":     text,
			"modified": time.Now().Format(time.RFC3339),
		},
	}
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	_, err = dat.db.Collection("Comments").UpdateOne(ctx, bson.M{"_id": objectID}, update)
	if err != nil {
		dat.logger.Log("Error updating comment: ", err.Error())
		return err
	}

	return nil
}

//RATINGS
func (dat *database) CreateRating(ctx context.Context, rating float64, userID string, chargerID string) error {
	userIDmongo, err := primitive.ObjectIDFromHex(userID)
	if err != nil {
		dat.logger.Log("Error creating rating: ", err.Error())
		return err
	}
	chargerIDmongo, err := primitive.ObjectIDFromHex(chargerID)
	if err != nil {
		dat.logger.Log("Error creating rating: ", err.Error())
		return err
	}
	ratingObj := Rating{
		ChargerID: chargerIDmongo,
		UserID:    userIDmongo,
		Rating:    rating,
		Created:   time.Now().Format(time.RFC3339),
		Modified:  time.Now().Format(time.RFC3339),
	}
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	_, err = dat.db.Collection("Ratings").InsertOne(ctx, ratingObj)
	if err != nil {
		dat.logger.Log("Error inserting rating into DB: ", err.Error())
		return err
	}
	avg, err := calculateChargerRating(dat.db, dat.logger, chargerIDmongo)
	if err != nil {
		dat.logger.Log("Error calculating charger rating: ", err.Error())
		return nil
	}
	val, _ := getConsulValue(dat.consul, dat.logger, "chargersService")

	requestBody, err := json.Marshal(PostChargerUpdateRequest{
		AverageRating: avg,
	})
	if err != nil {
		return err
	}
	chargersUri := val + "/chargers/"
	ratingChan := make(chan *http.Response)
	dat.logger.Log("Sending async request")
	go AsyncPost(chargersUri+chargerID, requestBody, ratingChan, dat.logger)
	//chargersResponse := <-ratingChan
	//defer chargersResponse.Body.Close()
	//updateChargerRating(val, dat.logger, chargerID, avg)
	dat.logger.Log("Finished creating new rating")
	return nil
}
func (dat *database) GetRating(ctx context.Context, id string) (Rating, error) {
	tempRating := Rating{}
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		dat.logger.Log("Error getting rating from DB: ", err.Error())
		return tempRating, err
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	err = dat.db.Collection("Ratings").FindOne(ctx, bson.M{"_id": objectID}).Decode(&tempRating)
	if err != nil {
		dat.logger.Log("Error getting rating from DB: ", err.Error())
		return tempRating, err
	}
	return tempRating, nil
}
func (dat *database) DeleteRating(ctx context.Context, id string) error {
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		dat.logger.Log("Error deleting rating from DB: ", err.Error())
		return err
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	filter := bson.M{"_id": objectID}
	res := dat.db.Collection("Ratings").FindOneAndDelete(ctx, filter)
	if res.Err() == mongo.ErrNoDocuments {
		dat.logger.Log("Error deleting rating from DB: ", err.Error())
		return err
	}
	return nil
}
func (dat *database) GetRatings(ctx context.Context) ([]Rating, error) {
	tempRating := Rating{}
	tempRatings := []Rating{}
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	cursor, err := dat.db.Collection("Ratings").Find(ctx, bson.D{})
	if err != nil {
		dat.logger.Log("Error getting ratings from DB: ", err.Error())
		return tempRatings, err
	}
	defer cursor.Close(ctx)
	for cursor.Next(ctx) {
		err := cursor.Decode(&tempRating)
		if err != nil {
			dat.logger.Log("Error getting ratings from DB: ", err.Error())
			return tempRatings, err
		}
		tempRatings = append(tempRatings, tempRating)
	}
	return tempRatings, nil
}
func (dat *database) UpdateRating(ctx context.Context, id string, rating float64) error {
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		dat.logger.Log("Error updating rating: ", err.Error())
		return err
	}
	update := bson.M{
		"$set": bson.M{
			"rating":   rating,
			"modified": time.Now().Format(time.RFC3339),
		},
	}
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	_, err = dat.db.Collection("Ratings").UpdateOne(ctx, bson.M{"_id": objectID}, update)
	if err != nil {
		dat.logger.Log("Error updating rating: ", err.Error())
		return err
	}

	return nil
}
func (dat *database) GetRatingsFilter(ctx context.Context, chargerID string, userID string) ([]Rating, error) {
	var mchargerID primitive.ObjectID
	var muserID primitive.ObjectID
	var err error
	filter := bson.M{}
	tempRating := Rating{}
	tempRatings := []Rating{}
	if chargerID != "" {
		mchargerID, err = primitive.ObjectIDFromHex(chargerID)
		if err != nil {
			dat.logger.Log("Error getting ratings from DB: ", err.Error())
			return tempRatings, err
		}
	}
	if userID != "" {
		muserID, err = primitive.ObjectIDFromHex(userID)
		if err != nil {
			dat.logger.Log("Error getting ratings from DB: ", err.Error())
			return tempRatings, err
		}
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if userID != "" && chargerID == "" {
		filter = bson.M{"userid": muserID}
	} else if userID == "" && chargerID != "" {
		filter = bson.M{"chargerid": mchargerID}
	} else if userID != "" && chargerID != "" {
		filter = bson.M{"chargerid": mchargerID, "userid": muserID}
	}
	cursor, err := dat.db.Collection("Ratings").Find(ctx, filter)
	if err != nil {
		dat.logger.Log("Error getting ratings from DB: ", err.Error())
		return tempRatings, err
	}
	defer cursor.Close(ctx)
	for cursor.Next(ctx) {
		err := cursor.Decode(&tempRating)
		if err != nil {
			dat.logger.Log("Error getting ratings from DB: ", err.Error())
			return tempRatings, err
		}
		tempRatings = append(tempRatings, tempRating)
	}
	return tempRatings, nil
}
func calculateChargerRating(db *mongo.Database, logger log.Logger, chargerID primitive.ObjectID) (float64, error) {
	tempRating := Rating{}
	tempRatings := []Rating{}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	cursor, err := db.Collection("Ratings").Find(ctx, bson.M{"chargerid": chargerID})
	if err != nil {
		logger.Log("Error getting ratings from DB: ", err.Error())
		return 0, err
	}
	defer cursor.Close(ctx)
	for cursor.Next(ctx) {
		err := cursor.Decode(&tempRating)
		if err != nil {
			logger.Log("Error getting ratings from DB: ", err.Error())
			return 0, err
		}
		tempRatings = append(tempRatings, tempRating)
	}

	var average float64 = 0.0
	for _, el := range tempRatings {
		average += el.Rating
	}
	average = average / float64(len(tempRatings))
	return average, nil
}
func updateChargerRating(chargersAddr string, logger log.Logger, chargerID string, average float64) (string, error) {
	requestBody, err := json.Marshal(PostChargerUpdateRequest{
		AverageRating: average,
	})
	if err != nil {
		return "", err
	}
	client := &http.Client{}
	chargersUri := chargersAddr + "/chargers/"
	req, err := http.NewRequest(http.MethodPut, chargersUri+chargerID, bytes.NewBuffer(requestBody))
	if err != nil {
		return "", err
	}
	req.Header.Set("Content-Type", "application/json; charset=utf-8")
	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	tempResponse := PostChargerUpdateResponse{}
	err = json.NewDecoder(resp.Body).Decode(&tempResponse)
	if err != nil {
		return "", err
	}
	if tempResponse.Status != "Ok" {
		return tempResponse.Status, nil
	}
	client.CloseIdleConnections()
	return "Ok", nil
}

func AsyncPost(url string, body []byte, rc chan *http.Response, logger log.Logger) {
	response, err := http.Post(url, "application/json", bytes.NewReader(body))
	if err != nil {
		logger.Log("Error sending async request: " + err.Error())
	}
	logger.Log("Async request finished")
	rc <- response
}

func getConsulValue(consul consulapi.Client, logger log.Logger, key string) (string, error) {
	kv := consul.KV()
	keyPair, _, err := kv.Get(key, nil)
	if err != nil {
		logger.Log("msg", "Failed getting consul key")
		return "", err
	}
	return string(keyPair.Value), nil
}
