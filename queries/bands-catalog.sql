CREATE DATABASE bands_catalog;

USE bands_catalog;

CREATE TABLE bands (
    uuid VARCHAR(36),
    band_name VARCHAR(100),
    year_of_foundation INT,
    biography VARCHAR(255),
    country VARCHAR(50),
    genre VARCHAR(25)
);

INSERT INTO bands
VALUES (uuid(), 'Nine Inch Nails', 1988, 'Founded by Trent Reznor', 'United States', 'Industrial');

SELECT * FROM bands;

CREATE PROCEDURE save_band(
	IN band_name VARCHAR(100),
    IN year_of_foundation INT,
    IN biography VARCHAR(255),
    IN country VARCHAR(50),
    IN genre VARCHAR(25)
)
INSERT INTO bands
VALUES (uuid(), band_name, year_of_foundation, biography, country, genre);

CALL save_band (
	'Shadow of Intent',
    2013,
    'The Shadow of Intent is the name of a spacecraft in the Halo game series.',
    'United States',
    'Deathcore'
);

