CREATE DATABASE my_db;

-- CREATE TABLE student (
-- StdId int NOT NULL,
-- FirstName varchar(45) NOT NULL,
-- LastName varchar(45) DEFAULT NULL,
-- Email varchar(45) NOT NULL,
-- PRIMARY KEY (StdId),
-- UNIQUE (Email)
-- )
CREATE DATABASE my_db;

\c my_db;

CREATE TABLE student (
  StdId SERIAL PRIMARY KEY,
  FirstName VARCHAR(45) NOT NULL,
  LastName VARCHAR(45),
  Email VARCHAR(45) UNIQUE NOT NULL
);
