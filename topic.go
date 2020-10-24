package warta

import (
	"fmt"
	"reflect"
)

type TopicName string

type Topic interface {
	getListeners() []interface{}
	AddListener(val interface{})
	RemoveListener(listener interface{}) (err error)
}

type topic struct {
	listeners []interface{}
}

func NewTopic() Topic {
	t := &topic{}

	return t
}

func (t *topic) AddListener(val interface{}) {
	t.listeners = append(t.listeners, val)
}

func (t *topic) getListeners() []interface{} {
	return t.listeners
}

func (t *topic) RemoveListener(listener interface{}) (err error) {
	if reflect.ValueOf(listener).Kind() != reflect.Func {
		err = ErrListenerNotFunction
		return
	}

	fmt.Printf("size t.listener before remove: %d\n", len(t.listeners))
	fmt.Printf("cap t.listener before remove: %d\n", cap(t.listeners))

	var nLis []interface{}
	oLis := t.listeners

	for _, l := range oLis {
		if reflect.ValueOf(l).Pointer() != reflect.ValueOf(listener).Pointer() {
			nLis = append(nLis, l)
		}
	}

	t.listeners = nLis

	fmt.Printf("size t.listener after remove: %d\n", len(t.listeners))
	fmt.Printf("cap t.listener after remove: %d\n", cap(t.listeners))

	return
}
