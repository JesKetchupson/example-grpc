package main

import (
	"example-grpc" //	"example-grpc/pkg/models"
	//	"context"
	//	"fmt"
	"os"
	"os/signal"
	"syscall"
	//	"time"
	//"google.golang.org/grpc"
)

func main() {
	//TODO: switch for db
	// app := example-grpc.NewService()
	newBs := example-grpc.NewServer()
	go newBs.Srv.Serve(newBs.Lis)

	/*move to test
	  c, err := grpc.Dial(":8080", grpc.WithInsecure())
	  	if err != nil {
	  	}
	  	cli := models.NewTasksClient(c)
	  	msg := &models.Text{Text: "someText", Desk: "myDeskription"}
	  	cli.Add(context.Background(), msg)
	  	tl, _ := cli.List(context.Background(), &models.Void{})
	  	fmt.Println(tl)
	*/
	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt, syscall.SIGTERM)
	<-interrupt
	newBs.Srv.GracefulStop()
}
