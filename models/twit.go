package models

type Twit struct {
	ID        string `json:"id"`
	UserID    string `json:"user_id"`
	Content   string `json:"content"`
	Media     string `json:"media"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
	DeletedAt string `json:"deleted_at"`
}

type Like struct {
	Twit_id    string `json:"twit_id"`
	Clicker_id string `json:"clicker_id"`
	CreatedAt  string `json:"created_at"`
}
