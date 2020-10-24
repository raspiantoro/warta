package warta

import (
	"errors"
	"reflect"
	"sync"
)

type Emitter interface {
	On(name TopicName, listener interface{}) error
	Emit(name TopicName, args ...interface{}) error
	Topic(name TopicName) (t Topic, err error)
}

type emitter struct {
	topics map[TopicName]Topic
	mu     sync.Mutex
}

//NewEmitter create emitter
func NewEmitter() Emitter {
	e := &emitter{
		topics: make(map[TopicName]Topic),
	}
	return e
}

func (e *emitter) Topic(name TopicName) (t Topic, err error) {
	var exists bool
	if t, exists = e.topics[name]; !exists {
		err = ErrTopicNotExists
	}
	return
}

func (e *emitter) On(name TopicName, listener interface{}) error {

	if reflect.ValueOf(listener).Kind() != reflect.Func {
		return ErrListenerNotFunction
	}

	if _, exists := e.topics[name]; !exists {
		e.topics[name] = NewTopic()
	}

	e.topics[name].AddListener(listener)

	return nil
}

func (e *emitter) Emit(name TopicName, args ...interface{}) error {

	if _, exists := e.topics[name]; !exists {
		return ErrTopicNotExists
	}

	callArgs := []reflect.Value{}

	for _, arg := range args {
		if reflect.ValueOf(arg).Kind() == reflect.Func {
			return errors.New("Cannot use func as arguments")
		}

		callArgs = append(callArgs, reflect.ValueOf(arg))
	}

	for _, listener := range e.topics[name].getListeners() {
		err := compareAndAnalyzeListener(listener, args)
		if err != nil {
			return err
		}
		reflect.ValueOf(listener).Call(callArgs)
	}

	return nil
}

func compareAndAnalyzeListener(listener interface{}, args []interface{}) error {
	x := reflect.TypeOf(listener)
	in := x.NumIn()

	if in != len(args) {
		return errors.New("Arguments length not match")
	}

	for i := 0; i < in; i++ {
		inV := x.In(i)
		argsV := args[i]

		if inV.Kind() != reflect.ValueOf(argsV).Kind() {
			return errors.New("Call using different kind of arguments")
		}
	}

	return nil
}
