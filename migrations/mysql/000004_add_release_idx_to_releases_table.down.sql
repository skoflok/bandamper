BEGIN;

ALTER TABLE releases DROP INDEX release_idx;

COMMIT;