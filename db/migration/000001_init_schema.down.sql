DROP TABLE IF EXISTS entries;
DROP TABLE IF EXISTS accounts;
DROP TABLE IF EXISTS transfers;
DELETE FROM schema_migrations WHERE version = -1;