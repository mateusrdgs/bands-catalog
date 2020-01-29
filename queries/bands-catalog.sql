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