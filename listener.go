package warta

//go:generate mockery -name=listener -inpkg

import (
	"reflect"

	"github.com/google/uuid"
)

type listener interface {
	Close()
	Callback() interface{}
	getID() string
}

type listen struct {
	id       string
	topic    topic
	callback interface{}
}

func newListener(topic topic, callback interface{}) (l listener, err error) {
	if reflect.ValueOf(callback).Kind() != reflect.Func {
		err = ErrCallbackNotFunction
		return
	}

	id := uuid.Must(uuid.NewRandom())

	l = &listen{
		id:       id.String(),
		topic:    topic,
		callback: callback,
	}

	return
}

func (l *listen) Close() {

	mu := l.topic.getMutex()
	mu.Lock()
	defer mu.Unlock()

	nListener := make(map[string]listener)

	oListeners := l.topic.getListeners()
	delete(oListeners, l.id)

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

func (l *listen) getID() string {
	return l.id
}
