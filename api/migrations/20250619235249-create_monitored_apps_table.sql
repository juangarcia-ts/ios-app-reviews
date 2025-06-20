-- +migrate Up
CREATE TABLE IF NOT EXISTS "monitored_apps" (
  app_id text NOT NULL PRIMARY KEY,
  app_name text NOT NULL,
  logo_url text NOT NULL,
  nickname text,
  last_synced_at timestamp,
  created_at timestamp NOT NULL DEFAULT now(),
  updated_at timestamp NOT NULL DEFAULT now()
);

-- +migrate Down
DROP TABLE IF EXISTS "monitored_apps";