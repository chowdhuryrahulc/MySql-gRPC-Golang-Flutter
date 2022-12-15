package main

import (
	"context"
	"fmt"
	"log"
	proto "mysqlgolang/gen"
	"net"
	"os"
	"sync"

	_ "github.com/go-sql-driver/mysql"
	"google.golang.org/grpc"
	_ "google.golang.org/grpc/encoding/gzip"
	glog "google.golang.org/grpc/grpclog"
	prt "google.golang.org/protobuf/proto" //implements the same stuff as done by "github.com/golang/protobuf/proto"
)

/*

 */

var grpcLog glog.LoggerV2 

func init() {
	grpcLog = glog.NewLoggerV2(os.Stdout, os.Stdout, os.Stdout)	
}

// This struct gets all the for loop mysql data 
type MySqlData struct {
    id int
    term string
    defination string
    favorite string     //! change to bool later
    image []byte
}

type Server struct {
	MySqlData []*MySqlData
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
   var x []MySqlData
   x = getSqlData() 
   return &proto.Vocab{Word: x}, nil

//todo How do we change []MySqlData to []Word. 
//todo Means os generted to grpc generated

//    &proto.Vocab{Id: 3, Term: "", Defination: "", Favorite: false, Image: []byte{}}, nil
//    x, nil
// PROBLEM: we have to send a list, not 1 term, as done in proto file
}

func main(){
    //todo Now send all this data via grpc
    //todo change from incomming msg to broadcast type
    //todo Favourite change to boolen in proto file
    
    var mySqlData []*MySqlData
	server := &Server{mySqlData, proto.UnimplementedBroadcastServer{}} //! Why did we write unimplement...?

	grpcServer := grpc.NewServer()
	grpcLog.Info("Starting server at port :8080")
	proto.RegisterBroadcastServer(grpcServer, server)
    // This server error is due to the server methord

    // Start the server.
	listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatalf("error creating the server %v", err)
	}
	grpcServer.Serve(listener)
    // fmt.Println(x)
}