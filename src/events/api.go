package events

import (
	"wkey-stock/src/events/image_event"
	"wkey-stock/src/events/script_event"
)

type ApiEvents struct {
	Script *script_event.Event
	Image  *image_event.Event
}

func Get() (*ApiEvents, error) {
	script, err := script_event.Create()
	if err != nil {
		return nil, err
	}

	image, err := image_event.Create()
	if err != nil {
		return nil, err
	}

	return &ApiEvents{
		Script: script,
		Image:  image,
	}, nil
}
