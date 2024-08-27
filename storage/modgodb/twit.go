package mongo

import (
	"context"
	"time"
	pb "twit-service/generated/twit"
	"twit-service/models"

	"github.com/google/uuid"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type TwitRepository interface {
	CreateTwit(ctx context.Context, in *models.Twit) (*pb.CreateTwitResp, error)
	UpdateTwit(ctx context.Context, in *models.Twit) (*pb.UpdateTwitResp, error)
	DeleteTwit(ctx context.Context, id string) error
	GetTwit(ctx context.Context, userId string) (*pb.GetTwitsResp, error)
	GetFollowerTwit(ctx context.Context, id []string) (*pb.GetTwitsResp, error)
}

type twitImpl struct {
	db *mongo.Database
}

func NewTwitRepository(db *mongo.Database) TwitRepository {
	return &twitImpl{db: db}
}

func (r twitImpl) CreateTwit(ctx context.Context, in *models.Twit) (*pb.CreateTwitResp, error) {
	collectin := r.db.Collection("twits")
	in.CreatedAt = time.Now().Format("2006/01/02")
	in.UpdatedAt = time.Now().Format("2006/01/02")

	in.ID = uuid.NewString()

	_, err := collectin.InsertOne(ctx, bson.M{
		"_id":        in.ID,
		"user_id":    in.UserID,
		"content":    in.Content,
		"media":      in.Media,
		"deleted_at": in.DeletedAt,
		"created_at": in.CreatedAt,
		"updated_at": in.UpdatedAt,
	})
	if err != nil {
		return nil, err
	}
	return &pb.CreateTwitResp{
		Id:      in.ID,
		Message: "creating new"}, nil
}

func (r twitImpl) UpdateTwit(ctx context.Context, in *models.Twit) (*pb.UpdateTwitResp, error) {
	collectin := r.db.Collection("twits")

	in.UpdatedAt = time.Now().Format("2006/01/02")
	filter := bson.M{"$and": []bson.M{
		{"_id": in.ID},
		{"deleted_at": ""},
	}}

	update := bson.M{"$set": bson.M{
		"content":    in.Content,
		"media":      in.Media,
		"user_id":    in.UserID,
		"updated_at": in.UpdatedAt,
	}}

	var resp *pb.UpdateTwitResp
	err := collectin.FindOneAndUpdate(ctx, filter, update).Decode(&resp)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (r twitImpl) DeleteTwit(ctx context.Context, id string) error {
	collectin := r.db.Collection("twits")

	filter := bson.M{"$and": []bson.M{
		{"_id": id},
		{"deleted_at": ""},
	}}

	update := bson.M{"$set": bson.M{
		"deleted_at": time.Now().Format("2006/01/02"),
	}}

	_, err := collectin.UpdateOne(ctx, filter, update)
	return err

}

func (r twitImpl) GetTwit(ctx context.Context, userId string) (*pb.GetTwitsResp, error) {
	filter := bson.M{"$and": []bson.M{{"user_id": userId}, {"deleted_at": ""}}}

	collectin := r.db.Collection("twits")
	cursor, err := collectin.Find(ctx, filter)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)
	var twits []*pb.Twit
	for cursor.Next(ctx) {
		var twit models.Twit
		err := cursor.Decode(&twit)
		if err != nil {
			return nil, err
		}
		var twitsResp pb.Twit

		twitsResp.Id = twit.ID
		twitsResp.UserId = twit.UserID
		twitsResp.Content = twit.Content
		twitsResp.Media = twit.Media
		twitsResp.CreatedAt = twit.CreatedAt
		twitsResp.UpdatedAt = twit.UpdatedAt

		twits = append(twits, &twitsResp)
	}

	return &pb.GetTwitsResp{Twits: twits}, nil
}

func (r twitImpl) GetFollowerTwit(ctx context.Context, id []string) (*pb.GetTwitsResp, error) {

	collectin := r.db.Collection("twits")
	var twits []*pb.Twit
	for _, v := range id {
		filter := bson.M{"$and": []bson.M{
			{"user_id": string(v)},
			{"deleted_at": ""},
		}}
		cursor, err := collectin.Find(ctx, filter)
		if err != nil {
			return nil, err
		}
		defer cursor.Close(ctx)
		for cursor.Next(ctx) {
			var twit models.Twit
			err := cursor.Decode(&twit)
			if err != nil {
				return nil, err
			}
			var twitsResp pb.Twit
			twitsResp.Id = twit.ID
			twitsResp.UserId = twit.UserID
			twitsResp.Content = twit.Content
			twitsResp.Media = twit.Media
			twitsResp.CreatedAt = twit.CreatedAt
			twitsResp.UpdatedAt = twit.UpdatedAt
			twits = append(twits, &twitsResp)
		}

	}
	return &pb.GetTwitsResp{Twits: twits}, nil
}
