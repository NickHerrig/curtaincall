# Data Model Sketch

```
CREATE TABLE theaters (
    theater_id INTEGER   NOT NULL   PRIMARY KEY,
    name       TEXT,
    ...
);

CREATE TABLE theaters_donors_bridge (
    theater_id   INTEGER    NOT NULL    REFERENCES theaters,
    donor_id     INTEGER    NOT NULL    REFERENCES donors,
    PRIMARY KEY ( theater_id, donor_id )
);

CREATE TABLE donors (
    donor_id INTEGER   NOT NULL   PRIMARY KEY,
    name     TEXT,
    donation INTEGER,
    ...
);

CREATE TABLE theaters_shows_bridge (
    theater_id   INTEGER    NOT NULL    REFERENCES theaters,
    show_id      INTEGER    NOT NULL    REFERENCES shows,
    PRIMARY KEY ( theater_id, show_id )
);

CREATE TABLE staff (
    staff_id    INTEGER  NOT NULL    PRIMARY KEY, 
    name        TEXT,
    position    TEXT,
    theater_id  INTEGER  NOT NULL    REFERENCES theaters 
    ...
 );

CREATE TABLE advertisers (
    advertiser_id  INTEGER  NOT NULL    PRIMARY KEY, 
    name           TEXT,
    theater_id     INTEGER  NOT NULL    REFERENCES theaters 
    ...
 );

CREATE TABLE advertisements (
    advertisment_id  INTEGER  NOT NULL    PRIMARY KEY, 
    name             TEXT,
    content          BLOB,
    advertiser_id    INTEGER  NOT NULL    REFERENCES advertisers 
    ...
 );

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


CREATE TABLE theaters (
    theater_id INTEGER   NOT NULL   PRIMARY KEY,
    theater_name,
    ...
);

```

# Data Brain Dump V1.0

## THE THEATER DATA Relationships

### One-to-many data relationships
One theater has many staff members.
One theater has many advertisers.

### Many-to-many data relationships
A theater has multiple shows, and a show may perform at multiple theaters
A theater has multiple donors, and a donor may donate to multiple theaters 

## THE ADDVERTISER DATA RELATIONSHIP
### One-to-many data relationship
An Advertiser will have multiple addvertisements.

##  SHOWS DATA RELATIONSHIP
### One-to-many data relationships
Each show will have a number of songs/musical numbers.

### Many-to-many data relationship
A show has multiple performers, and a performer can be in multiple shows

