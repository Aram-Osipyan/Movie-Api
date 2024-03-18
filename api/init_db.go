package main

const InitSqlSchemaMigration string = `
CREATE TABLE IF NOT EXISTS artists (
    id SERIAL PRIMARY KEY,
    name VARCHAR(100) NOT NULL,
    sex CHAR(1) NOT NULL CHECK (sex IN ('M', 'F')),
    birth_date DATE NOT NULL
);

CREATE TABLE IF NOT EXISTS movies (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    description TEXT,
    release_date DATE,
    rating NUMERIC(10,2) CHECK (rating BETWEEN 0.0 AND 5.0)
);

CREATE TABLE IF NOT EXISTS movie_artists (
    id SERIAL PRIMARY KEY,
    movie_id INTEGER REFERENCES movie(id),
    artist_id INTEGER REFERENCES artist(id)
);

CREATE TABLE IF NOT EXISTS users (
    id SERIAL PRIMARY KEY,
    username VARCHAR(50) NOT NULL UNIQUE,
    password VARCHAR(255) NOT NULL,
    role VARCHAR(20) NOT NULL DEFAULT 'user'
);
`
