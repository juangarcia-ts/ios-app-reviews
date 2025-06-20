-- +migrate Up
CREATE TABLE IF NOT EXISTS "monitored_apps" (
  app_id text NOT NULL PRIMARY KEY,
  nickname text NOT NULL,
  last_synced_at timestamp,
  created_at timestamp NOT NULL DEFAULT now(),
  updated_at timestamp NOT NULL DEFAULT now()
);

-- +migrate Down
DROP TABLE IF EXISTS "monitored_apps";