package warta

import (
	"fmt"
	"sync"
	"testing"

	"github.com/stretchr/testify/assert"
	mock "github.com/stretchr/testify/mock"
)

func mockFnListener(s string) {
	fmt.Println("hello")
}

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

	expectedErr := ErrTopicExists

	topic, err := warta.CreateTopic(name)

	assert.Equal(t, nil, topic)
	assert.Equal(t, expectedErr, err)
}

func Test_CreateTopic_TopicNotYetExists(t *testing.T) {
	var name TopicName = "test-topic"
	mu := &sync.Mutex{}

	warta := warta{
		topics: map[TopicName]topic{},
		mu:     mu,
	}

	listeners := make(map[string]listener)

	expectedTopic := &_topic{
		listeners: listeners,
		mu:        mu,
	}
	expectedTopics := map[TopicName]topic{name: expectedTopic}

	topic, err := warta.CreateTopic(name)

	assert.Equal(t, expectedTopic, topic)
	assert.Equal(t, expectedTopics, warta.topics)
	assert.Equal(t, nil, err)
}

func Test_Broadcast_TopicNotExists(t *testing.T) {
	var name TopicName = "test-topic"
	topics := map[TopicName]topic{}

	warta := warta{
		topics: topics,
		mu:     &sync.Mutex{},
	}

	expectedErr := ErrTopicNotExists

	err := warta.Broadcast(name)

	assert.Equal(t, expectedErr, err)
}

func Test_Broadcast_TopicExists(t *testing.T) {
	var name TopicName = "test-topic"
	mTopic := &mockTopic{}

	listenerID := "4324345"

	fn := mockFnListener

	listen := &listen{
		id:       listenerID,
		topic:    mTopic,
		callback: fn,
	}

	mTopic.On("getListeners").Return(map[string]listener{listenerID: listen})

	topics := map[TopicName]topic{name: mTopic}

	warta := warta{
		topics: topics,
		mu:     &sync.Mutex{},
	}

	err := warta.Broadcast(name, "hello")

	assert.Equal(t, nil, err)
}

func Test_BroadcastCreate_TopicNotExists(t *testing.T) {
	var name TopicName = "test-topic"

	topics := map[TopicName]topic{}

	warta := warta{
		topics: topics,
		mu:     &sync.Mutex{},
	}

	err := warta.BroadcastCreate(name, "hello")

	assert.Equal(t, nil, err)
}

func Test_BroadcastCreate_TopicExists(t *testing.T) {
	var name TopicName = "test-topic"
	mTopic := &mockTopic{}

	listenerID := "4324345"

	fn := mockFnListener

	listen := &listen{
		id:       listenerID,
		topic:    mTopic,
		callback: fn,
	}

	mTopic.On("getListeners").Return(map[string]listener{listenerID: listen})

	topics := map[TopicName]topic{name: mTopic}

	warta := warta{
		topics: topics,
		mu:     &sync.Mutex{},
	}

	var expectedErr error = nil

	err := warta.BroadcastCreate(name, "hello")

	assert.Equal(t, expectedErr, err)
}

func Test_BroadcastClose_TopicNotExists(t *testing.T) {
	var name TopicName = "test-topic"

	topics := map[TopicName]topic{}

	warta := warta{
		topics: topics,
		mu:     &sync.Mutex{},
	}

	expectedErr := ErrTopicNotExists

	err := warta.BroadcastClose(name, "hello")

	assert.Equal(t, expectedErr, err)
}

func Test_BroadcastClose_TopicExists(t *testing.T) {
	var name TopicName = "test-topic"
	mTopic := &mockTopic{}

	listenerID := "4324345"

	fnls := mockFnListener

	listen := &listen{
		id:       listenerID,
		topic:    mTopic,
		callback: fnls,
	}

	mTopic.On("getListeners").Return(map[string]listener{listenerID: listen})

	topics := map[TopicName]topic{name: mTopic}

	warta := warta{
		topics: topics,
		mu:     &sync.Mutex{},
	}

	err := warta.BroadcastClose(name, "hello")

	assert.Equal(t, nil, err)
}

func Test_Listen_TopicNotExists(t *testing.T) {
	var name TopicName = "test-topic"

	topics := map[TopicName]topic{}

	fn := mockFnListener

	warta := warta{
		topics: topics,
		mu:     &sync.Mutex{},
	}

	expectedErr := ErrTopicNotExists

	listener, err := warta.Listen(name, fn)

	assert.Equal(t, nil, listener)
	assert.Equal(t, expectedErr, err)
}

func Test_Listen_TopicExists(t *testing.T) {
	var name TopicName = "test-topic"

	mTopic := &mockTopic{}

	listenerID := "4324345"

	fn := mockFnListener

	expectedListener := &listen{
		id:       listenerID,
		topic:    mTopic,
		callback: fn,
	}

	mTopic.On("addListener", mock.AnythingOfType("func(string)")).Return(expectedListener, nil)

	topics := map[TopicName]topic{name: mTopic}

	warta := warta{
		topics: topics,
		mu:     &sync.Mutex{},
	}

	listener, err := warta.Listen(name, fn)

	assert.Equal(t, expectedListener, listener)
	assert.Equal(t, nil, err)
}

func Test_ListenCreate_TopicExists(t *testing.T) {
	var name TopicName = "test-topic"

	mu := &sync.Mutex{}

	listenerID := "4324345"

	fn := mockFnListener

	mTopic := &mockTopic{}

	expectedListener := &listen{
		callback: fn,
		topic:    mTopic,
		id:       listenerID,
	}

	mTopic.On("addListener", mock.AnythingOfType("func(string)")).Return(expectedListener, nil)

	topics := map[TopicName]topic{name: mTopic}

	warta := warta{
		topics: topics,
		mu:     mu,
	}

	listener, err := warta.ListenCreate(name, fn)

	assert.Equal(t, expectedListener, listener)
	assert.Equal(t, nil, err)
}

func Test_broadcast_ArgsIsFunction(t *testing.T) {
	var name TopicName = "test-topic"

	fn := mockFnListener

	expectedErr := ErrArgsIsFunc

	warta := &warta{}

	err := warta.broadcast(name, fn)

	assert.Equal(t, expectedErr, err)
}

func Test_broadcast_ArgsCountNotMatch(t *testing.T) {
	var name TopicName = "test-topic"
	mTopic := &mockTopic{}

	listenerID := "4324345"

	fn := mockFnListener

	listen := &listen{
		id:       listenerID,
		topic:    mTopic,
		callback: fn,
	}

	mTopic.On("getListeners").Return(map[string]listener{listenerID: listen})

	topics := map[TopicName]topic{name: mTopic}

	warta := warta{
		topics: topics,
		mu:     &sync.Mutex{},
	}

	err := warta.broadcast(name, "hello", "world")

	assert.Equal(t, ErrArgsLenNotMatch, err)
}

func Test_broadcast_ArgsDifferentType(t *testing.T) {
	var name TopicName = "test-topic"
	mTopic := &mockTopic{}

	listenerID := "4324345"

	fn := mockFnListener

	listen := &listen{
		id:       listenerID,
		topic:    mTopic,
		callback: fn,
	}

	mTopic.On("getListeners").Return(map[string]listener{listenerID: listen})

	topics := map[TopicName]topic{name: mTopic}

	warta := warta{
		topics: topics,
		mu:     &sync.Mutex{},
	}

	err := warta.broadcast(name, 1)

	assert.Equal(t, ErrArgsIsDifferent, err)
}
