-- +migrate Up
CREATE TABLE IF NOT EXISTS "app_reviews" (
  review_id text NOT NULL PRIMARY KEY,
  app_id text NOT NULL,
  title text NOT NULL,
  content text NOT NULL,
  author text NOT NULL,
  rating int NOT NULL,
  submitted_at timestamp NOT NULL,
  created_at timestamp NOT NULL DEFAULT now(),
  updated_at timestamp NOT NULL DEFAULT now()
);

-- +migrate Down
DROP TABLE IF EXISTS "app_reviews";