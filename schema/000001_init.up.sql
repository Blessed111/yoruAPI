CREATE TABLE teams (
    team_id SERIAL PRIMARY KEY,
    team_name VARCHAR(255) NOT NULL,
    player1 VARCHAR(255) NOT NULL,
    player2 VARCHAR(255) NOT NULL,
    player3 VARCHAR(255) NOT NULL,
    player4 VARCHAR(255) NOT NULL,
    player5 VARCHAR(255) NOT NULL,
    password VARCHAR(255) NOT NULL
);