package main

import (
	"context"
	"fmt"
	"log"
	proto "mysqlgolang/gen"
	"net"
	"database/sql"
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
	fmt.Println("getdata1")				//? When flutter is called, this is working (it is printing)
	x := getSqlData()
	fmt.Println("getdata2")
	fmt.Println(x)
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
	//! Error Solved: panic: runtime error: invalid memory address or nil pointer dereference
	// SOL: change var r * Rectangle to
	// 	1) r := &Rectangle{l:10, b:20}
	// 	2) r := &Rectangle{}
	// 	3) r := new(Rectangle)
	// (where Rectangle is a struct with l and b variables and some methords)

	dataRow := &proto.Word{}
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
		errr := getValues.Scan(&dataRow.Id, &dataRow.Term, &dataRow.Defination, &dataRow.Favorite, &dataRow.Image)

		if errr != nil {
			panic(err)
		}

		mysqlAllData = append(mysqlAllData, dataRow)
	}

	// be careful deferring Queries if you are using transactions
	defer getValues.Close()

	return mysqlAllData
}
