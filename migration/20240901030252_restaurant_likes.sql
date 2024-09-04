-- +goose Up
-- +goose StatementBegin
-- Create the table
CREATE TABLE IF NOT EXISTS restaurant_likes (
  restaurant_id INT NOT NULL,
  user_id INT NOT NULL,
  created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (restaurant_id, user_id)
);

-- Add an index on user_id
CREATE INDEX idx_restaurant_likes_user_id ON restaurant_likes (user_id);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS restaurant_likes;
-- +goose StatementEnd
