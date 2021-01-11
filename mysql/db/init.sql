CREATE DATABASE IF NOT EXISTS golang_rest_api;
USE golang_rest_api;

CREATE TABLE IF NOT EXISTS posts (
  id          INT PRIMARY KEY NOT NULL AUTO_INCREMENT,
  title       VARCHAR(255) NOT NULL,
  content     VARCHAR(255) NOT NULL,
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;