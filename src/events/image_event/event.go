package image_event

import "sync"

type Event struct {
	mutex sync.Mutex
}

func Create() (*Event, error) {
	return &Event{
		mutex: sync.Mutex{},
	}, nil
}
