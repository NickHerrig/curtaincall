CREATE TABLE theaters (
    theater_id INTEGER   NOT NULL   PRIMARY KEY,
    name       TEXT
);

CREATE TABLE shows (
    show_id INTEGER  NOT NULL    PRIMARY KEY,
    name    TEXT,
    company TEXT
 );

CREATE TABLE songs (
    song_id         INTEGER  NOT NULL    PRIMARY KEY, 
    song_name       TEXT,
    song_number     INTEGER,
    song_act_number INTEGER,
    show_id         INTEGER  NOT NULL    REFERENCES shows
 );

CREATE TABLE theaters_shows_bridge (
    theater_id   INTEGER    NOT NULL    REFERENCES theaters,
    show_id      INTEGER    NOT NULL    REFERENCES shows,
    PRIMARY KEY ( theater_id, show_id )
);

/* Setup Theaters */
INSERT INTO theaters (name)
VALUES ("Des Moines Civic Center");

INSERT INTO theaters (name)
VALUES ("The Fabulous Fox");

/* Setup Shows */
INSERT INTO shows (name, company)
VALUES ("Hamilton", "Company A");

INSERT INTO shows (name, company)
VALUES ("Dear Evan Hansen", "Company B");

INSERT INTO shows (name, company)
VALUES ("Mean Girls", "Company C");

/* Setup Mean Girls Songs */
INSERT INTO songs (song_name, song_number, song_act_number, show_id)
VALUES ("A Cautionary Tale", 1, 1, 3);
INSERT INTO songs (song_name, song_number, song_act_number, show_id)
VALUES ("It Roars", 2, 1, 3);
INSERT INTO songs (song_name, song_number, song_act_number, show_id)
VALUES ("Where Do You Belong?", 3, 1, 3);
INSERT INTO songs (song_name, song_number, song_act_number, show_id)
VALUES ("Meet the Plastics", 4, 1, 3);
INSERT INTO songs (song_name, song_number, song_act_number, show_id)
VALUES ("Stupid with Love", 5, 1, 3);
INSERT INTO songs (song_name, song_number, song_act_number, show_id)
VALUES ("Apex Predator", 6, 1, 3);
INSERT INTO songs (song_name, song_number, song_act_number, show_id)
VALUES ("What's Wrong with Me?", 7, 1, 3);
INSERT INTO songs (song_name, song_number, song_act_number, show_id)
VALUES ("Sexy", 8, 1, 3);
INSERT INTO songs (song_name, song_number, song_act_number, show_id)
VALUES ("Someone Gets Hurt", 9, 1, 3);
INSERT INTO songs (song_name, song_number, song_act_number, show_id)
VALUES ("Revenge Party", 10, 1, 3);
INSERT INTO songs (song_name, song_number, song_act_number, show_id)
VALUES ("Fearless", 11, 1, 3);
INSERT INTO songs (song_name, song_number, song_act_number, show_id)
VALUES ("Stop", 12, 2, 3);
INSERT INTO songs (song_name, song_number, song_act_number, show_id)
VALUES ("What's Wrong with Me? (Reprise)", 13, 2, 3);
INSERT INTO songs (song_name, song_number, song_act_number, show_id)
VALUES ("Whose House Is This?", 14, 2, 3);
INSERT INTO songs (song_name, song_number, song_act_number, show_id)
VALUES ("More Is Better", 15, 2, 3);
INSERT INTO songs (song_name, song_number, song_act_number, show_id)
VALUES ("World Burn", 16, 2, 3);
INSERT INTO songs (song_name, song_number, song_act_number, show_id)
VALUES ("I'd Rather Be Me", 17, 2, 3);
INSERT INTO songs (song_name, song_number, song_act_number, show_id)
VALUES ("Fearless (Reprise)", 18, 2, 3);
INSERT INTO songs (song_name, song_number, song_act_number, show_id)
VALUES ("Do This Thing", 19, 2, 3);
INSERT INTO songs (song_name, song_number, song_act_number, show_id)
VALUES ("I See Stars", 20, 2, 3);

/* Setup Show and Theater Many-to-many relationships */
INSERT INTO theaters_shows_bridge (theater_id, show_id)
VALUES (1,1);

INSERT INTO theaters_shows_bridge (theater_id, show_id)
VALUES (2,2);

