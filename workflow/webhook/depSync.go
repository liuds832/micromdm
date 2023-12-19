package webhook

import (
	"github.com/pkg/errors"

	"github.com/liuds832/micromdm/platform/dep/sync"
)


func depSyncEvent(topic string, data []byte) (*Event, error) {
	var ev sync.Event
	if err := sync.UnmarshalEvent(data, &ev); err != nil {
		return nil, errors.Wrap(err, "unmarshal depsync event for webhook")
	}

	webhookEvent := Event{
		Topic:     topic,
		DepSyncEvent: &ev,
	}

	return &webhookEvent, nil
}
