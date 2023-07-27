package events

import (
	"wkey-stock/src/events/image_event"
)

type ApiEvents struct {
	Image *image_event.Event
}

func Get() (*ApiEvents, error) {
	image, err := image_event.New()
	if err != nil {
		return nil, err
	}

	return &ApiEvents{
		Image: image,
	}, nil
}
