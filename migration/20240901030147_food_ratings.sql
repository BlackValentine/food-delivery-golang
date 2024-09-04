-- +goose Up
-- +goose StatementBegin
-- Create the table
CREATE TABLE IF NOT EXISTS food_ratings (
  id SERIAL PRIMARY KEY,
  user_id INT NOT NULL,
  food_id INT NOT NULL,
  point FLOAT DEFAULT 0,
  comment TEXT,
  status INT NOT NULL DEFAULT 1,
  created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- Add an index on food_id
CREATE INDEX idx_food_ratings_food_id ON food_ratings (food_id);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS food_ratings;
-- +goose StatementEnd
