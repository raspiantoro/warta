package warta

//go:generate mockery -name=topic -inpkg

import (
	"sync"

	uuid "github.com/satori/go.uuid"
)

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
	name := uuid.NewV4().String()

	l, err = newListener(name, t, val)
	if err != nil {
		return
	}

	t.listeners[name] = l
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

// func (t *_topic) RemoveListener(listener interface{}) (err error) {
// 	if reflect.ValueOf(listener).Kind() != reflect.Func {
// 		err = ErrListenerNotFunction
// 		return
// 	}

// 	fmt.Printf("size t.listener before remove: %d\n", len(t.listeners))
// 	fmt.Printf("cap t.listener before remove: %d\n", cap(t.listeners))

// 	var nLis []interface{}
// 	oLis := t.listeners

// 	for _, l := range oLis {
// 		if reflect.ValueOf(l).Pointer() != reflect.ValueOf(listener).Pointer() {
// 			nLis = append(nLis, l)
// 		}
// 	}

// 	t.listeners = nLis

// 	fmt.Printf("size t.listener after remove: %d\n", len(t.listeners))
// 	fmt.Printf("cap t.listener after remove: %d\n", cap(t.listeners))

// 	return
// }
