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
    // insert, err := db.Query("INSERT INTO vocabulary VALUES ( 2, 'Sinigalia', 'italy', true, LOAD_FILE('D:/Cdownloads/Flag_of_France.png') )")

    // Get all values
    getValues, err := db.Query("SELECT * FROM vocabulary")

	 // if there is an error inserting, handle it
	 if err != nil {
        panic(err.Error())
    }

    // the result object has a method called Next,
    // which is used to iterate through all returned rows.
    for getValues.Next() {
         
        var id int
        var term string
        var defination string
        var favorite bool
        var image []byte
         
        // The result object provided Scan  method
        // to read row data, Scan returns error,
        // if any. Here we read id and name returned.
        err = getValues.Scan(&id, &term, &defination, &favorite, &image)
        //// get.Scan(&id, &name)
         
        // handle error
        if err != nil {
            panic(err)
        }

        fmt.Print(id)
        fmt.Print(term)
        fmt.Print(defination)
        fmt.Println(favorite)
        fmt.Println(image)
    }

    fmt.Println(getValues)

    // be careful deferring Queries if you are using transactions
    defer getValues.Close()
}