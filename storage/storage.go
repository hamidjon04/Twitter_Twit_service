package storage

import (
	m "twit-service/storage/modgodb"

	"go.mongodb.org/mongo-driver/mongo"
)

type Twit interface {
	Twit() m.TwitRepository
	Like() m.LikeRepository
}

type twitImpl struct {
	mDB *mongo.Database
}

func NewTwitImpl(db *mongo.Database) Twit {
	return &twitImpl{
		mDB: db,
	}
}

func (t twitImpl) Twit() m.TwitRepository {
	return m.NewTwitRepository(t.mDB)
}

func (t twitImpl) Like() m.LikeRepository {
	return m.NewLikeRepository(t.mDB)
}
