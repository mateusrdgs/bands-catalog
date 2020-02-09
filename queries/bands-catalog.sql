CREATE DATABASE bands_catalog;

USE bands_catalog;

CREATE TABLE bands (
	uuid VARCHAR(36) PRIMARY KEY,
  name VARCHAR(100) NOT NULL,
  year_of_foundation INT,
  biography VARCHAR(255),
  country VARCHAR(50),
  genre VARCHAR(25) NOT NULL
);

CREATE TABLE artists (
	uuid VARCHAR(36) PRIMARY KEY,
  name VARCHAR(100) NOT NULL,
  age INT,
  country VARCHAR(60),
  genre VARCHAR(15),
  biography VARCHAR(200),
  instrument VARCHAR(100)
);

CREATE TABLE bands_artists (
	uuid VARCHAR(36) PRIMARY KEY,
  band_uuid VARCHAR(36),
  artist_uuid VARCHAR(36),
  FOREIGN KEY (band_uuid) REFERENCES bands(uuid),
  FOREIGN KEY (artist_uuid) REFERENCES artists(uuid)
);

CREATE TABLE albums (
	uuid VARCHAR(36) PRIMARY KEY,
  name VARCHAR(100) NOT NULL,
  type VARCHAR(25) NOT NULL,
  release_date DATE,
  label VARCHAR(50)
);

CREATE TABLE bands_albums (
	uuid VARCHAR(36) PRIMARY KEY,
  band_uuid VARCHAR(36),
  album_uuid VARCHAR(36),
  FOREIGN KEY (band_uuid) REFERENCES bands(uuid),
  FOREIGN KEY (album_uuid) REFERENCES albums(uuid)
);

CREATE TABLE songs (
	uuid VARCHAR(36) PRIMARY KEY,
  name VARCHAR(255) NOT NULL,
  number INT NOT NULL,
  duration INT NOT NULL,
  lyrics TEXT
);

CREATE TABLE albums_songs (
	uuid VARCHAR(36) PRIMARY KEY,
  album_uuid VARCHAR(36),
  song_uuid VARCHAR(36),
  FOREIGN KEY (album_uuid) REFERENCES albums(uuid),
  FOREIGN KEY (song_uuid) REFERENCES songs(uuid)
);

CREATE PROCEDURE insert_band (
	IN band_name VARCHAR(100),
  IN year_of_foundation INT,
  IN biography VARCHAR(255),
  IN country VARCHAR(50),
  IN genre VARCHAR(25)
)
INSERT INTO bands
VALUES (uuid(), band_name, year_of_foundation, biography, country, genre);

CALL insert_band (
	'Shadow of Intent',
  2013,
  'The Shadow of Intent is the name of a spacecraft in the Halo game series.',
  'United States',
  'Deathcore'
);

DELIMITER |
  CREATE PROCEDURE insert_album (
    IN name VARCHAR(100),
    IN type VARCHAR(25),
    IN release_date DATE,
    IN label VARCHAR(50),
    IN band_uuid VARCHAR(36)
  )
  BEGIN
    DECLARE album_uuid VARCHAR(36) DEFAULT uuid();
    DECLARE band_album VARCHAR(36) DEFAULT uuid();

    INSERT INTO albumsinsert_albuminsert_album
    VALUES (album_uuid, name, type, release_date, label);

    INSERT INTO bands_albums
    VALUES (band_album, band_uuid, album_uuid);
  END
|

CALL insert_album (
	'Melancholy',
  'Full-length',
  '2019-08-16',
  'Independent',
  'a69f3558-4960-11ea-bef1-8d5d28b20504'
);

SELECT b.name as band_name, a.name as album_name
FROM bands as b
INNER JOIN bands_albums as ba
ON b.uuid LIKE ba.band_uuid
INNER JOIN albums as a
ON ba.album_uuid LIKE a.uuid;

DELIMITER |
  CREATE PROCEDURE insert_song (
    IN name VARCHAR(255),
    IN number INT,
    IN duration INT,
    IN lyrics TEXT,
    IN album_uuid VARCHAR(36)
  )
  BEGIN
    DECLARE song_uuid VARCHAR(36) DEFAULT uuid();
    DECLARE album_song_uuid VARCHAR(36) DEFAULT uuid();

    INSERT INTO songs
    VALUES (song_uuid, name, number, duration, lyrics);

    INSERT INTO albums_songs
    VALUES (album_song_uuid, album_uuid, song_uuid);
  END
|

CALL insert_song (
	'Melancholy',
  1,
  315,
  'Deep inside the madness designed, you rot\nHear this foretelling of crepuscule doom most feral\nThe desolation of a human, an incarnate disgrace\nObliterate proceeds their fatal function; disease\nDevastation is the antidote\nChanting for the anathema of grand dolor\nDishonoring fathers through shedding of blood\nEnviron the influence rising through all\nInto the mouth of the maelstrom\nForetelling of crepuscule doom most feral\nEnviron the influence rising through all\nAll is lost\nFalling into oblivion\n
  It shall envelope all\nEngulfed by fire and cast to the flame\n
  Pure melancholy\nThey begin collapsing downward through the obscure godless chasm, plummeting an infinite depth inside the mind\nDeep inside the madness designed looking outward awaiting their death\nBehold!\nKnowledge of life and power\nThat was not what they found\nThey did not choose to survive\nA reason couldn’t be found\nIt took their lives',
  'b31bd6b4-4ad3-11ea-bef1-8d5d28b20504'
);

DELIMITER |
  CREATE PROCEDURE insert_artist (
    IN name VARCHAR(100),
    IN age INT,
    IN country VARCHAR(60),
    IN genre VARCHAR(15),
    IN biography VARCHAR(200),
    IN instrument VARCHAR(100),
    IN band_uuid VARCHAR(36)
  )
  BEGIN
    DECLARE artist_uuid VARCHAR(36) DEFAULT uuid();
    DECLARE band_artist_uuid VARCHAR(36) DEFAULT uuid();

    INSERT INTO artists
    VALUES (artist_uuid, name, age, country, genre, biography, instrument);

    INSERT INTO bands_artists
    VALUES (band_artist_uuid , band_uuid, artist_uuid);
  END
|

CALL insert_artist (
	'Ben Duerr',
  '25',
  'United States',
  'Male',
  '',
  'Vocals',
  'a69f3558-4960-11ea-bef1-8d5d28b20504'
);