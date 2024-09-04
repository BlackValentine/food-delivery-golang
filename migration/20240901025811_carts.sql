-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS carts (
  user_id INT NOT NULL,
  food_id INT NOT NULL,
  quantity INT NOT NULL,
  status INT NOT NULL DEFAULT 1,
  created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (user_id, food_id)
);
CREATE INDEX idx_food_id ON carts (food_id);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS carts
-- +goose StatementEnd
