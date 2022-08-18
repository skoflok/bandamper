BEGIN;

ALTER TABLE releases DROP INDEX release release_idx;

COMMIT;