# Data Model Sketch
```

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
    order_appearance INTEGER,
    performer_id     INTEGER  NOT NULL    REFERENCES performers,
    show_id          INTEGER  NOT NULL    REFERENCES shows 
);

CREATE TABLE performers_shows_bridge (
    performer_id INTEGER    NOT NULL    REFERENCES performers,
    show_id      INTEGER    NOT NULL    REFERENCES shows,
    PRIMARY KEY ( performer_id, show_id )
);

CREATE TABLE users (
    user_id INTEGER NOT NULL PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    email VARCHAR(255) NOT NULL UNIQUE,
    hashed_password CHAR(60) NOT NULL,
    created DATETIME NOT NULL,
    active BOOLEAN NOT NULL DEFAULT TRUE
)


```
