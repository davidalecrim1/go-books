-- schema.sql
CREATE TABLE IF NOT EXISTS books (
    ID     INTEGER PRIMARY KEY,
    Title  TEXT NOT NULL,
    Author TEXT NOT NULL,
    Genre  TEXT NOT NULL,
    Year   INTEGER NOT NULL
);