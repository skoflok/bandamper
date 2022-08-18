BEGIN;

ALTER TABLE releases ADD CONSTRAINT UNIQUE INDEX release_idx (release_id);

COMMIT;