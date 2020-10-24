package warta

import (
	"reflect"
)

type listener interface {
	Close() (err error)
	Callback() interface{}
}

type listen struct {
	name     string
	topic    topic
	callback interface{}
}

func newListener(name string, topic topic, callback interface{}) (l listener, err error) {
	if reflect.ValueOf(callback).Kind() != reflect.Func {
		err = ErrCallbackNotFunction
		return
	}

	l = &listen{
		name:     name,
		topic:    topic,
		callback: callback,
	}

	return
}

func (l *listen) Close() (err error) {

	mu := l.topic.getMutex()
	mu.Lock()
	defer mu.Unlock()

	nListener := make(map[string]listener)

	oListeners := l.topic.getListeners()
	delete(oListeners, l.name)

	for key, val := range oListeners {
		nListener[key] = val
	}

	oListeners = nil

	l.topic.setListeners(nListener)

	return
}

func (l *listen) Callback() interface{} {
	return l.callback
}
