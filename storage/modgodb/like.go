package mongo

import (
	"context"
	"time"
	pb "twit-service/generated/twit"
	"twit-service/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type LikeRepository interface {
	AddLike(ctx context.Context, in *models.Like) (*pb.AddLikeResp, error)
	RemoveLike(ctx context.Context, in *pb.DeleleLikeReq) error
}

type likeImpl struct {
	db *mongo.Database
}

func NewLikeRepository(db *mongo.Database) LikeRepository {
	return &likeImpl{db: db}
}

func (r likeImpl) AddLike(ctx context.Context, in *models.Like) (*pb.AddLikeResp, error) {
	collectin := r.db.Collection("likes")

	_, err := collectin.InsertOne(ctx, bson.M{
		"twit_id":    in.Twit_id,
		"clicker_id": in.Clicker_id,
		"created_at": in.CreatedAt,
		"deleted_at": "",
	})
	if err != nil {
		return nil, err
	}
	return &pb.AddLikeResp{Massage: "like added"}, nil
}

func (r likeImpl) RemoveLike(ctx context.Context, in *pb.DeleleLikeReq) error {
	collectin := r.db.Collection("likes")
	filter := bson.M{"$and": []bson.M{{"twit_id": in.TwitId}, {"clicker_id": in.ClickerId}, {"deleted_at": ""}}}

	_, err := collectin.UpdateOne(ctx, filter, bson.M{"$set": bson.M{"deleted_at": time.Now().Format("2006/01/02")}})
	return err
}
