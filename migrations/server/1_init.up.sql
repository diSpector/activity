CREATE TABLE IF NOT EXISTS activities
(
    id          INTEGER PRIMARY KEY,
    description TEXT NOT NULL,
    persons     INTEGER NOT NULL,
    created     DATETIME DEFAULT CURRENT_TIMESTAMP
);