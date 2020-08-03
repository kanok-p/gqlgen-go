package store

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/readpref"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	"github.com/jinzhu/copier"

	"graphql-gen/graph/model"
)

type Store struct {
	db             *mongo.Client
	dbName         string
	collectionName string
}

type id struct {
	ObjectID primitive.ObjectID `bson:"_id"`
}

func New(mongoEndpoint, dbName, collectionName string) *Store {
	clientOptions := options.Client().ApplyURI(mongoEndpoint)

	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)

	db, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	if err = db.Ping(ctx, readpref.Primary()); err != nil {
		log.Fatal(err)
	}

	return &Store{
		dbName:         dbName,
		collectionName: collectionName,
		db:             db,
	}
}

func (s *Store) collection() *mongo.Collection {
	return s.db.Database(s.dbName).Collection(s.collectionName)
}

func (s *Store) List(ctx context.Context, offset, limit int) (int, []*model.Response, error) {
	total, err := s.collection().CountDocuments(ctx, bson.M{})
	if err != nil {
		return int(total), nil, err
	}

	cursor, err := s.collection().Find(ctx, bson.M{}, options.Find().SetLimit(int64(limit)).SetSkip(int64(offset)))
	if err != nil {
		return int(total), nil, err
	}
	defer func() { _ = cursor.Close(ctx) }()

	list := make([]*model.Response, 0)
	for cursor.Next(ctx) {
		resp := &model.Response{}

		if err := cursor.Decode(resp); err != nil {
			return int(total), nil, err
		}

		id := &id{}
		if err := cursor.Decode(id); err != nil {
			return int(total), nil, err
		}

		resp.ID = id.ObjectID.Hex()
		list = append(list, resp)
	}

	return int(total), list, nil
}

func (s *Store) Get(ctx context.Context, ID string) (*model.Response, error) {
	resp := &model.Response{
		ID: ID,
	}
	oID, err := primitive.ObjectIDFromHex(ID)
	if err != nil {
		return resp, err
	}

	if err := s.collection().FindOne(ctx, bson.D{{"_id", oID}}).Decode(resp); err != nil {
		return nil, err
	}
	return resp, nil
}

func (s *Store) Create(ctx context.Context, input *model.CreateInput) (*model.Response, error) {
	resp := &model.Response{}

	r, err := s.collection().InsertOne(ctx, input)

	if err != nil {
		return resp, err
	}
	_ = copier.Copy(&resp, &input)
	resp.ID = r.InsertedID.(primitive.ObjectID).Hex()

	return resp, nil
}

func (s *Store) Update(ctx context.Context, input *model.UpdateInput) (*model.Response, error) {
	resp := &model.Response{
		ID: input.ID,
	}
	oID, err := primitive.ObjectIDFromHex(input.ID)
	if err != nil {
		return resp, err
	}

	filter := bson.D{{"_id", oID}}
	update := bson.D{{"$set", input}}

	_, err = s.collection().UpdateOne(ctx, filter, update, options.Update().SetUpsert(true))
	if err != nil {
		return resp, err
	}
	_ = copier.Copy(&resp, &input)

	return resp, nil
}

func (s *Store) Delete(ctx context.Context, id string) (*model.Response, error) {
	resp := &model.Response{
		ID: id,
	}

	oID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return resp, err
	}

	_, err = s.collection().DeleteOne(ctx, bson.M{"_id": oID})
	if err != nil {
		return resp, err
	}

	return resp, nil
}
