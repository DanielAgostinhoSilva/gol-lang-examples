create database go_course;

use go_course;

create table posts (
                       id INT AUTO_INCREMENT PRIMARY KEY ,
                       title VARCHAR(50) NOT NULL,
                       body TEXT
);