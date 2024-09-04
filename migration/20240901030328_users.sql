-- +goose Up
-- +goose StatementBegin
-- Create the table
CREATE TABLE IF NOT EXISTS users (
  id SERIAL PRIMARY KEY,
  email VARCHAR(50) NOT NULL UNIQUE,
  fb_id VARCHAR(50),
  gg_id VARCHAR(50),
  password VARCHAR(50) NOT NULL,
  salt VARCHAR(50),
  last_name VARCHAR(50) NOT NULL,
  first_name VARCHAR(50) NOT NULL,
  phone VARCHAR(20),
  role TEXT CHECK (role IN ('user', 'admin', 'shipper')) NOT NULL DEFAULT 'user',
  avatar JSONB,
  status INT NOT NULL DEFAULT 1,
  created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS users;
-- +goose StatementEnd
