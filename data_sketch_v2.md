# Data Model Additions

## Donors
CREATE TABLE theaters_donors_bridge (
    theater_id   INTEGER    NOT NULL    REFERENCES theaters,
    donor_id     INTEGER    NOT NULL    REFERENCES donors,
    PRIMARY KEY ( theater_id, donor_id )
);

CREATE TABLE donors (
    donor_id INTEGER   NOT NULL   PRIMARY KEY,
    name     TEXT,
    donation INTEGER
);

## Staff

CREATE TABLE staff (
    staff_id    INTEGER  NOT NULL    PRIMARY KEY, 
    name        TEXT,
    position    TEXT,
    theater_id  INTEGER  NOT NULL    REFERENCES theaters 
 );

## Advertisers and advertisement 

CREATE TABLE advertisers (
    advertiser_id  INTEGER  NOT NULL    PRIMARY KEY, 
    name           TEXT,
    theater_id     INTEGER  NOT NULL    REFERENCES theaters 
 );

CREATE TABLE advertisements (
    advertisment_id  INTEGER  NOT NULL    PRIMARY KEY, 
    name             TEXT,
    content          BLOB,
    advertiser_id    INTEGER  NOT NULL    REFERENCES advertisers 
 );
