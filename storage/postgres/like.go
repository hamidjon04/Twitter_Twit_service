package postgres

import (
	"database/sql"
	"fmt"
	"log"
	"twit-service/generated/twit"
	m "twit-service/storage/modgodb"

	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/mongo"
)

type LikeRepo interface {
	AddLike(Coll mongo.Database, req *twit.AddLikeReq) (*twit.AddLikeResp, error)
	DeleteLike(Coll mongo.Database, req *twit.DeleteLikeReq) (*twit.Message, error)
}

type likeImpl struct {
	DB *sql.DB
}

func NewLikeRepo(db *sql.DB) LikeRepo {
	return &likeImpl{
		DB: db,
	}
}

func (L *likeImpl) AddLike(DB mongo.Database, req *twit.AddLikeReq) (*twit.AddLikeResp, error) {
	cheack := m.IsTwitExists(DB.Collection("twits"), req.TwitId)
	if !cheack {
		return nil, fmt.Errorf("bunday twit mavjud emas")
	}
	id := uuid.NewString()
	query := `
				INSERT INTO likes(
					id, twit_id, clicker_id)
				VALUES
					($1, $2, $3)`
	_, err := L.DB.Exec(query, id, req.TwitId, req.ClickerId)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return &twit.AddLikeResp{
		Massage: "Like",
	}, nil
}

func (L *likeImpl) DeleteLike(DB mongo.Database, req *twit.DeleteLikeReq) (*twit.Message, error) {
	cheack := m.IsTwitExists(DB.Collection("twits"), req.TwitId)
	if !cheack {
		return nil, fmt.Errorf("bunday twit mavjud emas")
	}
	query := `
					DELETE FROM 
						likes
					WHERE 
						twit_id = $1`
	_, err := L.DB.Exec(query, req.TwitId)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return &twit.Message{
		Message: "Unlike",
	}, nil
}
