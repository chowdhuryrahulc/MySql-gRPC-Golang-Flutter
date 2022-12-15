package main

import "database/sql"

//  great source: https://www.geeksforgeeks.org/how-to-use-go-with-mysql/

func getSqlData() []MySqlData {
	var dataRow MySqlData
	mysqlAllData := []MySqlData{}

	// Open up our database connection.
	// I've set up a database by installing mysql
	// The database is called vocab
	db, err := sql.Open("mysql", "root:Chowdhury0511@@tcp(127.0.0.1:3306)/vocab")

	// if there is an error opening the connection, handle it
	if err != nil {
		panic(err.Error())
	}

	// defer the close till after the main function has finished executing
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

		// The result object provided Scan  method
		// to read row data, Scan returns error,
		// if any. Here we read id and name returned.
		err = getValues.Scan(&dataRow.id, &dataRow.term, &dataRow.defination, &dataRow.favorite, &dataRow.image)

		if err != nil {
			panic(err)
		}

		mysqlAllData = append(mysqlAllData, dataRow)
	}

	// be careful deferring Queries if you are using transactions
	defer getValues.Close()

	return mysqlAllData
}