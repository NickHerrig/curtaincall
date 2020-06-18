CREATE TABLE theaters (
    theater_id  INTEGER   NOT NULL   PRIMARY KEY,
    name        TEXT,
    address     TEXT,
    description TEXT
);

INSERT INTO theaters (name, address, description)
VALUES (
    "Des Moines Civic Center", 
    "221 Walnut St, Des Moines, IA 50309",
    "The Des Moines Civic Center is a 2,744-seat performing arts center belonging to Des Moines Performing Arts located in Des Moines, Iowa. It has been Iowa's largest theater since it opened on June 10, 1979, and is used for concerts, Broadway shows, ballets, and other special events.");

INSERT INTO theaters (name, address, description)
VALUES (
    "The Fox Theatre", 
    "527 N Grand Blvd, St. Louis, MO 63103",
    "The Fox Theatre, a former movie palace, is a performing arts center located at 527 N. Grand Blvd. in St. Louis, Missouri. Also known as The Fabulous Fox, it is situated in the arts district of the Grand Center area in Midtown St. Louis, one block north of Saint Louis University.");

CREATE TABLE shows (
    show_id     INTEGER NOT NULL PRIMARY KEY,
    name        TEXT,
    company     TEXT,
    description TEXT,
    theater_id  INTEGER NOT NULL REFERENCES theaters 
);

INSERT INTO shows (name, company, description, theater_id)
VALUES (
    "The Phantom of the Opera", 
    "Andrew Lloyd Webber",
    "A New York institution in itself as the longest-running show in Broadway history, you can't go wrong with a trip to see The Phantom of the Opera if you're looking for that classic Broadway experience. Featuring a timeless score by Andrew Lloyd Webber this Tony Award-winning 'Best Musical' officially opened at the Majestic Theatre in January 1988 and, to its credit, its production values have truly stood the test of time in the face of much younger competition. Based on the French novel Le Fantôme de l'Opéra by Gaston Leroux, the story centers around a dangerous love triangle between young soprano Christine Daaé, her childhood sweetheart, Viscount Raoul de Chagny, and her mysterious music tutor The Phantom, which unfolds at the Paris Opera House with disastrous consequences. Besides the iconic titular number, the musical also boasts such standards as All I Ask of You, Masquerade and The Music of the Night.",
    1 );

INSERT INTO shows (name, company, description, theater_id)
VALUES (
    "Hamilton", 
    "Lin-Manuel Miranda",
    "Our next recommendation perhaps needs no introduction, but here's one anyway... Ladies and Gentlemen, I give you the global phenomenon that is Hamilton! The Broadway premiere of this historic game-changer, conceived and written by musical mastermind Lin-Manuel Miranda, officially opened at the Richard Rodgers Theatre in August 2015... and skyrocketed into mainstream pop culture faster than any of its predecessors. The Grammy and Tony Award-winning 'Best Musical' tells the story of Alexander Hamilton and the Founding Fathers of America through a mix of Hip Hop, RnB and show tunes with a racially diverse cast that reflects American society today. From deathly political rivalries to the country's first sex scandal, Hamilton has often been described as the most hip history lesson you're ever likely to witness. So, don't throw away your shot and book early to avoid disappointment!",
    2 );
