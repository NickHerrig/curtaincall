# Data Model Sketch

```
CREATE TABLE theaters (
    theater_id INTEGER   NOT NULL   PRIMARY KEY,
    name       TEXT
);

CREATE TABLE theaters_shows_bridge (
    theater_id   INTEGER    NOT NULL    REFERENCES theaters,
    show_id      INTEGER    NOT NULL    REFERENCES shows,
    PRIMARY KEY ( theater_id, show_id )
);


CREATE TABLE shows (
    show_id INTEGER  NOT NULL    PRIMARY KEY,
    name    TEXT
    company TEXT
 );

CREATE TABLE songs (
    song_id    INTEGER  NOT NULL    PRIMARY KEY, 
    name       TEXT,
    number     INTEGER,
    act_number INTEGER,
    show_id    INTEGER  NOT NULL    REFERENCES shows
 );

CREATE TABLE performers (
    performer_id INTEGER NOT NULL   PRIMARY KEY,
    name         TEXT,
    bio          TEXT
);

CREATE TABLE roles (
    role_id          INTEGER NOT NULL   PRIMARY KEY,
    name             TEXT,
    order            INTEGER,
    performer_id     INTEGER  NOT NULL    REFERENCES performers 
    show_id          INTEGER  NOT NULL    REFERENCES shows 
);

CREATE TABLE performer_show_bridge (
    performer_id INTEGER    NOT NULL    REFERENCES performers,
    show_id      INTEGER    NOT NULL    REFERENCES shows,
    PRIMARY KEY ( performer_id, show_id )
);


```
