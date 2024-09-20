CREATE TABLE likes (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(), -- Yunik ID
    twit_id UUID NOT NULL, -- Twitga bog'langan ID
    clicker_id UUID NOT NULL, -- Like qo'shgan foydalanuvchi ID
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP, -- Like qo'shilgan vaqt
    CONSTRAINT unique_like UNIQUE (twit_id, clicker_id)
);

CREATE INDEX idx_likes_twit_id ON likes(twit_id);
CREATE INDEX idx_likes_user_id ON likes(user_id);
