# Data Model
This is a draft sketch of the database model.


## MUSICAL NUMBERS DATA (One-to-many data relationship)
Each show or event will have a number of songs/musical numbers.
This relationship is a one-to-many data relationship where a show_id is a foriegn key for a song.

'''

CREATE TABLE shows (
    show_id INTEGER  NOT NULL    PRIMARY KEY,
    name    TEXT
    company TEXT
    ...
 );


CREATE TABLE songs (
    song_id    INTEGER  NOT NULL    PRIMARY KEY, 
    name       TEXT,
    number     INTEGER,
    act_number INTEGER,
    show_id    INTEGER  NOT NULL    REFERENCES shows
    ...
 );

'''

## THE CAST DATA MODEL (Many-to-many data relationship)
Each Performer, _could_  be apart of multiple shows. 
Along with that, each show, will have multiple performers. 
This data relationship is a many-to-many design relationship. 
In order for this relationship to be represented a link/bridge table is needed.
This will link the shows table to the performers table with foreign keys.
this uses the show_id from table above to bridge the performers.

'''

CREATE TABLE performers (
    performer_id INTEGER NOT NULL   PRIMARY KEY,
    name         TEXT,
    bio          TEXT,
    ...
);

CREATE TABLE performer_show_bridge (
    performer_id INTEGER    NOT NULL    REFERENCES performers,
    show_id      INTEGER    NOT NULL    REFERENCES shows,
    PRIMARY KEY ( performer_id, show_id )
);

'''

## THE THEATER DATA MODEL (One-to-many data relationships)
The theater data model has a few one-to-many relationships.
One theater has many staff members, and many Advertisers.

'''

CREATE TABLE theaters (
    theater_id INTEGER   NOT NULL   PRIMARY KEY,
    theater_name,
    ...
);

'''


## THE THEATER DATA MODEL (Many-to-many data relationships)
The theater data model has a few many-to-many relationships.
A theater has multiple shows, and a show may performe at  multiple theaters
A theater has multiple donors, and a donor may donate to multiple theaters 

## THE ADDVERTISER MODEL (one-to-many data relationships)
The Addvertiser data model has a one-to-many relationships.
An Advertiser will have multiple adds.

