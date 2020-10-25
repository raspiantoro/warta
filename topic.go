package warta

//go:generate mockery -name=topic -inpkg -testonly

import (
	"sync"
)

// TopicName define new type for topic name
type TopicName string

type topic interface {
	getListeners() map[string]listener
	setListeners(listeners map[string]listener)
	addListener(val interface{}) (listener, error)
	getMutex() *sync.Mutex
}

type _topic struct {
	listeners map[string]listener
	mu        *sync.Mutex
}

func newTopic(mu *sync.Mutex) topic {
	listeners := make(map[string]listener)

	t := &_topic{
		listeners: listeners,
		mu:        mu,
	}

	return t
}

func (t *_topic) addListener(val interface{}) (l listener, err error) {

	l, err = newListener(t, val)
	if err != nil {
		return
	}

	t.listeners[l.getID()] = l
	return
}

func (t *_topic) getListeners() map[string]listener {
	return t.listeners
}

func (t *_topic) setListeners(listeners map[string]listener) {
	t.listeners = listeners
}

func (t *_topic) getMutex() *sync.Mutex {
	return t.mu
}
