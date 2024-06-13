create database bancovitor;
use bancovitor;

create table users(
id int auto_increment primary key,
nome varchar(50),
senha varchar(30)
);

select*from users;