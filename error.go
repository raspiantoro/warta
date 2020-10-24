package warta

import (
	"fmt"
)

var (
	ErrTopicNotExists      = fmt.Errorf("topic is not exists")
	ErrListenerNotFunction = fmt.Errorf("listener is not a function")
)
