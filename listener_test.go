package warta

import (
	"sync"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_newListener_GivenNonFuncCallback(t *testing.T) {
	topic := &_topic{}
	callback := "hello"

	expectedErr := ErrCallbackNotFunction

	listener, err := newListener(topic, callback)

	assert.Equal(t, nil, listener)
	assert.Equal(t, expectedErr, err)
}

func Test_newListener_ShouldReturnListener(t *testing.T) {
	topic := &_topic{}
	expectedCallback := mockFnListener

	listener, err := newListener(topic, expectedCallback)

	assert.NotNil(t, listener)
	assert.Nil(t, err)
	assert.NotEqual(t, "", listener.getID())
}

func Test_Close(t *testing.T) {

	topic := &_topic{
		mu: &sync.Mutex{},
	}

	fOne := func(i int) { return }
	fTwo := func(s string) { return }

	lOne := &listen{
		callback: fOne,
		id:       "1",
		topic:    topic,
	}

	lTwo := &listen{
		callback: fTwo,
		id:       "2",
		topic:    topic,
	}

	listeners := map[string]listener{"1": lOne, "2": lTwo}

	topic.listeners = listeners

	expectedListeners := map[string]listener{"2": lTwo}

	lOne.Close()

	assert.Equal(t, expectedListeners, topic.listeners)
}
