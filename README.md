# MoviesAppBackend
MoviesAppBackend

Simple Movie Directory, Exposing REST API, using mysql database.


Table Structure:
```
CREATE TABLE movies (
    id int NOT NULL AUTO_INCREMENT,
    name varchar(255) NOT NULL,
    year int(4) NOT NULL,
    director varchar(255) NOT NULL,
    created timestamp default CURRENT_TIMESTAMP ,
    updated timestamp default CURRENT_TIMESTAMP on update CURRENT_TIMESTAMP,
    PRIMARY KEY (ID)
);
```
```
CREATE TABLE users (
    id int NOT NULL AUTO_INCREMENT,
    email varchar(255) NOT NULL,
    name varchar(255) NOT NULL,
    lastname varchar(255) NOT NULL,
    password varchar(255),
    role varchar(255),
    created timestamp default CURRENT_TIMESTAMP ,
    updated timestamp default CURRENT_TIMESTAMP on update CURRENT_TIMESTAMP,
    PRIMARY KEY (ID)
);
```
