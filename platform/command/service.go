// Package command provides utilities for creating MDM Payloads.
package command

import (
	mdmsvc "github.com/liuds832/micromdm/mdm"
	"github.com/liuds832/micromdm/mdm/mdm"
	"github.com/liuds832/micromdm/platform/pubsub"
	"golang.org/x/net/context"
)

type Service interface {
	NewCommand(context.Context, *mdm.CommandRequest) (*mdm.CommandPayload, error)
	NewRawCommand(context.Context, *RawCommand) error
	ClearQueue(ctx context.Context, udid string) error
	ViewQueue(ctx context.Context, udid string) ([]*mdmsvc.Command, error)
}

// Queue is an MDM Command Queue.
type Queue interface {
	Clear(context.Context, mdmsvc.CheckinEvent) error
	ViewQueue(context.Context, mdmsvc.CheckinEvent) ([]*mdmsvc.Command, error)
}

type CommandService struct {
	publisher pubsub.Publisher
	queue     Queue
}

func New(pub pubsub.Publisher, queue Queue) (*CommandService, error) {
	svc := CommandService{
		publisher: pub,
		queue:     queue,
	}
	return &svc, nil
}
