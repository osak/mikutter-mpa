CREATE TABLE plugins(
    id int NOT NULL AUTO_INCREMENT,
    name varchar(256) NOT NULL,
    version varchar(32) NOT NULL,
    description text NOT NULL,
    url varchar(1024) NOT NULL,
    PRIMARY KEY (id)
);
