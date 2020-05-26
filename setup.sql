CREATE TABLE theaters (
    theater_id  INTEGER   NOT NULL   PRIMARY KEY,
    name        TEXT,
    address     TEXT,
    description TEXT
);

CREATE TABLE shows (
    show_id     INTEGER NOT NULL PRIMARY KEY,
    name        TEXT,
    company     TEXT,
    description string
);

CREATE TABLE theaters_shows_bridge (
    theater_id   INTEGER    NOT NULL    REFERENCES theaters,
    show_id      INTEGER    NOT NULL    REFERENCES shows,
    PRIMARY KEY ( theater_id, show_id )
);
