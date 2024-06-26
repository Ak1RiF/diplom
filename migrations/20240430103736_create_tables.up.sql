CREATE TABLE users(
    id  SERIAL PRIMARY KEY,
	username VARCHAR(255) NOT NULL,
	password_hash VARCHAR(255) NOT NULL,
	sumExperience INT DEFAULT 0,
	amountExperienceToLvl INT DEFAULT 50,
	lvl INT DEFAULT 1
);

CREATE TABLE quests(
    Id SERIAL PRIMARY KEY, 
    Title VARCHAR(255) NOT NULL,
    Description TEXT NOT NULL,
    Dificulty VARCHAR(255) NOT NULL,
    Completed BOOLEAN DEFAULT FALSE,
    User_Id INT REFERENCES users(id)
);



CREATE TABLE pets(
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    rarity VARCHAR(255) NOT NULL,
    user_id INT REFERENCES users(id)
);


CREATE TABLE eggs(
    id SERIAL PRIMARY KEY,
    rarity VARCHAR(255) NOT NULL
);

CREATE TABLE users_eggs(
    user_id INT REFERENCES users(id) ON DELETE CASCADE,
    egg_id INT REFERENCES eggs(id) ON DELETE CASCADE,
    count_eggs INT DEFAULT 0,
    PRIMARY KEY (user_id, egg_id)
);



INSERT INTO eggs (id, rarity) VALUES (1, 'ordinary');
INSERT INTO eggs (id, rarity) VALUES (2, 'rare');
INSERT INTO eggs (id, rarity) VALUES (3, 'epic');
INSERT INTO eggs (id, rarity) VALUES (4, 'legendary');