package warta

import (
	"fmt"
	"sync"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_New_ShouldReturnWarta(t *testing.T) {
	expected := &warta{
		topics: make(map[TopicName]topic),
		mu:     &sync.Mutex{},
	}

	actual := New()

	assert.Equal(t, expected, actual)
}

func Test_CreateTopic_TopicAlreadyExists(t *testing.T) {
	var name TopicName = "test-topic"
	topics := map[TopicName]topic{name: &_topic{}}

	warta := warta{
		topics: topics,
		mu:     &sync.Mutex{},
	}

	var expcTopic topic = nil
	expcErr := ErrTopicExists

	topic, err := warta.CreateTopic(name)

	assert.Equal(t, expcTopic, topic)
	assert.Equal(t, expcErr, err)
}

func Test_CreateTopic_TopicNotYetExists(t *testing.T) {
	var name TopicName = "test-topic"
	mu := &sync.Mutex{}

	warta := warta{
		topics: map[TopicName]topic{},
		mu:     mu,
	}

	listeners := make(map[string]listener)

	expcTopic := &_topic{
		listeners: listeners,
		mu:        mu,
	}
	expcTopics := map[TopicName]topic{name: expcTopic}
	var expcErr error = nil

	topic, err := warta.CreateTopic(name)

	assert.Equal(t, expcTopic, topic)
	assert.Equal(t, expcTopics, warta.topics)
	assert.Equal(t, expcErr, err)
}

func Test_Broadcast_TopicNotExists(t *testing.T) {
	var name TopicName = "test-topic"
	topics := map[TopicName]topic{}

	warta := warta{
		topics: topics,
		mu:     &sync.Mutex{},
	}

	expcErr := ErrTopicNotExists

	err := warta.Broadcast(name)

	assert.Equal(t, expcErr, err)
}

func mockFnListener(s string) {
	fmt.Println("hello")
}

func Test_Broadcast_TopicExists(t *testing.T) {
	var name TopicName = "test-topic"
	tp := &mockTopic{}

	lsName := "4324345"

	fnls := mockFnListener

	l := &listen{
		name:     lsName,
		topic:    tp,
		callback: fnls,
	}

	tp.On("getListeners").Return(map[string]listener{lsName: l})

	topics := map[TopicName]topic{name: tp}

	warta := warta{
		topics: topics,
		mu:     &sync.Mutex{},
	}

	var expcErr error = nil

	err := warta.Broadcast(name, "helo")

	assert.Equal(t, expcErr, err)
}

func Test_BroadcastCreate_TopicNotExists(t *testing.T) {
	var name TopicName = "test-topic"

	topics := map[TopicName]topic{}

	warta := warta{
		topics: topics,
		mu:     &sync.Mutex{},
	}

	var expcErr error = nil

	err := warta.BroadcastCreate(name, "helo")

	assert.Equal(t, expcErr, err)
}

func Test_BroadcastCreate_TopicExists(t *testing.T) {
	var name TopicName = "test-topic"
	tp := &mockTopic{}

	lsName := "4324345"

	fnls := mockFnListener

	l := &listen{
		name:     lsName,
		topic:    tp,
		callback: fnls,
	}

	tp.On("getListeners").Return(map[string]listener{lsName: l})

	topics := map[TopicName]topic{name: tp}

	warta := warta{
		topics: topics,
		mu:     &sync.Mutex{},
	}

	var expcErr error = nil

	err := warta.BroadcastCreate(name, "helo")

	assert.Equal(t, expcErr, err)
}
