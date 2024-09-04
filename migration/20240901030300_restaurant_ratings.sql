-- +goose Up
-- +goose StatementBegin
-- Create the table
CREATE TABLE IF NOT EXISTS restaurant_ratings (
  id SERIAL PRIMARY KEY,
  user_id INT NOT NULL,
  restaurant_id INT NOT NULL,
  point FLOAT NOT NULL DEFAULT 0,
  comment TEXT,
  status INT NOT NULL DEFAULT 1,
  created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- Add indexes on user_id and restaurant_id
CREATE INDEX idx_restaurant_ratings_user_id ON restaurant_ratings (user_id);
CREATE INDEX idx_restaurant_ratings_restaurant_id ON restaurant_ratings (restaurant_id);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS restaurant_ratings;
-- +goose StatementEnd
