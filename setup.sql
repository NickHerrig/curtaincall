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

INSERT INTO theaters (name, address, description)
VALUES (
    "Des Moines Civic Center", 
    "221 Walnut St, Des Moines, IA 50309",
    "The Des Moines Civic Center is a 2,744-seat performing arts\n
     center belonging to Des Moines Performing Arts located in Des Moines, Iowa.\n
     It has been Iowa's largest theater since it opened on June 10, 1979,\n
     and is used for concerts, Broadway shows, ballets, and other special events.");

INSERT INTO theaters (name, address, description)
VALUES (
    "The Fox Theatre", 
    "527 N Grand Blvd, St. Louis, MO 63103",
    "The Fox Theatre, a former movie palace, is a performing arts center located at\n
     527 N. Grand Blvd. in St. Louis, Missouri. Also known as The Fabulous Fox, it is\n
     situated in the arts district of the Grand Center area in Midtown St. Louis, one\n
     block north of Saint Louis University.");
