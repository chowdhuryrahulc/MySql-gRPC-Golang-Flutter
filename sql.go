package main

func SQLcommands (){
	// INT
	// DECIMAL(10,4)
	// VARCHAR(100)
	// BLOB
	// DATE
	// TIMESTAMP

/*
	show databases;                 : shows all databases available
	create database vocab;
	use vocab;
	show tables;
	1) create table holla(id integer, term text);
	2) CREATE TABLE Persons (
		Personid int NOT NULL AUTO_INCREMENT,
    	LastName varchar(255) NOT NULL,
    	FirstName varchar(255),
    	Age int,
    	PRIMARY KEY (Personid)
	);
	3) create table vocabulary(
		id int not null auto_increment primary key, 
		term varchar(25) not null default 'ThksIsTerm', 
		defination varchar(25), 
		favorite boolean,
		images blob
	);
	Cant store Lists (for lists, use foreign key references)
	drop table persons;
	show columns from vocab;			: show all column names (table structure) in a table
	1) ALTER TABLE Vocabulary ALTER defination SET DEFAULT 'defcon6';
	2) ALTER TABLE Vocabulary ALTER defination DROP DEFAULT;
	Error: cannot alter autoincrement default
	SELECT * FROM vocabulary;			: see all contants of table (vocabulary)
	!ERROR: (DIDNT WORK) INSERT INTO vocabulary VALUES ('Osaka', 'japan', true, LOAD_FILE('D:/Cdownloads/Flag_of_France.png');


CREATE TABLE vocab (id INTEGER PRIMARY KEY AUTOINCREMENT, term TEXT, defination TEXT, example TEXT, url TEXT, favorite BOOLEAN, archive BOOLEAN, current_set TEXT, picture BLOB)");
! Error: primary key autoincrement doesnt work in mysql shell
! Error: Why image is null? LoadImage didnt work. Why?
! IMAGE DIDNT WORK
//todo add default
todo foreign key
todo complex query with foreign keys
todo index searches
todo insert blob image (also as a default)




*/
}