package storage

import (
	"database/sql"
	m "twit-service/storage/modgodb"
	"twit-service/storage/postgres"

	"go.mongodb.org/mongo-driver/mongo"
)

type Twit interface {
	Twit() m.TwitRepository
	Like() postgres.LikeRepo
	MongoDatabase() mongo.Database
}

type twitImpl struct {
	mDB *mongo.Database
	db *sql.DB
}

func NewTwitImpl(db *sql.DB, mdb *mongo.Database) Twit {
	return &twitImpl{
		mDB: mdb,
		db: db,
	}
}

func (t twitImpl) Twit() m.TwitRepository {
	return m.NewTwitRepository(t.mDB)
}

func (t twitImpl) Like() postgres.LikeRepo {
	return postgres.NewLikeRepo(t.db)
}

func (t twitImpl) MongoDatabase()mongo.Database{
	return *t.mDB
}