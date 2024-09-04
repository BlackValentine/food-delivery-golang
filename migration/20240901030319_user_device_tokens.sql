-- +goose Up
-- +goose StatementBegin
-- Create the table
CREATE TABLE IF NOT EXISTS user_device_tokens (
  id SERIAL PRIMARY KEY,
  user_id INT,
  is_production BOOLEAN DEFAULT FALSE,
  os TEXT CHECK (os IN ('ios', 'android', 'web')) DEFAULT 'ios',
  token VARCHAR(255),
  status SMALLINT NOT NULL DEFAULT 1,
  created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- Add indexes on user_id and os
CREATE INDEX idx_user_device_tokens_user_id ON user_device_tokens (user_id);
CREATE INDEX idx_os ON user_device_tokens (os);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS user_device_tokens;
-- +goose StatementEnd
