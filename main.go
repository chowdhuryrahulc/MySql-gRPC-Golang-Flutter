package main

import (
    "fmt"
    "database/sql"
    _ "github.com/go-sql-driver/mysql"
)

/*
? Shell Commands
    \sql                            : change from js to sql queries
    \connect root@localhost:3306    : connect to localhost
    
!Error Solved: 
    Workbench Error: 19:36:23	Could not connect, server may not be running.	Unable to connect to localhost	
    Shell Error: MySQL Error 2003 (HY000): Can't connect to MySQL server on 'localhost:3306' (10061)
Sol: Ctrl+alt+del => task manager => services => search mysql80 => right click => start

? GOLANG
    source: https://www.geeksforgeeks.org/how-to-use-go-with-mysql/
*/


func main() {
    fmt.Println("Go MySQL Tutorial")

    // Open up our database connection.
    // I've set up a database by installing mysql
    // The database is called vocab
    db, err := sql.Open("mysql", "root:Chowdhury0511@@tcp(127.0.0.1:3306)/vocab")

    // if there is an error opening the connection, handle it
    if err != nil {
        panic(err.Error())
    }

    // defer the close till after the main function has finished
    // executing
    defer db.Close()

	// perform a db.Query insert
    //* Image also added
    insert, err := db.Query("INSERT INTO vocabulary VALUES ( 2, 'Otranaya', 'russia', true, LOAD_FILE('D:/Cdownloads/Flag_of_France.png') )")

	 // if there is an error inserting, handle it
	 if err != nil {
        panic(err.Error())
    }

    // be careful deferring Queries if you are using transactions
    defer insert.Close()
}