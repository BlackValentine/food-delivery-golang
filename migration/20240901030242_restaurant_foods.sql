-- +goose Up
-- +goose StatementBegin
-- Create the table
CREATE TABLE IF NOT EXISTS restaurant_foods (
  restaurant_id INT NOT NULL,
  food_id INT NOT NULL,
  status INT NOT NULL DEFAULT 1,
  created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (restaurant_id, food_id)
);

-- Add an index on food_id
CREATE INDEX idx_restaurant_foods_food_id ON restaurant_foods (food_id);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS restaurant_foods;
-- +goose StatementEnd
