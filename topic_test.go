package warta

import (
	"sync"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_newTopic_ShouldReturnTopic(t *testing.T) {
	mu := &sync.Mutex{}

	expected := &_topic{
		listeners: make(map[string]listener),
		mu:        mu,
	}

	actual := newTopic(mu)

	assert.Equal(t, expected, actual)
}

func Test_setListener(t *testing.T) {
	expected := &sync.Mutex{}

	topic := &_topic{
		mu: expected,
	}

	actual := topic.getMutex()

	assert.Equal(t, expected, actual)
}

func Test_getMutex(t *testing.T) {
	expected := map[string]listener{"123": &listen{}}

	topic := &_topic{}

	topic.setListeners(expected)

	assert.Equal(t, expected, topic.getListeners())
}
