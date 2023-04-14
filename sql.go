package main

func SQLcommands (){
	// INT
	// DECIMAL(10,4)
	// VARCHAR(100)
	// BLOB
	// DATE
	// TIMESTAMP

/*
? Shell Commands
    \sql                            : change from js to sql queries
    \connect root@localhost:3306    : connect to localhost
    
	!Error Solved: 
	Workbench Error: 19:36:23	Could not connect, server may not be running.	Unable to connect to localhost	
	Shell Error: MySQL Error 2003 (HY000): Can't connect to MySQL server on 'localhost:3306' (10061)
	Sol: Ctrl+alt+del => task manager => services => search mysql80 => right click => start
*/

/*
	show databases;                 		: shows all databases available
	create database vocab;
	use vocab;
	drop database xyz;						: delete database
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
	show columns from vocabulary;			: show all column names (table structure) in a table
	desc vocabulary;						: same as show columns
	1) ALTER TABLE Vocabulary ALTER defination SET DEFAULT 'defcon6';
	2) ALTER TABLE Vocabulary ALTER defination DROP DEFAULT;
	Error: cannot alter autoincrement default
	
	? Queries
	1) SELECT * FROM vocabulary;			: see all contants of table (vocabulary)
	2) 
	3) 
	4) 
	5) 
	6) 
	7) 
	
	INSERT: insert into vocabulary values (6, 'paris', 'france', false, load_file('D:/Cdownloads/Flag_of_France.png'));
	!Error solved: in insert, use 'single quotes', not double quotes. Also in image, use single quotes
	!Error2: you need to give id in insert
	

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