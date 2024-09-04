-- +goose Up
-- +goose StatementBegin
-- Create the table
CREATE TABLE IF NOT EXISTS order_details (
  id SERIAL PRIMARY KEY,
  order_id INT NOT NULL,
  food_origin JSONB DEFAULT NULL,
  price FLOAT NOT NULL,
  quantity INT NOT NULL,
  discount FLOAT DEFAULT 0,
  status INT NOT NULL DEFAULT 1,
  created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- Add an index on order_id
CREATE INDEX idx_order_details_order_id ON order_details (order_id);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS order_details;
-- +goose StatementEnd
