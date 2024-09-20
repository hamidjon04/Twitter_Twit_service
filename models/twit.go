package models

type Twit struct {
	ID        string `bson:"_id"`
	UserID    string `bson:"user_id"`
	Content   string `bson:"content"`
	Media     string `bson:"media"`
	CreatedAt string `bson:"created_at"`
	UpdatedAt string `bson:"updated_at"`
	DeletedAt string `bson:"deleted_at"`
}

type Like struct {
	Twit_id    string `bson:"twit_id"`
	Clicker_id string `bson:"clicker_id"`
}

type UpdateTwit struct{
	Id        string	`bson:"_id"`
	UserId    string	`bson:"user_id"`
	Content   string	`bson:"content"`
	Media     string	`bson:"media"`
	CreatedAt string	`bson:"created_at"`
	UpdatedAt string	`bson:"updated_at"`
}