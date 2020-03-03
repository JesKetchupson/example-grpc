package example

import (
	"context"
	"example/pkg/models"
	"net"

	"google.golang.org/grpc"
)

//Server instance
type Server struct {
	s    *Service
	Port string
	Srv  *grpc.Server
	Lis  net.Listener
}

// NewServer is a factiory method for Servise
func NewServer() *Server {
	newEs := &Server{}
	newEs.s = NewService()
	newEs.Port = ":8080" //os.Getenv("PORT")
	lis, err := net.Listen("tcp", newEs.Port)
	newEs.Lis = lis
	if err != nil {
		return nil
	}
	newEs.Srv = grpc.NewServer()
	models.RegisterTasksServer(newEs.Srv, newEs)
	return newEs
}

//List server wraper
func (es *Server) List(ctx context.Context, in *models.Void) (*models.TaskList, error) {
	tl, err := es.s.list()
	if err != nil {
		return nil, err
	}
	return tl, nil
}

//Add server wraper
func (es *Server) Add(ctx context.Context, in *models.Text) (*models.Task, error) {
	t, err := es.s.add(in.Text, in.Desk)
	if err != nil {
		return nil, err
	}
	return t, nil
}
