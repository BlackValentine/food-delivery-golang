-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS restaurants (
  id SERIAL PRIMARY KEY,
  owner_id INT DEFAULT NULL,
  name VARCHAR(255) NOT NULL,
  addr VARCHAR(255) NOT NULL,
  city_id INT DEFAULT NULL,
  lat DOUBLE PRECISION DEFAULT NULL,
  lng DOUBLE PRECISION DEFAULT NULL,
  cover JSON DEFAULT NULL,
  logo JSON DEFAULT NULL,
  shipping_fee_per_km DOUBLE PRECISION DEFAULT 0,
  status INT NOT NULL DEFAULT 1,
  liked_count INT NOT NULL DEFAULT 0,
  created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
-- Index on owner_id
CREATE INDEX idx_owner_id ON restaurants(owner_id);

-- Index on city_id
CREATE INDEX idx_city_id ON restaurants(city_id);

-- Index on status
CREATE INDEX idx_status ON restaurants(status);

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS restaurants
-- +goose StatementEnd
