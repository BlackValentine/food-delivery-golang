-- +goose Up
-- +goose StatementBegin
-- Create the table
CREATE TABLE IF NOT EXISTS food_likes (
  user_id INT NOT NULL,
  food_id INT NOT NULL,
  created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (user_id, food_id)
);

-- Add an index on food_id
CREATE INDEX idx_food_likes_food_id ON food_likes (food_id);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS food_likes;
-- +goose StatementEnd
