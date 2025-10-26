-- +goose Up

-- Create schema if it doesn't exist, owned by gator_migrator
CREATE SCHEMA IF NOT EXISTS app AUTHORIZATION gator_migrator;

-- Grants for roles
GRANT USAGE, CREATE ON SCHEMA app TO gator_migrator;
GRANT USAGE ON SCHEMA app TO gator_app;

-- Default privileges for future tables and sequences
ALTER DEFAULT PRIVILEGES FOR ROLE gator_migrator IN SCHEMA app
  GRANT SELECT, INSERT, UPDATE, DELETE ON TABLES TO gator_app;

ALTER DEFAULT PRIVILEGES FOR ROLE gator_migrator IN SCHEMA app
  GRANT USAGE, SELECT, UPDATE ON SEQUENCES TO gator_app;

-- Create goose_db_version table for goose
CREATE TABLE IF NOT EXISTS app.goose_db_version (
    id serial PRIMARY KEY,
    version_id bigint NOT NULL,
    is_applied boolean NOT NULL,
    tstamp timestamp NOT NULL DEFAULT now()
);

-- +goose Down
DROP TABLE IF EXISTS app.goose_db_version;
-- Optional: drop schema if you want a full rollback
-- DROP SCHEMA IF EXISTS app CASCADE;
