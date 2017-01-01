CREATE TABLE IF NOT EXISTS plugins (
    id int NOT NULL AUTO_INCREMENT,
    user_id int NOT NULL,
    name varchar(256) NOT NULL,
    version varchar(32) NOT NULL,
    description text NOT NULL,
    url varchar(1024) NOT NULL,
    PRIMARY KEY (id),
    KEY (user_id)
);

CREATE TABLE IF NOT EXISTS users(
    id int NOT NULL AUTO_INCREMENT,
    login varchar(32) NOT NULL,
    name varchar(256) NOT NULL,
    PRIMARY KEY (id),
    UNIQUE KEY (login)
);

CREATE TABLE IF NOT EXISTS sessions(
    id char(32) NOT NULL,
    user_id int NOT NULL,
    PRIMARY KEY(id)
);
