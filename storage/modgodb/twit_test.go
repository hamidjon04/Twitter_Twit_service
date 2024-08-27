package mongo

import (
	"context"
	"testing"
	"twit-service/models"

	"github.com/stretchr/testify/assert"
)

func TestCreateTwit(t *testing.T) {

	mDB, err := ConnectMongo()
	assert.NoError(t, err)

	repo := NewTwitRepository(mDB)
	req := models.Twit{
		UserID:  "user",
		Content: "Hello, World!",
		Media:   "https://example.com/image.jpg",
	}

	res, err := repo.CreateTwit(context.Background(), &req)
	assert.NoError(t, err)
	assert.NotEmpty(t, res)
}

func TestUpdateTwit(t *testing.T) {
	mDB, err := ConnectMongo()
    assert.NoError(t, err)

    repo := NewTwitRepository(mDB)
    req := models.Twit{
        ID:      "8677fc69-2c73-4940-a7d5-26a04aa32608",
        UserID:  "user",
        Content: "Hello, Updated World!",
        Media:   "https://example.com/updated_image.jpg",
    }

    _, err = repo.CreateTwit(context.Background(), &req)
    assert.NoError(t, err)

    newReq := models.Twit{
        ID:      "8677fc69-2c73-4940-a7d5-26a04aa32608",
        UserID:  "user",
        Content: "Hello, Updated Again World!",
        Media:   "https://example.com/updated_again_image.jpg",
    }

    res, err := repo.UpdateTwit(context.Background(), &newReq)
	assert.NoError(t, err)
	assert.Equal(t, newReq.Content, res.Content)
	assert.Equal(t, newReq.Media, res.Media)
}

func TestDeleteTwit(t *testing.T) {
	mDB, err := ConnectMongo()
    assert.NoError(t, err)

    repo := NewTwitRepository(mDB)
    req := models.Twit{
        ID:      "8677fc69-2c73-4940-a7d5-26a04aa32608",
        UserID:  "user",
        Content: "Hello, World!",
        Media:   "https://example.com/image.jpg",
    }

    _, err = repo.CreateTwit(context.Background(), &req)
    assert.NoError(t, err)

    err = repo.DeleteTwit(context.Background(), req.ID)
    assert.NoError(t, err)
}

func TestGetTwits(t *testing.T) {
	mDB, err := ConnectMongo()
    assert.NoError(t, err)

    repo := NewTwitRepository(mDB)
    req := models.Twit{
        UserID:  "user",
        Content: "Hello, World!",
        Media:   "https://example.com/image.jpg",
    }

    _, err = repo.CreateTwit(context.Background(), &req)
    assert.NoError(t, err)

    res, err := repo.GetTwit(context.Background(), req.UserID)
    assert.NoError(t, err)
    assert.NotEmpty(t, res)
}

func TestGetFollowerTwits(t *testing.T) {
	mDB, err := ConnectMongo()
    assert.NoError(t, err)

    repo := NewTwitRepository(mDB)
    followerIDs := []string{"follower1", "follower2"}

    res, err := repo.GetFollowerTwit(context.Background(), followerIDs)
    assert.NoError(t, err)
    assert.NotEmpty(t, res)
}
