BEGIN;

ALTER TABLE releases DROP COLUMN is_sent;

COMMIT;