package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	proto "mysqlgolang/gen"
	"net"
	"os"

	// "sync"

	_ "github.com/go-sql-driver/mysql"
	"google.golang.org/grpc"
	_ "google.golang.org/grpc/encoding/gzip"
	glog "google.golang.org/grpc/grpclog"
	// prt "google.golang.org/protobuf/proto" //implements the same stuff as done by "github.com/golang/protobuf/proto"
)

/*
! ERROR: check if data is reaching to grocfront-end
! i think the methord is not getting called, otherwise we would have get a print response
! probably we need a pagination methord toarive from client side
! or make something like CreteStream in TensorFlutterDocker side which can get called automatically
todo check how they did that
*/

var grpcLog glog.LoggerV2

func init() {
	grpcLog = glog.NewLoggerV2(os.Stdout, os.Stdout, os.Stdout)
}

// This struct gets all the for loop mysql data
// type MySqlData struct {
//     id int
//     term string
//     defination string
//     favorite string     //! change to bool later
//     image []byte
// }

type Server struct {
	MySqlData []*proto.Word
	proto.UnimplementedBroadcastServer
}

// func (s *Server) GetData(paginate *proto.Paginate) (*proto.Vocab, error) {
// 	// wait := sync.WaitGroup{} 										// to implement go routines
// 	// done := make(chan int)   										// to know when all the go routines are finished
//     //todo Didnt implement size, bcoz it is recieved msg size, not send msg size
//     // size := prt.Size(msg) 											//! gets msg size
// 	// grpcLog.Info("SIZEEEE BROD", size)								// returns 126, 184, 291 etc with compression and without compression
//     // for _, data := range s.MySqlData{
// 		// wait.Add(1)
// 		// go func(conn *Connection) {}
//     // }
//     return &proto.Vocab{}, nil
// }

func (s *Server) GetData(ctx context.Context, paginate *proto.Paginate) (*proto.Vocab, error) {
	//! ERROR SOLVED: context is required, otherwise below in registerBroadcastServer will give an error
	//    var x []MySqlData
	fmt.Println("getdata1") //? When flutter is called, this is working (it is printing)
	x := getSqlData()
	fmt.Println("getdata2")
	// fmt.Println(x)
	return &proto.Vocab{Word: x}, nil

	// 	//todo How do we change []MySqlData to []Word.
	// 	//todo Means os generted to grpc generated

	// 	//    &proto.Vocab{Id: 3, Term: "", Defination: "", Favorite: false, Image: []byte{}}, nil
	// 	//    x, nil
	// 	// PROBLEM: we have to send a list, not 1 term, as done in proto file
}

func main() {
	//todo Now send all this data via grpc
	//todo change from incomming msg to broadcast type
	//// Favourite change to boolen in proto file

	//! testing sql function
	// getSqlData() //! error1: cant import from other file, bcoz it shows error

	// fmt.Println("exeo")
	// x := getSqlData()
	// fmt.Println("goexeo")
	// fmt.Println("x value", x)

	//// uncomment below all
	var mySqlData []*proto.Word
	server := &Server{mySqlData, proto.UnimplementedBroadcastServer{}} //! Why did we write unimplement...?

	grpcServer := grpc.NewServer()
	grpcLog.Info("Starting server at port :8080")
	proto.RegisterBroadcastServer(grpcServer, server)
	//! Error Solved: server error (here) is generally due to the server methords

	// Start the server.
	listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatalf("error creating the server %v", err)
	}
	grpcServer.Serve(listener)
	// fmt.Println(x)
}

func getSqlData() []*proto.Word {
	//? *[]Users would be a pointer to a slice of Users
	//? []*Users would be a slice of pointers to Users.
	//! Problem is all the pointers are pointing to the same memory address. That is dataRow.
	//! Thats why all the values are same
	// SOL: make many dataRows. Like a slice or array. And each pointer to be returned can
	// point to different slice (1,2,3,4...)
	//		! Error encountered: didnt know the meaning of &, and at the for loop entry, we got error
	// 		! Also the for loop i gave an error. So use it as a seperate experiment.
	// 		How to get i value from a .Next() loop
	// SOL2: mysqlAllData is not a slice of pointers, rather a ...
	// Problem: we have to return a slice of pointers.
	// todo WTF is &, or reference?????
	// & = 0xc0000160c0

	//! Error Solved: panic: runtime error: invalid memory address or nil pointer dereference
	// SOL: change var r * Rectangle to
	// 	1) r := &Rectangle{l:10, b:20}
	// 	2) r := &Rectangle{}
	// 	3) r := new(Rectangle)
	// (where Rectangle is a struct with l and b variables and some methords)

	// dataRow := proto.Word{}	//! satish

	//**************************************TESTING***************************************
	// dataRowArray := new([]proto.Word)			//! DELETE LATER
	// dataRowArray = append(dataRowArray, proto.Word{Id:3, Term:"pattaya", Defination:"thailand",  Favorite:true, Image: []byte{1, 8, 90, 44, 63, 77, 90}})
	// fmt.Println(dataRowArray)
	// What we need is to put all this in a scan function. An get a response
	//! Even this code throws error
	// Hypothesis1: Probably bcoz o empty image value, it causes error. So give a fake image byte value
	// 	(DIDNT WORK), we tried adding image value as []byte{1, 8, 90, 44, 63, 77, 90}
	// Real Error: panic: runtime error: index out of range [0] with length 0
	// Means range error while adding values to the dataRowArray[0]
	//
	//**************************************TESTING***************************************

	// x := make([]proto.Word{}, 0)
	mysqlAllData := []*proto.Word{}

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
	// fmt.Println(getValues)// &{0xc0000d6090 0x115b5a0 0xc0000a40f0 <nil> <nil> {{0 0} 0 0 0 0} false <nil> []}

	// if there is an error inserting, handle it
	if err != nil {
		panic(err.Error())
	}

	// the result object has a method called Next,
	// which is used to iterate through all returned rows.

	//// for i := 0; getValues.Next(); i++ {
	//// fmt.Println("ERRORRRRRR")
	//// }
	for getValues.Next() {
		dataRow := proto.Word{} //! SATISH

		// The result object provided Scan  method
		// to read row data, Scan returns error,
		// if any. Here we read id and name returned.
		errr := getValues.Scan(&dataRow.Id, &dataRow.Term, &dataRow.Defination, &dataRow.Favorite, &dataRow.Image)

		if errr != nil {
			fmt.Println(err)
			// panic(err)
		}

		//todo Wat you can do, is put all into a .....
		//todo Why cant it happen? Why cant array go and get the value, and append it into mysqlAllData?
		//todo Is just the dataRowArray[1] giving error?

		fmt.Println("SQLLL")
		fmt.Println(&dataRow)                         // All values in loop are comming
		mysqlAllData = append(mysqlAllData, &dataRow) //! ERROR OCCURING due to pointers append
		fmt.Println(mysqlAllData)
	}

	// be careful deferring Queries if you are using transactions
	defer getValues.Close()

	return mysqlAllData
}

// func getSqlDataRetry() []*proto.Word {
// 	dataRow := proto.Word{}
// 	mysqlAllData := []*proto.Word{}
// 	db, err := sql.Open("mysql", "root:Chowdhury0511@@tcp(127.0.0.1:3306)/vocab")
// 	if err != nil {
// 		panic(err.Error())
// 	}
// 	defer db.Close()
// 	getValues, err := db.Query("SELECT * FROM vocabulary")
// 	if err != nil {
// 		panic(err.Error())
// 	}
// 	for getValues.Next() {
// 		errr := getValues.Scan(&dataRow.Id, &dataRow.Term, &dataRow.Defination, &dataRow.Favorite, &dataRow.Image)
// 		if errr != nil {
// 			fmt.Println(err)
// 		}
// 		fmt.Println("SQLLL")
// 		fmt.Println(&dataRow)                         // All values in loop are comming
// 		mysqlAllData = append(mysqlAllData, &dataRow) //! ERROR OCCURING due to pointers append
// 		fmt.Println(mysqlAllData)
// 	}
// 	defer getValues.Close()

// 	return mysqlAllData
// }
