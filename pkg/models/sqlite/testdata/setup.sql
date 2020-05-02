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

INSERT INTO theaters (name)
VALUES ("Des Moines Civic Center");


INSERT INTO users (name, email, hashed_password, created) VALUES (
    "nick", 
    "neherrig@gmail.com", 
    "aoskdjf;lkasdjflkasjf",
    "2018-12-23 17:25:22"
);
