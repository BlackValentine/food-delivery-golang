-- +goose Up
-- +goose StatementBegin
-- Create the table
CREATE TABLE IF NOT EXISTS foods (
  id SERIAL PRIMARY KEY,
  restaurant_id INT NOT NULL,
  category_id INT,
  name VARCHAR(255) NOT NULL,
  description TEXT,
  price FLOAT NOT NULL,
  images JSONB NOT NULL,
  status INT NOT NULL DEFAULT 1,
  created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- Add indexes on restaurant_id, category_id, and status
CREATE INDEX idx_restaurant_id ON foods (restaurant_id);
CREATE INDEX idx_category_id ON foods (category_id);
CREATE INDEX idx_foods_status ON foods (status);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS foods;
-- +goose StatementEnd
