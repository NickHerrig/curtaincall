CREATE TABLE theaters (
    theater_id INTEGER   NOT NULL   PRIMARY KEY,
    name       TEXT
);

CREATE TABLE users (
    user_id INTEGER NOT NULL PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    email VARCHAR(255) NOT NULL UNIQUE,
    hashed_password CHAR(60) NOT NULL,
    created DATETIME NOT NULL,
    active BOOLEAN NOT NULL DEFAULT TRUE
);

CREATE TABLE shows (
    show_id INTEGER  NOT NULL    PRIMARY KEY,
    name    TEXT,
    company TEXT
 );

CREATE TABLE theaters_shows_bridge (
    theater_id   INTEGER    NOT NULL    REFERENCES theaters,
    show_id      INTEGER    NOT NULL    REFERENCES shows,
    PRIMARY KEY ( theater_id, show_id )
);

INSERT INTO theaters (name)
VALUES ("Des Moines Civic Center");

INSERT INTO theaters (name)
VALUES ("The Fabulous Fox");

INSERT INTO shows (name, company)
VALUES ("Hamilton", "Company A");

INSERT INTO shows (name, company)
VALUES ("Dear Evan Hansen", "Company B");

INSERT INTO theaters_shows_bridge (theater_id, show_id)
VALUES (1,1);

INSERT INTO theaters_shows_bridge (theater_id, show_id)
VALUES (2,2);

INSERT INTO users (name, email, hashed_password, created) VALUES (
    "nick", 
    "neherrig@gmail.com", 
    "aoskdjf;lkasdjflkasjf",
    "2018-12-23 17:25:22"
);
