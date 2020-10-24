package warta

import (
	"errors"
)

var (
	// ErrTopicExists is returned by CreateTopic when topic that
	// want to create is already exists
	ErrTopicExists = errors.New("topic is already exists")

	// ErrTopicNotExists is returned any operation that want to
	// broadcast or listen to any topic that is not already created
	ErrTopicNotExists = errors.New("topic is not exists")

	// ErrCallbackNotFunction is returned by listener when the given
	// callback is not a function
	ErrCallbackNotFunction = errors.New("listener is not a function")

	// ErrArgsIsFunc is returned by broadcast operation when  one of the
	// given args is a function
	ErrArgsIsFunc = errors.New("Cannot use func as arguments")
)
