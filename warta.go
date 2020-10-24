package warta

import (
	"reflect"
	"sync"
)

type Warta interface {
	CreateTopic(name TopicName) (topic, error)
	Broadcast(topic TopicName, args ...interface{}) error
	BroadcastCreate(topic TopicName, args ...interface{}) error
	Listen(topic TopicName, callback interface{}) (listener, error)
	ListenCreate(topic TopicName, callback interface{}) (listener, error)
}

type warta struct {
	topics map[TopicName]topic
	mu     *sync.Mutex
}

func New() Warta {
	return &warta{
		topics: make(map[TopicName]topic),
		mu:     &sync.Mutex{},
	}
}

func (w *warta) CreateTopic(name TopicName) (t topic, err error) {
	w.mu.Lock()
	defer w.mu.Unlock()

	if _, exists := w.topics[name]; exists {
		err = ErrTopicExists
		return
	}

	t = newTopic(w.mu)
	w.topics[name] = t

	return
}

func (w *warta) Broadcast(topic TopicName, args ...interface{}) (err error) {
	if _, exists := w.topics[topic]; !exists {
		return ErrTopicNotExists
	}

	err = w.broadcast(topic, args...)
	return
}

func (w *warta) BroadcastCreate(topic TopicName, args ...interface{}) (err error) {
	w.CreateTopic(topic)
	err = w.broadcast(topic, args...)
	return
}

func (w *warta) broadcast(topic TopicName, args ...interface{}) (err error) {

	callArgs := []reflect.Value{}

	for _, arg := range args {
		if reflect.ValueOf(arg).Kind() == reflect.Func {
			return ErrArgsIsFunc
		}

		callArgs = append(callArgs, reflect.ValueOf(arg))
	}

	w.mu.Lock()
	defer w.mu.Unlock()

	for _, listener := range w.topics[topic].getListeners() {
		lArgs := listener.Callback()
		err := compareAndAnalyze(lArgs, args)
		if err != nil {
			return err
		}
		reflect.ValueOf(lArgs).Call(callArgs)
	}

	return nil
}

func (w *warta) Listen(topic TopicName, callback interface{}) (l listener, err error) {
	w.mu.Lock()
	defer w.mu.Unlock()

	if _, exists := w.topics[topic]; !exists {
		err = ErrTopicNotExists
		return
	}

	l, err = w.listen(topic, callback)
	return
}

func (w *warta) ListenCreate(topic TopicName, callback interface{}) (l listener, err error) {
	w.CreateTopic(topic)
	l, err = w.listen(topic, callback)
	return
}

func (w *warta) listen(topic TopicName, callback interface{}) (l listener, err error) {
	l, err = w.topics[topic].addListener(callback)
	return
}
