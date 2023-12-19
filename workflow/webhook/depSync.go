package webhook

import (
	"github.com/pkg/errors"

	"github.com/liuds832/micromdm/mdm"
	"github.com/liuds832/micromdm/platform/dep/sync"
	"github.com/liuds832/micromdm/dep"
)


func depSyncEvent(topic string, data []byte) (*Event, error) {
	var ev sync.Event
	if err := sync.UnmarshalEvent(data, &ev); err != nil {
		return errors.Wrap(err, "unmarshal depsync event")
	}

	webhookEvent := Event{
		Topic:     topic,
		DepSyncEvent: &ev,
	}

	return &webhookEvent, nil
}
