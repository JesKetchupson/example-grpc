package main

import (
	"example"
	"example/pkg/models"
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"

	"google.golang.org/grpc"
)

func main() {
	//TODO: switch for db
	// app := example.NewService()
	newBs := example.NewServer()
	go newBs.Srv.Serve(newBs.Lis)

	time.Sleep(2 * time.Second)
	c, err := grpc.Dial(":8080", grpc.WithInsecure())
	if err != nil {
	}
	cli := models.NewTasksClient(c)
	msg := &models.Text{Text: "someText", Desk: "myDeskription"}
	cli.Add(context.Background(), msg)
	tl, _ := cli.List(context.Background(), &models.Void{})
	fmt.Println(tl)

	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt, syscall.SIGTERM)
	<-interrupt
	newBs.Srv.GracefulStop()
}
