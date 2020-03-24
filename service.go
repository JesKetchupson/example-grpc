package example-grpc

import (
	"context"
	"math/rand"
	"time"

	"example-grpc/pkg/config"
	"example-grpc/pkg/db"
	"example-grpc/pkg/log"
	"example-grpc/pkg/models"

	"github.com/golang/protobuf/proto"
)

// Service instance
type Service struct {
	database      db.DB
	log           log.Log
	writeConnPoll chan db.DB
	readConnPoll  chan db.DB
	ctx           context.Context
}

// NewService is a factiory method for Servise
func NewService() *Service {
	app := &Service{}
	c := config.Parse()
	switch c.DBType {
	case 1:
		app.database = &db.FileDB{Path: c.DBUri}
	default:
		app.database = &db.FileDB{Path: c.DBUri}
	}
	app.log = &log.StdOutLog{Out: c.Log}
	app.ctx = context.Background()
	app.readConnPoll = make(chan db.DB, 4*10)
	app.writeConnPoll = make(chan db.DB, 1)
	return app
}

func (s *Service) getReadConnection() (db.DB, error) {
	if len(s.readConnPoll) == 0 {
		conn, err := s.database.Open()
		if err != nil {
			return nil, err
		}
		s.readConnPoll <- conn
	}
	return <-s.readConnPoll, nil
}

//getWrieConnection gets write connection from database
func (s *Service) getWrieConnection() (db.DB, error) {
	if len(s.writeConnPoll) == 0 {
		conn, err := s.database.Open()
		if err != nil {
			return nil, err
		}
		s.writeConnPoll <- conn
	}
	return <-s.writeConnPoll, nil
}

func (s *Service) returnReadConnection(conn db.DB) {
	s.readConnPoll <- conn
}

func (s *Service) returnWriteConnection(conn db.DB) {
	s.writeConnPoll <- conn
}
func (s *Service) add(text string, desk string) (*models.Task, error) {
	conn, err := s.getWrieConnection()
	if err != nil {
		s.log.WriteFLog("getWrieConnection err: %w", err)
		return nil, err
	}
	defer s.returnWriteConnection(conn)
	rand.Seed(time.Now().Unix())
	task := &models.Task{
		Data: text,
		Done: false,
		Desc: desk,
		Id:   rand.Uint64(), //TODO uuid
	}

	b, err := proto.Marshal(task)
	if err != nil {
		s.log.WriteFLog("proto.Marshal error: %w", err)
		return nil, err
	}

	err = conn.Insert(b)
	if err != nil {
		s.log.WriteFLog("conn.Insert  error: %w", err)
		return nil, err
	}

	return task, nil
}
func (s *Service) list() (*models.TaskList, error) {
	conn, err := s.getReadConnection()
	if err != nil {
		s.log.WriteFLog(" getReadConnection err: %w", err)
		return nil, err
	}
	defer s.returnReadConnection(conn)

	b, err := conn.List()
	if err != nil {
		s.log.WriteFLog("database list error: %w", err)
		return nil, err
	}

	tasks := &models.TaskList{}
	for i := range b {
		var task models.Task
		if len(b[i]) > 0 {
			err = proto.Unmarshal(b[i], &task)
			if err != nil {
				s.log.WriteFLog("proto.Unmarshal error: %w", err)
				return nil, err
			}
			if !task.GetDone() {
				tasks.Tasks = append(tasks.Tasks, &task)
			}
		}
	}

	// for {
	// 	if len(b) == 0 || len(b) < 8 {
	// 		return tasks, nil
	// 	}
	// 	var l uint64
	// 	shift := unsafe.Sizeof(l)
	// 	binary.Read(bytes.NewReader(b[:shift]), binary.LittleEndian, &l)
	// 	b = b[shift:]
	// 	var task models.Task
	// 	err = proto.Unmarshal(b[:l], &task)
	// 	b = b[l:]
	// 	tasks.Tasks = append(tasks.Tasks, &task)
	// 	if err != nil {
	// 		s.log.WriteFLog("%v", err)
	// 		return nil, err
	// 	}
	// }
	s.readConnPoll <- conn
	return tasks, nil
}
