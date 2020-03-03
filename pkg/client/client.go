package client

import (
	"context"
	"example/pkg/models"
	"fmt"
)

func ClientList(ctx context.Context, client models.TasksClient) (*models.TaskList, error) {
	l, err := client.List(ctx, &models.Void{})
	if err != nil {
		return nil, err
	}
	for _, t := range l.Tasks {
		fmt.Println(t)
	}
	return nil, fmt.Errorf("%v", "not iml")
}
