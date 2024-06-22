package fyne_state

import (
	"sync"
)

var (
	state       = make(map[string]interface{})
	subscribers = make(map[string][]chan struct{})
	mutex       sync.Mutex
)

func Set(key string, value interface{}) {
	mutex.Lock()
	defer mutex.Unlock()
	state[key] = value
	for _, ch := range subscribers[key] {
		select {
		case ch <- struct{}{}:
		default:
		}
	}
}

func Get[T any](key string) T {
	mutex.Lock()
	defer mutex.Unlock()

	return state[key].(T)
}

func Subscribe(key string) chan struct{} {
	mutex.Lock()
	defer mutex.Unlock()

	ch := make(chan struct{}, 1)
	subscribers[key] = append(subscribers[key], ch)
	return ch
}

func Unsubscribe(key string, ch chan struct{}) {
	mutex.Lock()
	defer mutex.Unlock()

	chs := subscribers[key]
	for i := range chs {
		if chs[i] == ch {
			subscribers[key] = append(chs[:i], chs[i+1:]...)
			// close(ch)
			break
		}
	}
}
