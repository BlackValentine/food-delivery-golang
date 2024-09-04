-- +goose Up
-- +goose StatementBegin
-- Create the table
CREATE TABLE IF NOT EXISTS order_trackings (
  id SERIAL PRIMARY KEY,
  order_id INT NOT NULL,
  state TEXT CHECK (state IN ('waiting_for_shipper', 'preparing', 'on_the_way', 'delivered', 'cancel')) NOT NULL,
  status INT NOT NULL DEFAULT 1,
  created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- Add an index on order_id
CREATE INDEX idx_order_trackings_order_id ON order_trackings (order_id);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS order_trackings;
-- +goose StatementEnd
