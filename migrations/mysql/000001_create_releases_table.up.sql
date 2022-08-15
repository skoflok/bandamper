BEGIN;

CREATE TABLE IF NOT EXISTS releases (
    id BIGINT UNSIGNED NOT NULL AUTO_INCREMENT PRIMARY KEY,
	type VARCHAR(255) DEFAULT "",
    release_id BIGINT UNSIGNED NOT NULL,
    band_id BIGINT UNSIGNED NOT NULL, INDEX (band_id),
    is_preorder SMALLINT default 0,
    publish_date DATETIME NOT NULL,
    genre VARCHAR(255) DEFAULT "",
    album TEXT DEFAULT "",
    artist TEXT DEFAULT "",
    featured_track TEXT DEFAULT ""
);

COMMIT;