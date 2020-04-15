# Test Data for V1.0


INSERT INTO theaters (name)
VALUES ("Des Moines Civic Center");

INSERT INTO shows (name, company)
VALUES ("Hamilton", "The Public Theater");

INSERT INTO shows (name, company)
VALUES ("Dear Evan Hansen", "The Public Theater");


INSERT INTO theaters_shows_bridge (theater_id, show_id) 
VALUES (1,1);

INSERT INTO theaters_shows_bridge (theater_id, show_id) 
VALUES (1,2);

INSERT INTO songs (name, number, act_number, show_id)
VALUES ("Alexander Hamilton", 1, 1, 1);

INSERT INTO songs (name, number, act_number, show_id)
VALUES ("The Story of Tonight", 3, 1, 1);

INSERT INTO songs (name, number, act_number, show_id)
VALUES ("What'd I Miss", 18, 2, 1);

INSERT INTO songs (name, number, act_number, show_id)
VALUES ("Waving Through a Window", 2, 1, 2);

INSERT INTO songs (name, number, act_number, show_id)
VALUES ("Requiem", 5, 1, 2);

INSERT INTO songs (name, number, act_number, show_id)
VALUES ("Only Us", 11, 2, 2);

INSERT INTO performers (name, bio)
VALUES ("Joseph Morales", "Joseph Morales' Bio! He's great..");

INSERT INTO performers (name, bio)
VALUES ("Ben Levi Ross", "Ben Levi Ross' Bio! He's also great..");

INSERT INTO performers_shows_bridge (performer_id, show_id) 
VALUES (1,1);

INSERT INTO performers_shows_bridge (perforomer_id, show_id) 
VALUES (1,2);

INSERT INTO roles (name, order_appearance, performer_id, show_id) 
VALUES ("Alexander Hamilton", 1, 1, 1);

INSERT INTO roles (name, order_appearance , performer_id, show_id) 
VALUES ("Evan Hansen", 7, 2, 2);
