-- +goose Up
-- +goose StatementBegin
-- Create the table
CREATE TABLE IF NOT EXISTS user_addresses (
  id SERIAL PRIMARY KEY,
  user_id INT NOT NULL,
  city_id INT NOT NULL,
  title VARCHAR(100),
  icon JSONB,
  addr VARCHAR(255) NOT NULL,
  lat DOUBLE PRECISION,
  lng DOUBLE PRECISION,
  status INT NOT NULL DEFAULT 1,
  created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- Add indexes on user_id and city_id
CREATE INDEX idx_user_addresses_user_id ON user_addresses (user_id);
CREATE INDEX idx_user_addresses_city_id ON user_addresses (city_id);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS user_addresses;
-- +goose StatementEnd
