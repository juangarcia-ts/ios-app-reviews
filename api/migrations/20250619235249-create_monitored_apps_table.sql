-- +migrate Up
CREATE TABLE IF NOT EXISTS "monitored_apps" (
  app_id text NOT NULL PRIMARY KEY,
  polling_interval_in_minutes int NOT NULL DEFAULT 60,
  created_at timestamp NOT NULL DEFAULT now(),
  updated_at timestamp NOT NULL DEFAULT now()
);

-- +migrate Down
DROP TABLE IF EXISTS "monitored_apps";