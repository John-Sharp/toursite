CREATE DATABASE IF NOT EXISTS toursite;
GRANT ALL ON toursite.* TO 'toursite_usr'@'localhost' IDENTIFIED BY 'toursite';
USE toursite;

CREATE TABLE IF NOT EXISTS users (
    id int NOT NULL AUTO_INCREMENT,
    username varchar(255),
    PRIMARY KEY (id)
) ENGINE = INNODB;

CREATE TABLE IF NOT EXISTS tours (
    id int NOT NULL AUTO_INCREMENT,
    operator int NOT NULL,
    title varchar(255),
    description text,
    PRIMARY KEY (id),
    FOREIGN KEY (operator)
        REFERENCES users(id)
        ON DELETE CASCADE
) ENGINE = INNODB;

CREATE TABLE IF NOT EXISTS tour_images (
    id int NOT NULL AUTO_INCREMENT,
    tour_id int NOT NULL,
    path varchar(255),
    PRIMARY KEY (id),
    FOREIGN KEY (tour_id)
        REFERENCES tours(id)
        ON DELETE CASCADE
) ENGINE = INNODB;

CREATE TABLE IF NOT EXISTS categories (
    id int NOT NULL AUTO_INCREMENT,
    name varchar(255),
    PRIMARY KEY(id)
) ENGINE = INNODB;

CREATE TABLE IF NOT EXISTS category_tour_links (
    id int NOT NULL AUTO_INCREMENT,
    category_id int NOT NULL,
    tour_id int NOT NULL,
    PRIMARY KEY(id),
    FOREIGN KEY (category_id)
        REFERENCES categories(id),
    FOREIGN KEY (tour_id)
        REFERENCES tours(id)
) ENGINE = INNODB;
